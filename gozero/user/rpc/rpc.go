// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package main

import (
	"context"
	"datacenter/user/rpc/internal/config"
	"datacenter/user/rpc/internal/server"
	"datacenter/user/rpc/internal/svc"
	"datacenter/user/rpc/user"
	"flag"
	"fmt"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/zrpc"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/rpc.yaml", "the config file")
var limiter = rate.NewLimiter(rate.Limit(100), 100)

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewUserServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, srv)
	})
	defer s.Stop()
	s.AddUnaryInterceptors(authInterceptor)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

//拦截验证器
func authInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	if !limiter.Allow() {
		fmt.Println("限流了")
		return nil, nil
	}
	logx.Info("好玩")
	return handler(ctx, req)
}
