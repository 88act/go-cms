package main

import (
	"flag"
	"fmt"
	"go-cms/app/basic/cmd/api/internal/config"
	"go-cms/app/basic/cmd/api/internal/handler"

	"go-cms/app/basic/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/basic.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("启动 basic api server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
