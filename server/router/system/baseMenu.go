package system

import (
	v1 "go-cms/api/v1"
	"go-cms/middleware"

	"github.com/gin-gonic/gin"
)

type MenuRouter struct {
}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	menuRouter := Router.Group("base").Use(middleware.OperationRecord())
	menuRouterWithoutRecord := Router.Group("base")
	var baseMenuApi = v1.ApiGroupApp.SystemApiGroup.BaseMenuApi
	{
		menuRouterWithoutRecord.GET("getMenu", baseMenuApi.GetMenu) // 获取菜单树

	}
	return menuRouter
}
