package service

import (
	"github.com/google/uuid"
	"math/rand"
	"net/http"
	"new-world-robot/internal/biz"
	"new-world-robot/internal/conf"
	"strconv"
	"time"
)

type ManagerService struct {
	bs     *conf.Bootstrap
	robots map[int64]*biz.Robot
	rd     *rand.Rand
	hc     *http.Client
}

func NewManagerService(bs *conf.Bootstrap) *ManagerService {
	return &ManagerService{
		bs:     bs,
		robots: map[int64]*biz.Robot{},
		rd:     rand.New(rand.NewSource(time.Now().UnixNano())),
		hc: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (m *ManagerService) AllRobotsByGroup(groupSize int) [][]*biz.Robot {
	rst := make([][]*biz.Robot, groupSize)
	gs := int64(groupSize)
	for k, robot := range m.robots {
		rst[k%gs] = append(rst[k%gs], robot)
	}
	return rst
}

func (m *ManagerService) CreateRobot(count int) {
	if count <= 0 {
		return
	}
	for i := 0; i < count; i++ {
		id := int64(uuid.New().ID())
		face := strconv.Itoa(int(m.bs.Robots.Auth.RobotStartUid) + i)
		robot := biz.NewRobot(id, m.hc, m.bs, m.rd, face, m.bs.Robots.Auth.FKey)
		m.robots[id] = robot

	}
}

func (m *ManagerService) RemoveRobot(count int) {
	if count <= 0 {
		return
	}
	size := len(m.robots)
	if size < count {
		count = size
	}
	idx := 0
	keys := make([]int64, count)
	for _, r := range m.robots {
		if idx <= count {
			break
		}
		keys[idx] = r.Id
		idx++
	}
	for _, key := range keys {
		m.robots[key] = nil
		delete(m.robots, key)
	}
}

func (m *ManagerService) Size() int {
	return len(m.robots)
}

func (m *ManagerService) Statistics() (map[string]int64, map[string]int64) {
	cms := map[string]int64{}
	tms := map[string]int64{}
	for _, robot := range m.robots {
		cm, tm := robot.Statistics()
		for _, url := range cm.Keys() {
			if v, ok := cms[url]; ok {
				get, _ := cm.Get(url)
				cms[url] = v + get
			} else {
				get, _ := cm.Get(url)
				cms[url] = get
			}
		}
		for _, url := range tm.Keys() {
			if v, ok := tms[url]; ok {
				get, _ := tm.Get(url)
				tms[url] = v + get
			} else {
				get, _ := tm.Get(url)
				tms[url] = get
			}
		}
	}
	return cms, tms
}
