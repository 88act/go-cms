package global

import (
	"github.com/88act/go-cms/server/utils/timer"
	"github.com/songzhibin97/gkit/cache/local_cache"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/88act/go-cms/server/config"

	"github.com/go-redis/redis/v8"
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
