package business

import (
	v1 "go-cms/api/v1"
	"go-cms/middleware"

	"github.com/gin-gonic/gin"
)

type SysApisRouter struct {
}

// InitSysApisRouter 初始化 SysApis 路由信息
func (s *SysApisRouter) InitSysApisRouter(Router *gin.RouterGroup) {
	router := Router.Group("sysApis").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("sysApis")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.SysApisApi
	{
		router.POST("createSysApis", apiV1.CreateSysApis)           // 新建SysApis
		router.POST("deleteSysApis", apiV1.DeleteSysApis)           // 删除SysApis
		router.POST("deleteSysApisByIds", apiV1.DeleteSysApisByIds) // 批量删除SysApis
		router.POST("updateSysApis", apiV1.UpdateSysApis)           // 更新SysApis
		router.POST("quickEdit", apiV1.QuickEdit)                   // 快速编辑
	}
	{
		routerNoRecord.GET("findSysApis", apiV1.FindSysApis)       // 根据ID获取SysApis
		routerNoRecord.GET("getSysApisList", apiV1.GetSysApisList) // 获取SysApis列表
		routerNoRecord.GET("excelList", apiV1.ExcelList)           // 分页导出excel SysApis列表
	}
}
