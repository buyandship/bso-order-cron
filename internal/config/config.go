package config

import (
	"github.com/buyandship/bns-golib/config"
	"github.com/buyandship/bns-golib/log"
	"github.com/buyandship/bns-golib/log/rollwriter"
	"github.com/cloudwego/kitex/pkg/klog"
	"io"
	"os"
)

func LoadConfig() {
	if err := config.LoadGlobalConfig("config.yaml"); err != nil {
		panic(err)
	}
	klog.SetLogger(log.NewKitexZapLogger())
	klog.SetLevel(klog.Level(config.GlobalAppConfig.Log.RpcLogConfig.Level))
	w1, err := rollwriter.NewRollWriter(config.GlobalAppConfig.Log.RpcLogConfig.File,
		rollwriter.WithMaxSize(100), rollwriter.WithMaxBackups(10))
	if err != nil {
		klog.Fatal(err)
	}

	// set output to both stdout and log file
	multiWriter := io.MultiWriter(os.Stdout, w1)
	klog.SetOutput(multiWriter)
}
