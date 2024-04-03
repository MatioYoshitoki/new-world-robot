package biz

import (
	"bytes"
	"errors"
	"fmt"
	errors2 "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	cmap "github.com/orcaman/concurrent-map/v2"
	"io"
	"math/rand"
	"net/http"
	apipb "new-world-robot/api/new-world-api/v1"
	v1 "new-world-robot/api/new-world-robot/v1"
	sharedpb "new-world-robot/api/shared/v1"
	"new-world-robot/internal/conf"
	"new-world-robot/pkg/biz_errors"
	"new-world-robot/pkg/consts"
	"new-world-robot/pkg/convert"
	"new-world-robot/pkg/net_utils"
	"new-world-robot/pkg/utils"
	"time"
)

type debugInfo struct {
	LatestSleep   time.Time
	LatestAlive   time.Time
	LatestTimer   time.Time
	LatestWaterId int64
}

type Robot struct {
	Id                                 int64
	ticker                             time.Ticker
	hc                                 *http.Client
	rd                                 *rand.Rand
	memory                             *v1.RobotMemory
	bs                                 *conf.Bootstrap
	statisticsRequestCountMap          cmap.ConcurrentMap[string, int64]
	statisticsRequestTotalCostMap      cmap.ConcurrentMap[string, int64]
	defaultRateMap                     []func() error
	canCreateFishRateMap               []func() error
	canExpandParkingRateMap            []func() error
	bothCreateFishExpandParkingRateMap []func() error
	noFishNoParkingRateMap             []func() error
	di                                 debugInfo
}

func NewRobot(id int64, hc *http.Client, bs *conf.Bootstrap, rd *rand.Rand, face, fKey string) *Robot {
	robot := &Robot{
		Id:                            id,
		hc:                            hc,
		bs:                            bs,
		rd:                            rd,
		statisticsRequestCountMap:     cmap.New[int64](),
		statisticsRequestTotalCostMap: cmap.New[int64](),
		memory: &v1.RobotMemory{
			Auth: &v1.RobotMemory_Auth{
				Face: face,
				FKey: fKey,
				Uid:  face,
			},
			Assets: &v1.RobotMemory_Assets{},
			Fishes: &v1.RobotMemory_Fishes{
				FishList: make([]*v1.RobotMemory_Fishes_Fish, 0),
			},
			ParkingMap: make(map[string]string),
		},
		di: debugInfo{
			LatestTimer: time.UnixMilli(0),
		},
	}
	robot.defaultRateMap = []func() error{robot.tryRandomOperateFish}
	robot.canCreateFishRateMap = []func() error{robot.tryRandomOperateFish, robot.tryCreateFish}
	robot.canExpandParkingRateMap = []func() error{robot.tryRandomOperateFish, robot.tryExpandParking}
	robot.bothCreateFishExpandParkingRateMap = []func() error{robot.tryRandomOperateFish, robot.tryExpandParking, robot.tryCreateFish}
	robot.noFishNoParkingRateMap = []func() error{robot.tryExpandParking}
	return robot
}

func (r *Robot) Work(waterId int64) {
	time.Sleep(utils.RandomDuration(0, 200, time.Millisecond, r.rd))
	if r.di.LatestWaterId != 0 && waterId < r.di.LatestWaterId {
		log.Debugf("wi: %d,lwi: %d", waterId, r.di.LatestWaterId)
	}
	r.di.LatestWaterId = waterId
	if err := r.tryRefreshAccessToken(); err != nil {
		log.Error("auth failed ", err)
		return
	}
	if err := r.tryRefreshAssets(); err != nil {
		log.Error("refresh assets failed ", err)
		return
	}
	if err := r.tryRefreshFishList(); err != nil {
		log.Error("refresh fish list failed ", err)
		return
	}
	if err := r.tryRefreshParkingMap(); err != nil {
		log.Error("refresh parking list failed ", err)
		return
	}
	fishCount := len(r.memory.Fishes.FishList)
	usefulCount, inactiveCount := r.parkingCountByStatus()
	//log.Context(ctx).Debugf("uid=%s, fish_count=%d, useful_count=%d, inactive_count=%d", r.memory.Auth.Uid, fishCount, usefulCount, inactiveCount)
	op := r.tryRandomOperateFish
	if usefulCount > 0 && inactiveCount > 0 && fishCount > 0 {
		oIdx := r.rd.Intn(len(r.bothCreateFishExpandParkingRateMap))
		op = r.bothCreateFishExpandParkingRateMap[oIdx]
	} else if usefulCount > 0 && fishCount > 0 {
		oIdx := r.rd.Intn(len(r.canCreateFishRateMap))
		op = r.canCreateFishRateMap[oIdx]
	} else if inactiveCount > 0 && fishCount > 0 {
		oIdx := r.rd.Intn(len(r.canExpandParkingRateMap))
		op = r.canExpandParkingRateMap[oIdx]
	} else if fishCount == 0 && usefulCount > 0 {
		op = r.tryCreateFish
	} else if fishCount == 0 && inactiveCount > 0 {
		oIdx := r.rd.Intn(len(r.noFishNoParkingRateMap))
		op = r.noFishNoParkingRateMap[oIdx]
	} else {
		oIdx := r.rd.Intn(len(r.defaultRateMap))
		op = r.defaultRateMap[oIdx]
	}
	if err := op(); err != nil {
		log.Errorf("[uid=%s]op failed %v", r.memory.Auth.Uid, err)
		return
	}
}

