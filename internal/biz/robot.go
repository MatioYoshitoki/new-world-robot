package biz

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	errors2 "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
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

type Robot struct {
	Id                                 int64
	ticker                             time.Ticker
	hc                                 *http.Client
	rd                                 *rand.Rand
	memory                             *v1.RobotMemory
	bs                                 *conf.Bootstrap
	defaultRateMap                     []func(ctx context.Context) error
	canCreateFishRateMap               []func(ctx context.Context) error
	canExpandParkingRateMap            []func(ctx context.Context) error
	bothCreateFishExpandParkingRateMap []func(ctx context.Context) error
	noFishNoParkingRateMap             []func(ctx context.Context) error
}

func NewRobot(id int64, hc *http.Client, bs *conf.Bootstrap, rd *rand.Rand, face, fKey string) *Robot {
	robot := &Robot{
		Id: id,
		hc: hc,
		bs: bs,
		rd: rd,
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
	}
	robot.defaultRateMap = []func(ctx context.Context) error{robot.tryRandomOperateFish}
	robot.canCreateFishRateMap = []func(ctx context.Context) error{robot.tryRandomOperateFish, robot.tryCreateFish}
	robot.canExpandParkingRateMap = []func(ctx context.Context) error{robot.tryRandomOperateFish, robot.tryExpandParking}
	robot.bothCreateFishExpandParkingRateMap = []func(ctx context.Context) error{robot.tryRandomOperateFish, robot.tryExpandParking, robot.tryCreateFish}
	robot.noFishNoParkingRateMap = []func(ctx context.Context) error{robot.tryExpandParking}
	return robot
}

