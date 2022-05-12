package svc

import (
	"looklook/app/act/cmd/api/internal/config"
	"looklook/app/act/cmd/rpc/act"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                config.Config
	ActRpc                act.Act
	SetUidToCtxMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		ActRpc: act.NewAct(zrpc.MustNewClient(c.ActRpcConf)),
	}
}
