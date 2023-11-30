package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type MemDepartRouter struct {
}

// InitMemDepartRouter 初始化 MemDepart 路由信息
func (s *MemDepartRouter) InitMemDepartRouter(Router *gin.RouterGroup) {
	router := Router.Group("memDepart").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("memDepart")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.MemDepartApi
	{
		router.POST("createMemDepart", apiV1.CreateMemDepart)   // 新建MemDepart
		router.POST("deleteMemDepart", apiV1.DeleteMemDepart) // 删除MemDepart
		router.POST("deleteMemDepartByIds", apiV1.DeleteMemDepartByIds) // 批量删除MemDepart
		router.POST("updateMemDepart", apiV1.UpdateMemDepart)    // 更新MemDepart
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findMemDepart", apiV1.FindMemDepart)        // 根据ID获取MemDepart
		routerNoRecord.GET("getMemDepartList", apiV1.GetMemDepartList)  // 获取MemDepart列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel MemDepart列表
	}
}
