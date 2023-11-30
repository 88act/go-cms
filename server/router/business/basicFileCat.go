package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type BasicFileCatRouter struct {
}

// InitBasicFileCatRouter 初始化 BasicFileCat 路由信息
func (s *BasicFileCatRouter) InitBasicFileCatRouter(Router *gin.RouterGroup) {
	router := Router.Group("basicFileCat").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("basicFileCat")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.BasicFileCatApi
	{
		router.POST("createBasicFileCat", apiV1.CreateBasicFileCat)   // 新建BasicFileCat
		router.POST("deleteBasicFileCat", apiV1.DeleteBasicFileCat) // 删除BasicFileCat
		router.POST("deleteBasicFileCatByIds", apiV1.DeleteBasicFileCatByIds) // 批量删除BasicFileCat
		router.POST("updateBasicFileCat", apiV1.UpdateBasicFileCat)    // 更新BasicFileCat
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findBasicFileCat", apiV1.FindBasicFileCat)        // 根据ID获取BasicFileCat
		routerNoRecord.GET("getBasicFileCatList", apiV1.GetBasicFileCatList)  // 获取BasicFileCat列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel BasicFileCat列表
	}
}
