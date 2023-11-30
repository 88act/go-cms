package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type SysRoleMenuRouter struct {
}

// InitSysRoleMenuRouter 初始化 SysRoleMenu 路由信息
func (s *SysRoleMenuRouter) InitSysRoleMenuRouter(Router *gin.RouterGroup) {
	router := Router.Group("sysRoleMenu").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("sysRoleMenu")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.SysRoleMenuApi
	{
		router.POST("createSysRoleMenu", apiV1.CreateSysRoleMenu)   // 新建SysRoleMenu
		router.POST("deleteSysRoleMenu", apiV1.DeleteSysRoleMenu) // 删除SysRoleMenu
		router.POST("deleteSysRoleMenuByIds", apiV1.DeleteSysRoleMenuByIds) // 批量删除SysRoleMenu
		router.POST("updateSysRoleMenu", apiV1.UpdateSysRoleMenu)    // 更新SysRoleMenu
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findSysRoleMenu", apiV1.FindSysRoleMenu)        // 根据ID获取SysRoleMenu
		routerNoRecord.GET("getSysRoleMenuList", apiV1.GetSysRoleMenuList)  // 获取SysRoleMenu列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel SysRoleMenu列表
	}
}
