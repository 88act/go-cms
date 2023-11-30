package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type MemUserSafeRouter struct {
}

// InitMemUserSafeRouter 初始化 MemUserSafe 路由信息
func (s *MemUserSafeRouter) InitMemUserSafeRouter(Router *gin.RouterGroup) {
	router := Router.Group("memUserSafe").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("memUserSafe")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.MemUserSafeApi
	{
		router.POST("createMemUserSafe", apiV1.CreateMemUserSafe)   // 新建MemUserSafe
		router.POST("deleteMemUserSafe", apiV1.DeleteMemUserSafe) // 删除MemUserSafe
		router.POST("deleteMemUserSafeByIds", apiV1.DeleteMemUserSafeByIds) // 批量删除MemUserSafe
		router.POST("updateMemUserSafe", apiV1.UpdateMemUserSafe)    // 更新MemUserSafe
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findMemUserSafe", apiV1.FindMemUserSafe)        // 根据ID获取MemUserSafe
		routerNoRecord.GET("getMemUserSafeList", apiV1.GetMemUserSafeList)  // 获取MemUserSafe列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel MemUserSafe列表
	}
}
