package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type MemLogsRouter struct {
}

// InitMemLogsRouter 初始化 MemLogs 路由信息
func (s *MemLogsRouter) InitMemLogsRouter(Router *gin.RouterGroup) {
	router := Router.Group("memLogs").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("memLogs")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.MemLogsApi
	{
		router.POST("createMemLogs", apiV1.CreateMemLogs)   // 新建MemLogs
		router.POST("deleteMemLogs", apiV1.DeleteMemLogs) // 删除MemLogs
		router.POST("deleteMemLogsByIds", apiV1.DeleteMemLogsByIds) // 批量删除MemLogs
		router.POST("updateMemLogs", apiV1.UpdateMemLogs)    // 更新MemLogs
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findMemLogs", apiV1.FindMemLogs)        // 根据ID获取MemLogs
		routerNoRecord.GET("getMemLogsList", apiV1.GetMemLogsList)  // 获取MemLogs列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel MemLogs列表
	}
}
