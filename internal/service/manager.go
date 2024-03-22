package service

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"new-world-robot/internal/biz"
	"new-world-robot/internal/conf"
)

type ManagerService struct {
	bs     *conf.Bootstrap
	robots map[string]*biz.Robot
	hc     *http.Client
}

func NewManagerService(bs *conf.Bootstrap) *ManagerService {
	return &ManagerService{
		bs: bs,
		hc: &http.Client{},
	}
}

func (m *ManagerService) CreateRobot(ctx context.Context, count int) {
	if count <= 0 {
		return
	}
	for i := 0; i < count; i++ {
		id := uuid.New().String()
		robot := biz.NewRobot(id, m.hc)
		m.robots[id] = robot
	}
}

func (m *ManagerService) RemoveRobot(ctx context.Context, count int) {
	if count <= 0 {
		return
	}
	size := len(m.robots)
	if size < count {
		count = size
	}
	idx := 0
	keys := make([]string, count)
	for _, r := range m.robots {
		if idx <= count {
			break
		}
		keys[idx] = r.Id
		idx++
	}
	for _, key := range keys {
		m.robots[key].Stop()
		m.robots[key] = nil
		delete(m.robots, key)
	}
}

func (m *ManagerService) Size() int {
	return len(m.robots)
}
