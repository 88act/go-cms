package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DataSource string
	}
	Cache cache.CacheConf

	//Redis redis.RedisConf
	//邮件发送 Kq : pub sub
	SendEmailConf kq.KqConf
	//短信发送 Kq : pub sub
	SendSmsConf kq.KqConf
	//image zip 图片压缩
	ImageZipConf kq.KqConf
}
