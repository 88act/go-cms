package svc

import (
	"go-cms/app/basic/cmd/mq/internal/config"
)

type ServiceContext struct {
	Config config.Config

	//	OrderRpc      order.Order
	///	UsercenterRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		//OrderRpc:      order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		//UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
