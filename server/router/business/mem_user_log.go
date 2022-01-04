package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type MemUserLogRouter struct {
}

// InitMemUserLogRouter 初始化 MemUserLog 路由信息
func (s *MemUserLogRouter) InitMemUserLogRouter(Router *gin.RouterGroup) {
	router := Router.Group("memUserLog").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("memUserLog")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.MemUserLogApi
	{
		router.POST("createMemUserLog", apiV1.CreateMemUserLog)   // 新建MemUserLog
		router.DELETE("deleteMemUserLog", apiV1.DeleteMemUserLog) // 删除MemUserLog
		router.DELETE("deleteMemUserLogByIds", apiV1.DeleteMemUserLogByIds) // 批量删除MemUserLog
		router.PUT("updateMemUserLog", apiV1.UpdateMemUserLog)    // 更新MemUserLog
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findMemUserLog", apiV1.FindMemUserLog)        // 根据ID获取MemUserLog
		routerNoRecord.GET("getMemUserLogList", apiV1.GetMemUserLogList)  // 获取MemUserLog列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel MemUserLog列表
	}
}
