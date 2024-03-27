package server

import (
	"context"
	"new-world-robot/internal/biz"
	"new-world-robot/internal/conf"
	"new-world-robot/internal/service"
	"time"
)

type TimerServer struct {
	started bool
	bs      *conf.Bootstrap
	ms      *service.ManagerService
	ticker  *time.Ticker
}

func NewTimerServer(ms *service.ManagerService, bs *conf.Bootstrap) *TimerServer {
	return &TimerServer{
		ms:      ms,
		bs:      bs,
		started: false,
	}
}

func (m *TimerServer) Start(ctx context.Context) error {
	m.started = true
	m.ticker = time.NewTicker(15 * time.Second)
	go m.run(ctx)
	return nil
}

func (m *TimerServer) Stop(ctx context.Context) error {
	m.started = false
	m.ticker.Stop()
	return nil
}

func (m *TimerServer) run(ctx context.Context) {
	for m.started {
		select {
		case <-m.ticker.C:
			robotsByGroup := m.ms.AllRobotsByGroup(int(m.bs.App.RunnerGroupSize))
			for _, robots := range robotsByGroup {
				go m.notifyRobots(ctx, robots)
			}
		}
	}
}

func (m *TimerServer) notifyRobots(ctx context.Context, robots []*biz.Robot) {
	for _, robot := range robots {
		robot.Work(ctx)
	}
}
