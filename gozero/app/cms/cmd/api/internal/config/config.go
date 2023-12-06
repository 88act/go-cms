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

	Redis    redis.RedisConf
	LocalRes struct {
		BaseUrl  string
		BasePath string
		Path     string
		PathUser string
	}
	LogConf logc.LogConf
}
