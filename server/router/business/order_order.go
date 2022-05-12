package business

import (
	v1 "go-cms/api/v1"
	"go-cms/middleware"

	"github.com/gin-gonic/gin"
)

type OrderOrderRouter struct {
}

// InitOrderOrderRouter 初始化 OrderOrder 路由信息
func (s *OrderOrderRouter) InitOrderOrderRouter(Router *gin.RouterGroup) {
	router := Router.Group("orderOrder").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("orderOrder")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.OrderOrderApi
	{
		router.POST("createOrderOrder", apiV1.CreateOrderOrder)             // 新建OrderOrder
		router.DELETE("deleteOrderOrder", apiV1.DeleteOrderOrder)           // 删除OrderOrder
		router.DELETE("deleteOrderOrderByIds", apiV1.DeleteOrderOrderByIds) // 批量删除OrderOrder
		router.PUT("updateOrderOrder", apiV1.UpdateOrderOrder)              // 更新OrderOrder
		router.POST("quickEdit", apiV1.QuickEdit)                           // 快速编辑
	}
	{
		routerNoRecord.GET("findOrderOrder", apiV1.FindOrderOrder)       // 根据ID获取OrderOrder
		routerNoRecord.GET("getOrderOrderList", apiV1.GetOrderOrderList) // 获取OrderOrder列表
		routerNoRecord.GET("excelList", apiV1.ExcelList)                 // 分页导出excel OrderOrder列表
	}
}
