package scheduler

import (
	"github.com/buyandship/bns-golib/config"
	"github.com/buyandship/bso-order-cron/internal/jobs"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-co-op/gocron/v2"
	"sync"
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
			gocron.CronJob(
				config.GlobalAppConfig.GetString("create_shipment_order"),
				false,
			),
			gocron.NewTask(
				jobs.CreateShipmentJob,
			),
		)

		_, err = s.NewJob(
			gocron.CronJob(
				config.GlobalAppConfig.GetString("close_order"),
				false,
			),
			gocron.NewTask(
				jobs.CloseOrderJob,
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
	klog.Infof("scheduler is started...")
}

func (s *Scheduler) Shutdown() error {
	return s.scheduler.Shutdown()
}
