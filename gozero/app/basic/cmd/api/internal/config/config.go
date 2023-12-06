package config

import (
	"go-cms/common"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	DB common.DbConf

	Redis redis.RedisConf

	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	LocalRes struct {
		BaseUrl  string
		BasePath string
		Path     string
		PathUser string
	}
	LogConf logc.LogConf
	//Cache   cache.CacheConf
	//BasicRpcConf zrpc.RpcClientConf
}
