package business

import (
	v1 "github.com/88act/go-cms/server/api/v1"
	"github.com/88act/go-cms/server/middleware"
	"github.com/gin-gonic/gin"
)

type BusinessExampleRouter struct {
}

func (s *BusinessExampleRouter) InitSysBusinessExampleRouter(Router *gin.RouterGroup) {
	businessExampleRouter := Router.Group("businessExample").Use(middleware.OperationRecord())
	businessExampleRouterWithoutRecord := Router.Group("businessExample")
	var businessExampleApi = v1.ApiGroupApp.BusinessApiGroup.BusinessExampleApi
	{
		businessExampleRouter.POST("createSysBusinessExample", businessExampleApi.CreateBusinessExample)   // 新建BusinessExample
		businessExampleRouter.DELETE("deleteSysBusinessExample", businessExampleApi.DeleteBusinessExample) // 删除BusinessExample
		businessExampleRouter.PUT("updateSysBusinessExample", businessExampleApi.UpdateBusinessExample)    // 更新BusinessExample
	}
	{
		businessExampleRouterWithoutRecord.GET("findSysBusinessExample", businessExampleApi.FindBusinessExample)       // 根据ID获取BusinessExample
		businessExampleRouterWithoutRecord.GET("getSysBusinessExampleList", businessExampleApi.GetBusinessExampleList) // 获取BusinessExample列表
	}
}
