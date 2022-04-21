package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type BbPiDeviceRouter struct {
}

// InitBbPiDeviceRouter 初始化 BbPiDevice 路由信息
func (s *BbPiDeviceRouter) InitBbPiDeviceRouter(Router *gin.RouterGroup) {
	router := Router.Group("bbPiDevice").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("bbPiDevice")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.BbPiDeviceApi
	{
		router.POST("createBbPiDevice", apiV1.CreateBbPiDevice)   // 新建BbPiDevice
		router.DELETE("deleteBbPiDevice", apiV1.DeleteBbPiDevice) // 删除BbPiDevice
		router.DELETE("deleteBbPiDeviceByIds", apiV1.DeleteBbPiDeviceByIds) // 批量删除BbPiDevice
		router.PUT("updateBbPiDevice", apiV1.UpdateBbPiDevice)    // 更新BbPiDevice
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findBbPiDevice", apiV1.FindBbPiDevice)        // 根据ID获取BbPiDevice
		routerNoRecord.GET("getBbPiDeviceList", apiV1.GetBbPiDeviceList)  // 获取BbPiDevice列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel BbPiDevice列表
	}
}
