package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type SysLogsRouter struct {
}

// InitSysLogsRouter 初始化 SysLogs 路由信息
func (s *SysLogsRouter) InitSysLogsRouter(Router *gin.RouterGroup) {
	router := Router.Group("sysLogs").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("sysLogs")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.SysLogsApi
	{
		router.POST("createSysLogs", apiV1.CreateSysLogs)   // 新建SysLogs
		router.POST("deleteSysLogs", apiV1.DeleteSysLogs) // 删除SysLogs
		router.POST("deleteSysLogsByIds", apiV1.DeleteSysLogsByIds) // 批量删除SysLogs
		router.POST("updateSysLogs", apiV1.UpdateSysLogs)    // 更新SysLogs
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findSysLogs", apiV1.FindSysLogs)        // 根据ID获取SysLogs
		routerNoRecord.GET("getSysLogsList", apiV1.GetSysLogsList)  // 获取SysLogs列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel SysLogs列表
	}
}
