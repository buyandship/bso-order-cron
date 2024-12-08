package main

import (
	"github.com/buyandship/bso-order-cron/internal/config"
	"github.com/buyandship/bso-order-cron/internal/scheduler"
	"github.com/cloudwego/kitex/pkg/klog"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// load config
	config.LoadConfig()

	// create a scheduler
	s := scheduler.NewScheduler()
	s.Start()
	defer func() {
		err := s.Shutdown()
		if err != nil {
			klog.Errorf("shutdown scheduler failed, err:%v", err)
		}
	}()

	select {
	case <-sigs:
		{
			klog.Infof("manually shutdown the scheduler")
		}
	}

	return
}