func (r *Robot) parkingCountByStatus() (usefulCount, inactiveCount int) {
	for _, status := range r.memory.ParkingMap {
		if status == apipb.FishParkingStatus_fish_parking_unused.String() {
			usefulCount++
		} else if status == apipb.FishParkingStatus_fish_parking_inactive.String() {
			inactiveCount++
		}
	}
	return usefulCount, inactiveCount
}

func (r *Robot) tryRefreshFishList() error {
	fishes, err := r.fishList()
	if err != nil {
		if errors.Is(err, biz_errors.AuthError) {
			r.forgetMe()
		}
		return err
	}
	fs := make([]*v1.RobotMemory_Fishes_Fish, len(fishes.List))
	for i, d := range fishes.List {
		fs[i] = &v1.RobotMemory_Fishes_Fish{
			FishId: d.Id,
			Statue: d.Statue,
		}
	}
	r.memory.Fishes.FishList = fs
	r.randomSleep()
	return nil
}

func (r *Robot) tryRefreshAssets() error {
	if assetRst, err := r.asset(); err != nil {
		if errors.Is(err, biz_errors.AuthError) {
			r.forgetMe()
		}
		return err
	} else {
		r.memory.Assets.Gold = assetRst.Gold
		r.memory.Assets.Exp = assetRst.Exp
		r.memory.Assets.Level = assetRst.Level
	}
	r.randomSleep()
	return nil
}

func (r *Robot) tryRefreshAccessToken() error {
	if r.memory.Auth.AccessToken == "" {
		if authRst, err := r.auth(); err != nil {
			return err
		} else {
			r.memory.Auth.AccessToken = authRst.AccessToken
		}
		r.randomSleep()
	}
	return nil
}

func (r *Robot) tryRefreshParkingMap() error {
	if parkingListResult, err := r.parkingList(); err != nil {
		if errors.Is(err, biz_errors.AuthError) {
			r.forgetMe()
		}
		return err
	} else {
		for _, parking := range parkingListResult.ParkingList {
			r.memory.ParkingMap[parking.Parking.String()] = parking.Status.String()
		}
	}
	r.randomSleep()
	return nil
}

func (r *Robot) tryRandomOperateFish() error {
	fishSize := len(r.memory.Fishes.FishList)
	fIdx := r.rd.Intn(fishSize)
	fish := r.memory.Fishes.FishList[fIdx]
	var err error
	if fish.Statue == sharedpb.FishStatus_alive {
		_, err = r.fishSleep(fish.FishId)
	} else if fish.Statue == sharedpb.FishStatus_sleep {
		_, err = r.fishAlive(fish.FishId)
	} else if fish.Statue == sharedpb.FishStatus_dead {
		_, err = r.fishRefining(fish.FishId)
	}
	if err != nil {
		if errors.Is(err, biz_errors.AuthError) {
			r.forgetMe()
		}
		log.Errorf("fish: %+v", fish)
		return err
	}
	r.randomSleep()
	return nil
}

func (r *Robot) tryCreateFish() error {
	if _, err := r.createFish(); err != nil {
		if errors.Is(err, biz_errors.AuthError) {
			r.forgetMe()
		}
		return err
	}
	r.randomSleep()
	return nil
}

