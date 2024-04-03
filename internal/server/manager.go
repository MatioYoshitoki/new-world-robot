package server

import (
	"context"
	"new-world-robot/internal/conf"
	"new-world-robot/internal/service"
	"time"
)

type ManagerServer struct {
	started bool
	bs      *conf.Bootstrap
	ms      *service.ManagerService
	ticker  *time.Ticker
}

func NewManagerServer(ms *service.ManagerService, bs *conf.Bootstrap) *ManagerServer {
	return &ManagerServer{
		ms:      ms,
		bs:      bs,
		started: false,
	}
}

func (m *ManagerServer) Start(ctx context.Context) error {
	m.started = true
	m.ticker = time.NewTicker(1 * time.Second)
	go m.run()
	return nil
}

func (m *ManagerServer) Stop(ctx context.Context) error {
	m.started = false
	m.ticker.Stop()
	return nil
}

func (m *ManagerServer) run() {
	for m.started {
		select {
		case <-m.ticker.C:
			size := m.ms.Size()
			targetCount := int(m.bs.Robots.Auth.RobotCount)
			if targetCount > size {
				m.ms.CreateRobot(targetCount - size)
			}
		}
	}
}
