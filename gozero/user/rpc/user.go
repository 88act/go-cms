// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package main

import (
	"flag"
	"fmt"

	"datacenter/user/rpc/internal/config"
	"datacenter/user/rpc/internal/server"
	"datacenter/user/rpc/internal/svc"
	"datacenter/user/rpc/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()
	fmt.Println("rpc  ------1---------")
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewUserServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
