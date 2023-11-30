package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type SysRoleRouter struct {
}

// InitSysRoleRouter 初始化 SysRole 路由信息
func (s *SysRoleRouter) InitSysRoleRouter(Router *gin.RouterGroup) {
	router := Router.Group("sysRole").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("sysRole")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.SysRoleApi
	{
		router.POST("createSysRole", apiV1.CreateSysRole)   // 新建SysRole
		router.POST("deleteSysRole", apiV1.DeleteSysRole) // 删除SysRole
		router.POST("deleteSysRoleByIds", apiV1.DeleteSysRoleByIds) // 批量删除SysRole
		router.POST("updateSysRole", apiV1.UpdateSysRole)    // 更新SysRole
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findSysRole", apiV1.FindSysRole)        // 根据ID获取SysRole
		routerNoRecord.GET("getSysRoleList", apiV1.GetSysRoleList)  // 获取SysRole列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel SysRole列表
	}
}
