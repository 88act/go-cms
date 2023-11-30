package global

import (
	"go-cms/utils/timer"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"go-cms/config"

	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	REDIS  *redis.Client
	CONFIG config.Server
	VP     *viper.Viper
	//GO_LOG    *oplogging.Logger
	LOG                 *zap.Logger
	Timer               timer.Timer = timer.NewTimerTask()
	Concurrency_Control             = &singleflight.Group{}
	BlackCache          local_cache.Cache
)
