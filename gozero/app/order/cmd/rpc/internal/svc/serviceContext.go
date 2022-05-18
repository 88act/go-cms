package svc

import (
	"go-cms/app/act/cmd/rpc/act"
	"go-cms/app/order/cmd/rpc/internal/config"
	"go-cms/app/order/model"
	"go-cms/app/usercenter/cmd/rpc/usercenter"
	"os"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ServiceContext struct {
	Config                  config.Config
	AsynqClient             *asynq.Client
	RedisClient             *redis.Redis
	ActRpc                  act.Act
	UsercenterRpc           usercenter.Usercenter
	OrderOrderSev           *model.OrderOrderSev
	KqSendEmailClient       *kq.Pusher
	KqSendSmsClient         *kq.Pusher
	KqUpdatePayStatusClient *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {

	//sqlConn := sqlx.NewMysql(c.DB.DataSource)
	gormDB := GormMysql(c.DB.DataSource)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		AsynqClient:             newAsynqClient(c),
		OrderOrderSev:           model.NewOrderOrderSev(gormDB, c.Cache),
		ActRpc:                  act.NewAct(zrpc.MustNewClient(c.ActRpcConf)),
		UsercenterRpc:           usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		KqSendEmailClient:       kq.NewPusher(c.KqSendEmailConf.Brokers, c.KqSendEmailConf.Topic),
		KqSendSmsClient:         kq.NewPusher(c.KqSendSmsConf.Brokers, c.KqSendSmsConf.Topic),
		KqUpdatePayStatusClient: kq.NewPusher(c.KqSendSmsConf.Brokers, c.KqPaymentUpdatePayStatusConf.Topic),
	}
}

func GormMysql(dsn string) *gorm.DB {
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
	logx.Info("order rpc MySQL启动 success")
	return db

}
