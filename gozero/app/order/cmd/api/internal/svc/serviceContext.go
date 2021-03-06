package svc

import (
	"go-cms/app/order/cmd/api/internal/config"
	"go-cms/app/order/cmd/rpc/order"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                config.Config
	OrderRpc              order.Order
	SetUidToCtxMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		OrderRpc: order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
	}
}
