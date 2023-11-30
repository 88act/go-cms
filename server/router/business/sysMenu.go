package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type SysMenuRouter struct {
}

// InitSysMenuRouter 初始化 SysMenu 路由信息
func (s *SysMenuRouter) InitSysMenuRouter(Router *gin.RouterGroup) {
	router := Router.Group("sysMenu").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("sysMenu")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.SysMenuApi
	{
		router.POST("createSysMenu", apiV1.CreateSysMenu)   // 新建SysMenu
		router.POST("deleteSysMenu", apiV1.DeleteSysMenu) // 删除SysMenu
		router.POST("deleteSysMenuByIds", apiV1.DeleteSysMenuByIds) // 批量删除SysMenu
		router.POST("updateSysMenu", apiV1.UpdateSysMenu)    // 更新SysMenu
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findSysMenu", apiV1.FindSysMenu)        // 根据ID获取SysMenu
		routerNoRecord.GET("getSysMenuList", apiV1.GetSysMenuList)  // 获取SysMenu列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel SysMenu列表
	}
}
