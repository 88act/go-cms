package svc

import (
	"context"
	"go-cms/app/cms/cmd/api/internal/config"
	"go-cms/app/cms/model"
	"go-cms/common"
	"go-cms/common/mycache"
	"go-cms/common/myconfig"
	"os"
	"time"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config        config.Config
	RedisClient   *redis.Redis
	CmsArtSev     *model.CmsArtSev
	CmsCatSev     *model.CmsCatSev
	CmsDiscussSev *model.CmsDiscussSev
	//BasicRpc     basic.Basic
}

func NewServiceContext(c config.Config) *ServiceContext {
	logc.MustSetup(c.LogConf)
	logc.Info(context.Background(), c.Name, " 服务启动...", c.Host, " port=", c.Port)
	gormDB := GormMysql(c.Mode, c.DB)
	mycache.InitObj(c.Redis.Host, c.Redis.Pass, 0)
	myconfig.HttpRoot = c.LocalRes.BaseUrl

	// 新建 gorm 数据库 链接
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		CmsArtSev:     model.NewCmsArtSev(gormDB),
		CmsCatSev:     model.NewCmsCatSev(gormDB),
		CmsDiscussSev: model.NewCmsDiscussSev(gormDB),
		//BasicRpc:     basic.NewBasic(zrpc.MustNewClient(c.BasicRpcConf)),
	}

}

func GormMysql(mode string, dbConf common.DbConf) *gorm.DB {
	//dsn := "root:PXDN93VRKUm8TeE7@tcp(mysql:33069)/go-cms_usercenter?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	ormLogger := common.NewGormLogger(mode) // logger.Default.LogMode(logger.Info)
	dialector := mysql.New(mysql.Config{
		DSN:                       dbConf.DNS, // data source name
		DefaultStringSize:         256,        // string 类型字段的默认长度
		DisableDatetimePrecision:  true,       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,      // 根据当前 MySQL 版本自动配置
	})
	option := &gorm.Config{
		//禁用默认全局事务
		SkipDefaultTransaction: true,
		//开启预编译sql
		PrepareStmt: true,
		Logger:      ormLogger,
		//NamingStrategy: schema.NamingStrategy{
		//	TablePrefix: "ucenter_", // 表名前缀，`User` 对应的表名是 `tb_users`
		//},
	}

	db, err := gorm.Open(dialector, option)
	if err != nil {
		logx.Error("basicApi MySQL启动异常", err.Error())
		os.Exit(0)
		return nil
	}
	sqlDb, err := db.DB()
	if err != nil {
		logx.Error("basicApi MySQL启动异常2", err.Error())
		os.Exit(0)
		return nil
	}
	sqlDb.SetMaxOpenConns(dbConf.MaxOpenConns)
	sqlDb.SetMaxIdleConns(dbConf.MaxIdleConns)
	sqlDb.SetConnMaxIdleTime(time.Second * time.Duration(dbConf.ConnMaxIdleTime))
	sqlDb.SetConnMaxLifetime(time.Second * time.Duration(dbConf.ConnMaxLifetime))
	//sCtx.Db = db
	logx.Infof("%+v", sqlDb.Stats())
	logx.Info("cmsApi  MySQL启动 success")
	return db

}
