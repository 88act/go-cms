package svc

import (
	"go-cms/app/act/cmd/rpc/internal/config"
	"go-cms/app/act/model"
	"os"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis

	ActActSev *model.ActActSev
}

func NewServiceContext(c config.Config) *ServiceContext {

	//sqlConn := sqlx.NewMysql(c.DB.DataSource)
	gormDB := GormMysql(c.DB.DataSource)
	// 新建 gorm 数据库 链接

	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		ActActSev: model.NewActActSev(gormDB, c.Cache),
	}
}

func GormMysql(dsn string) *gorm.DB {
	//dsn := "root:PXDN93VRKUm8TeE7@tcp(mysql:33069)/go-cms_usercenter?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	var ormLogger logger.Interface
	//if cfg.Debug {
	ormLogger = logger.Default.LogMode(logger.Info)
	//} else {
	//ormLogger = logger.Default
	//}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: ormLogger,
		//NamingStrategy: schema.NamingStrategy{
		//	TablePrefix: "ucenter_", // 表名前缀，`User` 对应的表名是 `tb_users`
		//},
	})
	if err != nil {
		logx.Errorf(err.Error())
		logx.Errorf("MySQL启动异常")
		os.Exit(0)
		return nil
	}
	logx.Info("act rpc MySQL启动 success")
	return db

}
