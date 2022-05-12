package main

import (
	"context"
	"flag"
	"go-cms/app/mqueue/cmd/scheduler/internal/config"
	"go-cms/app/mqueue/cmd/scheduler/internal/logic"
	"go-cms/app/mqueue/cmd/scheduler/internal/svc"
	"os"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/mqueue.yaml", "Specify the config file")

func main() {
	flag.Parse()
	var c config.Config

	conf.MustLoad(*configFile, &c)

	logx.DisableStat()
	// log、prometheus、trace、metricsUrl.
	if err := c.SetUp(); err != nil {
		panic(err)
	}

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()
	mqueueScheduler := logic.NewCronScheduler(ctx, svcContext)
	mqueueScheduler.Register()

	if err := svcContext.Scheduler.Run(); err != nil {
		logx.Errorf("!!!MqueueSchedulerErr!!!  run err:%+v", err)
		os.Exit(1)
	}

}
