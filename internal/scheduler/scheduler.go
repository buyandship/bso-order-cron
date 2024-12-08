package scheduler

import (
	"github.com/buyandship/bso-order-cron/internal/jobs"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-co-op/gocron/v2"
	"sync"
	"time"
)

type Scheduler struct {
	scheduler gocron.Scheduler
}

var once sync.Once

var instance *Scheduler

func NewScheduler() *Scheduler {
	once.Do(func() {
		s, err := gocron.NewScheduler()
		if err != nil {
			klog.Fatal(err)
		}

		// add jobs
		_, err = s.NewJob(
			gocron.DurationJob(
				5*time.Minute,
			),
			gocron.NewTask(
				jobs.CreateShipment,
			),
		)
		instance = &Scheduler{
			scheduler: s,
		}
	})
	return instance
}

func (s *Scheduler) Start() {
	s.scheduler.Start()
}

func (s *Scheduler) Shutdown() error {
	return s.scheduler.Shutdown()
}
