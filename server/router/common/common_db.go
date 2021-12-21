package commom

import (
	v1 "go-cms/api/v1"

	"github.com/gin-gonic/gin"
)

type CommonDbRouter struct {
}

// InitCommDbRouter 初始化 QuickEdit 路由信息
func (s *CommonDbRouter) InitCommonDbRouter(Router *gin.RouterGroup) {
	var dbApi = v1.ApiGroupApp.CommonApiGroup.CommonDbApi
	commonDbRouter := Router.Group("commDb")
	//commDbRouterWithoutRecord := Router.Group("commDb")
	{
		commonDbRouter.POST("quickEdit", dbApi.QuickEdit) // 更新 QuickEdit
		commonDbRouter.GET("test_db_get", dbApi.Test_db_get)
		commonDbRouter.GET("test_db", dbApi.Test_db)
	}

}