func (r *Robot) tryExpandParking() error {
	if _, err := r.expandParking(); err != nil {
		if errors.Is(err, biz_errors.AuthError) {
			r.forgetMe()
		}
		return err
	}
	r.randomSleep()
	return nil
}

func (r *Robot) randomSleep() {
	time.Sleep(utils.RandomDuration(1000, 3000, time.Millisecond, r.rd))
}

func (r *Robot) forgetMe() {
	r.memory.Auth.AccessToken = ""
}

func (r *Robot) auth() (*apipb.AuthResult, error) {
	start := time.Now().UnixMilli()
	defer r.countDown(start, r.bs.App.AuthUrl)
	face := r.memory.Auth.Face
	fk := r.memory.Auth.FKey
	b := convert.StringToBytes(fmt.Sprintf(consts.AuthRequestParamTemplate, face, fk))
	request, err := r.basePostRequest(r.bs.App.AuthUrl, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.AuthResult](r.hc, request)
	if err != nil {
		return nil, err
	}
	if rst.Code != 0 {
		return nil, errors2.New(int(rst.Code), rst.Message, rst.Message)
	}
	return &rst.Data, nil
}

func (r *Robot) asset() (*apipb.AssetResult, error) {
	start := time.Now().UnixMilli()
	defer r.countDown(start, r.bs.App.AssetUrl)
	request, err := r.basePostRequest(r.bs.App.AssetUrl, bytes.NewBuffer(consts.EmptyRequestParam))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.AssetResult](r.hc, request)
	if err != nil {
		return nil, err
	}
	if rst.Code == 401 {
		return nil, biz_errors.AuthError
	} else if rst.Code != 0 {
		return nil, errors2.New(int(rst.Code), rst.Message, rst.Message)
	}
	return &rst.Data, nil
}

func (r *Robot) createFish() (*apipb.FishCreateResult, error) {
	start := time.Now().UnixMilli()
	defer r.countDown(start, r.bs.App.CreateFishUrl)
	request, err := r.basePostRequest(r.bs.App.CreateFishUrl, bytes.NewBuffer(consts.EmptyRequestParam))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.FishCreateResult](r.hc, request)
	if err != nil {
		return nil, err
	}
	if rst.Code == 401 {
		return nil, biz_errors.AuthError
	} else if rst.Code != 0 {
		return nil, errors2.New(int(rst.Code), rst.Message, rst.Message)
	}
	return &rst.Data, nil
}

func (r *Robot) fishList() (*apipb.FishListResult, error) {
	start := time.Now().UnixMilli()
	defer r.countDown(start, r.bs.App.FishListUrl)
	request, err := r.basePostRequest(r.bs.App.FishListUrl, bytes.NewBuffer(consts.EmptyRequestParam))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.FishListResult](r.hc, request)
	if err != nil {
		return nil, err
	}
	if rst.Code == 401 {
		return nil, biz_errors.AuthError
	} else if rst.Code != 0 {
		return nil, errors2.New(int(rst.Code), rst.Message, rst.Message)
	}
	return &rst.Data, nil
}

func (r *Robot) fishSleep(fishId int64) (*apipb.FishSleepResult, error) {
	start := time.Now().UnixMilli()
	defer r.countDown(start, r.bs.App.FishSleepUrl)
	param := convert.StringToBytes(fmt.Sprintf(consts.OnlyFishIdParamTemplate, fishId))
	request, err := r.basePostRequest(r.bs.App.FishSleepUrl, bytes.NewBuffer(param))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.FishSleepResult](r.hc, request)
	if err != nil {
		return nil, err
	}
	if rst.Code == 401 {
		return nil, biz_errors.AuthError
	} else if rst.Code != 0 {
		log.Warnf("latest sleep time: %s , now: %s", r.di.LatestSleep.Format(time.DateTime), time.Now().Format(time.DateTime))
		return nil, errors2.New(int(rst.Code), rst.Message, rst.Message)
	}
	r.di.LatestSleep = time.Now()
	return &rst.Data, nil
}

