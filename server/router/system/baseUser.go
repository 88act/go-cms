package system

import (
	v1 "go-cms/api/v1"
	"go-cms/middleware"

	"github.com/gin-gonic/gin"
)

type BaseRouter struct {
}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base").Use(middleware.OperationRecord())
	baseRouterWithoutRecord := Router.Group("base")

	var baseUserApi = v1.ApiGroupApp.SystemApiGroup.BaseUserApi
	{
		baseRouter.POST("login", baseUserApi.Login)
	}
	{
		baseRouterWithoutRecord.POST("captcha", baseUserApi.Captcha)
	}
	return baseRouter
}

func (s *BaseRouter) InitBasePriRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base").Use(middleware.OperationRecord())
	baseRouterWithoutRecord := Router.Group("base")
	var systemApi = v1.ApiGroupApp.SystemApiGroup.SystemApi
	var baseUserApi = v1.ApiGroupApp.SystemApiGroup.BaseUserApi
	{

		baseRouter.GET("getUserInfo", baseUserApi.GetUserInfo)
		baseRouter.POST("changePwd", baseUserApi.ChangePwd)
		baseRouter.POST("changePwdMy", baseUserApi.ChangePwdMy)

	}
	{
		baseRouterWithoutRecord.GET("getServerInfo", systemApi.GetServerInfo)
		baseRouterWithoutRecord.POST("showQrcode", baseUserApi.ShowQrcode)
	}
	return baseRouter
}
