package initialize

import (
	"context"
	"os"

	"go-cms/global"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.LOG.Error("redis启动失败, err:", zap.Any("err", err))
		os.Exit(0)
	} else {
		global.LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.REDIS = client
	}
}