func (r *Robot) fishAlive(fishId int64) (*apipb.FishAliveResult, error) {
	start := time.Now().UnixMilli()
	defer r.countDown(start, r.bs.App.FishAliveUrl)
	param := convert.StringToBytes(fmt.Sprintf(consts.OnlyFishIdParamTemplate, fishId))
	request, err := r.basePostRequest(r.bs.App.FishAliveUrl, bytes.NewBuffer(param))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.FishAliveResult](r.hc, request)
	if err != nil {
		return nil, err
	}
	if rst.Code == 401 {
		return nil, biz_errors.AuthError
	} else if rst.Code != 0 {
		log.Warnf("latest alive time: %s , now: %s", r.di.LatestAlive.Format(time.DateTime), time.Now().Format(time.DateTime))
		return nil, errors2.New(int(rst.Code), rst.Message, rst.Message)
	}
	r.di.LatestAlive = time.Now()
	return &rst.Data, nil
}

func (r *Robot) fishRefining(fishId int64) (*apipb.FishRefiningResult, error) {
	start := time.Now().UnixMilli()
	defer r.countDown(start, r.bs.App.FishRefiningUrl)
	param := convert.StringToBytes(fmt.Sprintf(consts.OnlyFishIdParamTemplate, fishId))
	request, err := r.basePostRequest(r.bs.App.FishRefiningUrl, bytes.NewBuffer(param))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.FishRefiningResult](r.hc, request)
	if err != nil {
		return nil, err
	}
	if rst.Code == 401 {
		return nil, biz_errors.AuthError
	} else if rst.Code != 0 {
		return nil, errors2.New(int(rst.Code), rst.Message, rst.Message)
	}
	return &rst.Data, nil
}

func (r *Robot) poolRank(page, pageSize int32) (*apipb.FishPoolRankResult, error) {
	start := time.Now().UnixMilli()
	defer r.countDown(start, r.bs.App.PoolRankUrl)
	param := convert.StringToBytes(fmt.Sprintf(consts.OnlyPageParamTemplate, page, pageSize))
	request, err := r.basePostRequest(r.bs.App.PoolRankUrl, bytes.NewBuffer(param))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.FishPoolRankResult](r.hc, request)
	if err != nil {
		return nil, err
	}
	if rst.Code == 401 {
		return nil, biz_errors.AuthError
	} else if rst.Code != 0 {
		return nil, errors2.New(int(rst.Code), rst.Message, rst.Message)
	}
	return &rst.Data, nil
}

func (r *Robot) expandParking() (*apipb.ExpandParkingResult, error) {
	start := time.Now().UnixMilli()
	defer r.countDown(start, r.bs.App.ExpandParkingUrl)
	request, err := r.basePostRequest(r.bs.App.ExpandParkingUrl, bytes.NewBuffer(consts.EmptyRequestParam))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.ExpandParkingResult](r.hc, request)
	if err != nil {
		return nil, err
	}
	if rst.Code == 401 {
		return nil, biz_errors.AuthError
	} else if rst.Code != 0 {
		return nil, errors2.New(int(rst.Code), rst.Message, rst.Message)
	}
	return &rst.Data, nil
}

func (r *Robot) parkingList() (*apipb.ParkingListResult, error) {
	start := time.Now().UnixMilli()
	defer r.countDown(start, r.bs.App.ParkingListUrl)
	request, err := r.basePostRequest(r.bs.App.ParkingListUrl, bytes.NewBuffer(consts.EmptyRequestParam))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.ParkingListResult](r.hc, request)
	if err != nil {
		return nil, err
	}
	if rst.Code == 401 {
		return nil, biz_errors.AuthError
	} else if rst.Code != 0 {
		return nil, errors2.New(int(rst.Code), rst.Message, rst.Message)
	}
	return &rst.Data, nil
}

func (r *Robot) marketList(page, pageSize int32) (*apipb.MarketListResult, error) {
	start := time.Now().UnixMilli()
	defer r.countDown(start, r.bs.App.MarketListUrl)
	param := convert.StringToBytes(fmt.Sprintf(consts.OnlyPageParamTemplate, page, pageSize))
	request, err := r.basePostRequest(r.bs.App.MarketListUrl, bytes.NewBuffer(param))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.MarketListResult](r.hc, request)
	if err != nil {
		return nil, err
	}
	if rst.Code == 401 {
		return nil, biz_errors.AuthError
	} else if rst.Code != 0 {
		return nil, errors2.New(int(rst.Code), rst.Message, rst.Message)
	}
	return &rst.Data, nil
}

