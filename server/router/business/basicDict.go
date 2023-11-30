package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type BasicDictRouter struct {
}

// InitBasicDictRouter 初始化 BasicDict 路由信息
func (s *BasicDictRouter) InitBasicDictRouter(Router *gin.RouterGroup) {
	router := Router.Group("basicDict").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("basicDict")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.BasicDictApi
	{
		router.POST("createBasicDict", apiV1.CreateBasicDict)   // 新建BasicDict
		router.POST("deleteBasicDict", apiV1.DeleteBasicDict) // 删除BasicDict
		router.POST("deleteBasicDictByIds", apiV1.DeleteBasicDictByIds) // 批量删除BasicDict
		router.POST("updateBasicDict", apiV1.UpdateBasicDict)    // 更新BasicDict
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findBasicDict", apiV1.FindBasicDict)        // 根据ID获取BasicDict
		routerNoRecord.GET("getBasicDictList", apiV1.GetBasicDictList)  // 获取BasicDict列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel BasicDict列表
	}
}
