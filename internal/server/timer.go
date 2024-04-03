package server

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"new-world-robot/internal/biz"
	"new-world-robot/internal/conf"
	"new-world-robot/internal/service"
	"new-world-robot/pkg/convert"
	"os"
	"time"
)

type TimerServer struct {
	started          bool
	bs               *conf.Bootstrap
	ms               *service.ManagerService
	ticker           *time.Ticker
	statisticsTicker *time.Ticker
	testId           int64
}

func NewTimerServer(ms *service.ManagerService, bs *conf.Bootstrap) *TimerServer {
	return &TimerServer{
		ms:      ms,
		bs:      bs,
		started: false,
		testId:  time.Now().UnixMilli(),
	}
}

func (m *TimerServer) Start(ctx context.Context) error {
	m.started = true
	m.ticker = time.NewTicker(20 * time.Second)
	m.statisticsTicker = time.NewTicker(360 * time.Second)
	go m.run(ctx)
	return nil
}

func (m *TimerServer) Stop(ctx context.Context) error {
	m.started = false
	m.ticker.Stop()
	m.statisticsTicker.Stop()
	return nil
}

func (m *TimerServer) run(ctx context.Context) {
	var waterId int64 = 0
	for m.started {
		select {
		case <-m.ticker.C:
			robotsByGroup := m.ms.AllRobotsByGroup(int(m.bs.App.RunnerGroupSize))
			waterId++
			//log.Debugf("begin dispatch timer: %d", waterId)
			for _, robots := range robotsByGroup {
				r := robots
				go m.notifyRobots(waterId, r)
			}
		case <-m.statisticsTicker.C:
			log.Debug("reporting...")
			m.Report()
		}
	}
}

func (m *TimerServer) notifyRobots(waterId int64, robots []*biz.Robot) {
	for _, robot := range robots {
		r := robot
		go r.Work(waterId)
	}
}

func (m *TimerServer) Report() {
	cms, tms := m.ms.Statistics()
	filePath := fmt.Sprintf("%s/report_%d.log", m.bs.App.ReportPath, m.testId)
	reportTemplate := `
%s ---
request times:
%s
request total time:
%s
request avg time:
%s
------

	`
	requestTimes := ""
	for url, c := range cms {
		requestTimes = fmt.Sprintf("%s%s ===> %d\n", requestTimes, url, c)
	}
	requestTotalTime := ""
	for url, t := range tms {
		requestTotalTime = fmt.Sprintf("%s%s ===> %d ms\n", requestTotalTime, url, t)
	}
	requestAvgTime := ""
	for url, t := range tms {
		c := cms[url]
		requestAvgTime = fmt.Sprintf("%s%s ===> %f ms\n", requestAvgTime, url, float32(t)/float32(c))
	}
	now := time.Now().Format(time.DateTime)
	report := fmt.Sprintf(reportTemplate, now, requestTimes, requestTotalTime, requestAvgTime)

	if err := os.WriteFile(filePath, convert.StringToBytes(report), os.ModePerm); err != nil {
		log.Error(err)
		return
	}
	log.Debug("report completed")
}
