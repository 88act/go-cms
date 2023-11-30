package commom

import (
	v1 "go-cms/api/v1"
	"go-cms/middleware"

	"github.com/gin-gonic/gin"
)

type CommonDbRouter struct {
}

// InitCommDbRouter 初始化 QuickEdit 路由信息
func (s *CommonDbRouter) InitCommonDbRouter(Router *gin.RouterGroup) {
	var dbApi = v1.ApiGroupApp.CommonApiGroup.CommonDbApi
	router := Router.Group("commDb").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("commDb")
	{
		router.POST("quickEdit", dbApi.QuickEdit) // 更新 QuickEdit
		router.POST("updatePidSort", dbApi.UpdatePidSort)
	}
	{
		routerNoRecord.POST("pidData", dbApi.PidData)         // 获取 PidData
		routerNoRecord.POST("pidTreeData", dbApi.PidTreeData) // 获取 PidTreeData
		routerNoRecord.POST("getDict", dbApi.GetDict)
		routerNoRecord.POST("getDict2", dbApi.GetDict2)
		routerNoRecord.POST("getDictTree", dbApi.GetDictTree)
	}

}

// InitCommDbRouter 初始化 QuickEdit 路由信息
func (s *CommonDbRouter) InitCommonDbRouterPub(Router *gin.RouterGroup) {
	//var dbApi = v1.ApiGroupApp.CommonApiGroup.CommonDbApi
	//router := Router.Group("commDb").Use(middleware.OperationRecord())

	{

		//routerNoRecord.GET("test_db", dbApi.Test_db)
	}

}
