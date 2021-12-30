package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type MemUserRouter struct {
}

// InitMemUserRouter 初始化 MemUser 路由信息
func (s *MemUserRouter) InitMemUserRouter(Router *gin.RouterGroup) {
	router := Router.Group("memUser").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("memUser")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.MemUserApi
	{
		router.POST("createMemUser", apiV1.CreateMemUser)   // 新建MemUser
		router.DELETE("deleteMemUser", apiV1.DeleteMemUser) // 删除MemUser
		router.DELETE("deleteMemUserByIds", apiV1.DeleteMemUserByIds) // 批量删除MemUser
		router.PUT("updateMemUser", apiV1.UpdateMemUser)    // 更新MemUser
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findMemUser", apiV1.FindMemUser)        // 根据ID获取MemUser
		routerNoRecord.GET("getMemUserList", apiV1.GetMemUserList)  // 获取MemUser列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel MemUser列表
	}
}
