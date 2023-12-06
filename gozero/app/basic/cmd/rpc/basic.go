package main

import (
	"flag"
	"fmt"

	"go-cms/app/basic/cmd/rpc/internal/config"
	"go-cms/app/basic/cmd/rpc/internal/server"
	"go-cms/app/basic/cmd/rpc/internal/svc"
	"go-cms/app/basic/cmd/rpc/pb"
	"go-cms/common/interceptor/rpcserver"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/basic.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewBasicServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterBasicServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	//rpc log
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)

	//defer s.Stop()

	fmt.Printf("启动 basic rpc server at %s...\n", c.ListenOn)
	//s.Start()

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	// for _, mq := range listen.Mqs(c) {
	// 	fmt.Printf("启动消息队列监听 basic mq  server ...\n")
	// 	serviceGroup.Add(mq)
	// }
	serviceGroup.Add(s)
	serviceGroup.Start()

}
