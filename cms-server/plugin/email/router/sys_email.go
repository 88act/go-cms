package router

import (
	"github.com/88act/go-cms/server/middleware"
	"github.com/88act/go-cms/server/plugin/email/api"
	"github.com/gin-gonic/gin"
)

type EmailRouter struct {
}

func (s *EmailRouter) InitEmailRouter(Router *gin.RouterGroup) {
	emailRouter := Router.Use(middleware.OperationRecord())
	var EmailApi = api.ApiGroupApp.EmailApi.EmailTest
	var SendEmail = api.ApiGroupApp.EmailApi.SendEmail
	{
		emailRouter.POST("emailTest", EmailApi)  // 发送测试邮件
		emailRouter.POST("sendEmail", SendEmail) // 发送邮件
	}
}