func (r *Robot) marketDetail(productId int64) (*apipb.MarketDetailResult, error) {
	start := time.Now().UnixMilli()
	defer r.countDown(start, r.bs.App.MarketDetailUrl)
	param := convert.StringToBytes(fmt.Sprintf(consts.OnlyProductIdParamTemplate, productId))
	request, err := r.basePostRequest(r.bs.App.MarketDetailUrl, bytes.NewBuffer(param))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.MarketDetailResult](r.hc, request)
	if err != nil {
		return nil, err
	}
	if rst.Code == 401 {
		return nil, biz_errors.AuthError
	} else if rst.Code != 0 {
		return nil, errors2.New(int(rst.Code), rst.Message, rst.Message)
	}
	return &rst.Data, nil
}

func (r *Robot) marketSell(fishId, price int64) (*apipb.MarketSellResult, error) {
	start := time.Now().UnixMilli()
	defer r.countDown(start, r.bs.App.MarketSellUrl)
	param := convert.StringToBytes(fmt.Sprintf(consts.MarketSellParamTemplate, fishId, price))
	request, err := r.basePostRequest(r.bs.App.MarketSellUrl, bytes.NewBuffer(param))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.MarketSellResult](r.hc, request)
	if err != nil {
		return nil, err
	}
	if rst.Code == 401 {
		return nil, biz_errors.AuthError
	} else if rst.Code != 0 {
		return nil, errors2.New(int(rst.Code), rst.Message, rst.Message)
	}
	return &rst.Data, nil
}

func (r *Robot) marketStopSell(productId int64) (*apipb.MarketStopSellResult, error) {
	start := time.Now().UnixMilli()
	defer r.countDown(start, r.bs.App.MarketStopSellUrl)
	param := convert.StringToBytes(fmt.Sprintf(consts.OnlyProductIdParamTemplate, productId))
	request, err := r.basePostRequest(r.bs.App.MarketStopSellUrl, bytes.NewBuffer(param))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.MarketStopSellResult](r.hc, request)
	if err != nil {
		return nil, err
	}
	if rst.Code == 401 {
		return nil, biz_errors.AuthError
	} else if rst.Code != 0 {
		return nil, errors2.New(int(rst.Code), rst.Message, rst.Message)
	}
	return &rst.Data, nil
}

func (r *Robot) marketBuy(productId int64) (*apipb.MarketBuyResult, error) {
	start := time.Now().UnixMilli()
	defer r.countDown(start, r.bs.App.MarketBuyUrl)
	param := convert.StringToBytes(fmt.Sprintf(consts.OnlyProductIdParamTemplate, productId))
	request, err := r.basePostRequest(r.bs.App.MarketBuyUrl, bytes.NewBuffer(param))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.MarketBuyResult](r.hc, request)
	if err != nil {
		return nil, err
	}
	if rst.Code == 401 {
		return nil, biz_errors.AuthError
	} else if rst.Code != 0 {
		return nil, errors2.New(int(rst.Code), rst.Message, rst.Message)
	}
	return &rst.Data, nil
}

func (r *Robot) basePostRequest(url string, params io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodPost, r.bs.App.BaseUrl+url, params)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")
	if r.memory.Auth.Uid != "" {
		request.Header.Add(consts.HeaderUid, r.memory.Auth.Uid)
	}
	if r.memory.Auth.AccessToken != "" {
		request.Header.Add(consts.HeaderAccessToken, r.memory.Auth.AccessToken)
	}
	return request, nil
}

func (r *Robot) countDown(startMs int64, url string) {
	if get, exist := r.statisticsRequestCountMap.Get(url); exist {
		r.statisticsRequestCountMap.Set(url, get+1)
	} else {
		r.statisticsRequestCountMap.Set(url, 1)
	}
	nMs := time.Now().UnixMilli()
	if get, exist := r.statisticsRequestTotalCostMap.Get(url); exist {
		r.statisticsRequestTotalCostMap.Set(url, get+nMs-startMs)
	} else {
		r.statisticsRequestTotalCostMap.Set(url, nMs-startMs)
	}
}

func (r *Robot) Statistics() (cmap.ConcurrentMap[string, int64], cmap.ConcurrentMap[string, int64]) {
	return r.statisticsRequestCountMap, r.statisticsRequestTotalCostMap
}
