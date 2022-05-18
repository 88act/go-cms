package config

import (
	"go-cms/common/kqueue"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DataSource string
	}
	Cache                        cache.CacheConf
	ActRpcConf                   zrpc.RpcClientConf
	UsercenterRpcConf            zrpc.RpcClientConf
	KqPaymentUpdatePayStatusConf kqueue.KqConfig
	KqSendEmailConf              kqueue.KqConfig
	KqSendSmsConf                kqueue.KqConfig
}
