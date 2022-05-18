package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Config struct {
	service.ServiceConf

	Redis redis.RedisConf
	//邮件发送 Kq : pub sub
	SendEmailConf kq.KqConf
	//短信发送 Kq : pub sub
	SendSmsConf kq.KqConf
	//image zip 图片压缩
	ImageZipConf kq.KqConf
}
