package main

import (
	"flag"
	"fmt"

	"go-cms/app/order/cmd/rpc/internal/config"
	"go-cms/app/order/cmd/rpc/internal/mqs/listen"
	"go-cms/app/order/cmd/rpc/internal/server"
	"go-cms/app/order/cmd/rpc/internal/svc"
	"go-cms/app/order/cmd/rpc/pb"
	"go-cms/common/interceptor/rpcserver"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/order.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewOrderServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterOrderServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	//rpc log
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)
	fmt.Printf("Starting order  rpc server at %s...\n", c.ListenOn)

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	for _, mq := range listen.Mqs(c) {
		fmt.Printf("\n 启动消息队列监听  order mq listen... ")
		serviceGroup.Add(mq)
	}
	serviceGroup.Add(s)
	serviceGroup.Start()
}
