package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type MemAddressRouter struct {
}

// InitMemAddressRouter 初始化 MemAddress 路由信息
func (s *MemAddressRouter) InitMemAddressRouter(Router *gin.RouterGroup) {
	router := Router.Group("memAddress").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("memAddress")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.MemAddressApi
	{
		router.POST("createMemAddress", apiV1.CreateMemAddress)   // 新建MemAddress
		router.DELETE("deleteMemAddress", apiV1.DeleteMemAddress) // 删除MemAddress
		router.DELETE("deleteMemAddressByIds", apiV1.DeleteMemAddressByIds) // 批量删除MemAddress
		router.PUT("updateMemAddress", apiV1.UpdateMemAddress)    // 更新MemAddress
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findMemAddress", apiV1.FindMemAddress)        // 根据ID获取MemAddress
		routerNoRecord.GET("getMemAddressList", apiV1.GetMemAddressList)  // 获取MemAddress列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel MemAddress列表
	}
}
