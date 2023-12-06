package config

import (
	"go-cms/common"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	DB common.DbConf
	// MaxOpenConns             int    `json:",default=10"`
	// MaxIdleConns             int    `json:",default=5"`
	// ConnMaxIdleTime          int    `json:",default=60"`
	// ConnMaxLifetime          int    `json:",default=60"`
	// SlowThresholdMillisecond int64  `json:",default=1000"`

	Redis      redis.RedisConf
	WxMiniConf WxMiniConf
	LocalRes   struct {
		BaseUrl  string
		BasePath string
		Path     string
		PathUser string
	}
	LogConf logc.LogConf

	//UsercenterRpcConf zrpc.RpcClientConf
	//BasicRpcConf      zrpc.RpcClientConf
	//Cache             cache.CacheConf
}