func (r *Robot) Work(ctx context.Context) {
	if err := r.tryRefreshAccessToken(ctx); err != nil {
		log.Context(ctx).Error("auth failed ", err)
		return
	}
	if err := r.tryRefreshAssets(ctx); err != nil {
		log.Context(ctx).Error("refresh assets failed ", err)
		return
	}
	if err := r.tryRefreshFishList(ctx); err != nil {
		log.Context(ctx).Error("refresh fish list failed ", err)
		return
	}
	if err := r.tryRefreshParkingMap(ctx); err != nil {
		log.Context(ctx).Error("refresh parking list failed ", err)
		return
	}
	fishCount := len(r.memory.Fishes.FishList)
	usefulCount, inactiveCount := r.parkingCountByStatus()
	log.Context(ctx).Debugf("uid=%s, fish_count=%d, useful_count=%d, inactive_count=%d", r.memory.Auth.Uid, fishCount, usefulCount, inactiveCount)
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
	if err := op(ctx); err != nil {
		log.Context(ctx).Error("op failed ", err)
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

func (r *Robot) tryRefreshFishList(ctx context.Context) error {
	fishes, err := r.fishList(ctx)
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

func (r *Robot) tryRefreshAssets(ctx context.Context) error {
	if assetRst, err := r.asset(ctx); err != nil {
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

func (r *Robot) tryRefreshAccessToken(ctx context.Context) error {
	if r.memory.Auth.AccessToken == "" {
		if authRst, err := r.auth(ctx); err != nil {
			return err
		} else {
			r.memory.Auth.AccessToken = authRst.AccessToken
		}
		r.randomSleep()
	}
	return nil
}

func (r *Robot) tryRefreshParkingMap(ctx context.Context) error {
	if parkingListResult, err := r.parkingList(ctx); err != nil {
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

func (r *Robot) tryRandomOperateFish(ctx context.Context) error {
	fishSize := len(r.memory.Fishes.FishList)
	fIdx := r.rd.Intn(fishSize)
	fish := r.memory.Fishes.FishList[fIdx]
	var err error
	if fish.Statue == sharedpb.FishStatus_alive {
		_, err = r.fishSleep(ctx, fish.FishId)
	} else if fish.Statue == sharedpb.FishStatus_sleep {
		_, err = r.fishAlive(ctx, fish.FishId)
	} else if fish.Statue == sharedpb.FishStatus_dead {
		_, err = r.fishRefining(ctx, fish.FishId)
	}
	if err != nil {
		if errors.Is(err, biz_errors.AuthError) {
			r.forgetMe()
		}
		return err
	}
	r.randomSleep()
	return nil
}

func (r *Robot) tryCreateFish(ctx context.Context) error {
	if _, err := r.createFish(ctx); err != nil {
		if errors.Is(err, biz_errors.AuthError) {
			r.forgetMe()
		}
		return err
	}
	r.randomSleep()
	return nil
}

func (r *Robot) tryExpandParking(ctx context.Context) error {
	if _, err := r.expandParking(ctx); err != nil {
		if errors.Is(err, biz_errors.AuthError) {
			r.forgetMe()
		}
		return err
	}
	r.randomSleep()
	return nil
}

func (r *Robot) randomSleep() {
	time.Sleep(utils.RandomDuration(1, 3, r.rd))
}

func (r *Robot) forgetMe() {
	r.memory.Auth.AccessToken = ""
}

func (r *Robot) auth(ctx context.Context) (*apipb.AuthResult, error) {
	face := r.memory.Auth.Face
	fk := r.memory.Auth.FKey
	b := convert.StringToBytes(fmt.Sprintf(consts.AuthRequestParamTemplate, face, fk))
	request, err := r.basePostRequest(r.bs.App.AuthUrl, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.AuthResult](ctx, r.hc, request)
	if err != nil {
		return nil, err
	}
	if rst.Code != 0 {
		return nil, errors2.New(int(rst.Code), rst.Message, rst.Message)
	}
	return &rst.Data, nil
}

func (r *Robot) asset(ctx context.Context) (*apipb.AssetResult, error) {
	request, err := r.basePostRequest(r.bs.App.AssetUrl, bytes.NewBuffer(consts.EmptyRequestParam))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.AssetResult](ctx, r.hc, request)
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

func (r *Robot) createFish(ctx context.Context) (*apipb.FishCreateResult, error) {
	request, err := r.basePostRequest(r.bs.App.CreateFishUrl, bytes.NewBuffer(consts.EmptyRequestParam))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.FishCreateResult](ctx, r.hc, request)
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

func (r *Robot) fishList(ctx context.Context) (*apipb.FishListResult, error) {
	request, err := r.basePostRequest(r.bs.App.FishListUrl, bytes.NewBuffer(consts.EmptyRequestParam))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.FishListResult](ctx, r.hc, request)
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

func (r *Robot) fishSleep(ctx context.Context, fishId int64) (*apipb.FishSleepResult, error) {
	param := convert.StringToBytes(fmt.Sprintf(consts.OnlyFishIdParamTemplate, fishId))
	request, err := r.basePostRequest(r.bs.App.FishSleepUrl, bytes.NewBuffer(param))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.FishSleepResult](ctx, r.hc, request)
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

func (r *Robot) fishAlive(ctx context.Context, fishId int64) (*apipb.FishAliveResult, error) {
	param := convert.StringToBytes(fmt.Sprintf(consts.OnlyFishIdParamTemplate, fishId))
	request, err := r.basePostRequest(r.bs.App.FishAliveUrl, bytes.NewBuffer(param))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.FishAliveResult](ctx, r.hc, request)
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

func (r *Robot) fishRefining(ctx context.Context, fishId int64) (*apipb.FishRefiningResult, error) {
	param := convert.StringToBytes(fmt.Sprintf(consts.OnlyFishIdParamTemplate, fishId))
	request, err := r.basePostRequest(r.bs.App.FishRefiningUrl, bytes.NewBuffer(param))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.FishRefiningResult](ctx, r.hc, request)
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

func (r *Robot) poolRank(ctx context.Context, page, pageSize int32) (*apipb.FishPoolRankResult, error) {
	param := convert.StringToBytes(fmt.Sprintf(consts.OnlyPageParamTemplate, page, pageSize))
	request, err := r.basePostRequest(r.bs.App.PoolRankUrl, bytes.NewBuffer(param))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.FishPoolRankResult](ctx, r.hc, request)
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

func (r *Robot) expandParking(ctx context.Context) (*apipb.ExpandParkingResult, error) {
	request, err := r.basePostRequest(r.bs.App.ExpandParkingUrl, bytes.NewBuffer(consts.EmptyRequestParam))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.ExpandParkingResult](ctx, r.hc, request)
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

func (r *Robot) parkingList(ctx context.Context) (*apipb.ParkingListResult, error) {
	request, err := r.basePostRequest(r.bs.App.ParkingListUrl, bytes.NewBuffer(consts.EmptyRequestParam))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.ParkingListResult](ctx, r.hc, request)
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

func (r *Robot) marketList(ctx context.Context, page, pageSize int32) (*apipb.MarketListResult, error) {
	param := convert.StringToBytes(fmt.Sprintf(consts.OnlyPageParamTemplate, page, pageSize))
	request, err := r.basePostRequest(r.bs.App.MarketListUrl, bytes.NewBuffer(param))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.MarketListResult](ctx, r.hc, request)
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

func (r *Robot) marketDetail(ctx context.Context, productId int64) (*apipb.MarketDetailResult, error) {
	param := convert.StringToBytes(fmt.Sprintf(consts.OnlyProductIdParamTemplate, productId))
	request, err := r.basePostRequest(r.bs.App.MarketDetailUrl, bytes.NewBuffer(param))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.MarketDetailResult](ctx, r.hc, request)
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

func (r *Robot) marketSell(ctx context.Context, fishId, price int64) (*apipb.MarketSellResult, error) {
	param := convert.StringToBytes(fmt.Sprintf(consts.MarketSellParamTemplate, fishId, price))
	request, err := r.basePostRequest(r.bs.App.MarketSellUrl, bytes.NewBuffer(param))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.MarketSellResult](ctx, r.hc, request)
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

func (r *Robot) marketStopSell(ctx context.Context, productId int64) (*apipb.MarketStopSellResult, error) {
	param := convert.StringToBytes(fmt.Sprintf(consts.OnlyProductIdParamTemplate, productId))
	request, err := r.basePostRequest(r.bs.App.MarketStopSellUrl, bytes.NewBuffer(param))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.MarketStopSellResult](ctx, r.hc, request)
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

func (r *Robot) marketBuy(ctx context.Context, productId int64) (*apipb.MarketBuyResult, error) {
	param := convert.StringToBytes(fmt.Sprintf(consts.OnlyProductIdParamTemplate, productId))
	request, err := r.basePostRequest(r.bs.App.MarketBuyUrl, bytes.NewBuffer(param))
	if err != nil {
		return nil, err
	}
	rst, err := net_utils.RequestToStruct[apipb.MarketBuyResult](ctx, r.hc, request)
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
