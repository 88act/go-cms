package business

import (
	v1 "go-cms/api/v1"
	"go-cms/middleware"

	"github.com/gin-gonic/gin"
)

type PayPaymentRouter struct {
}

// InitPayPaymentRouter 初始化 PayPayment 路由信息
func (s *PayPaymentRouter) InitPayPaymentRouter(Router *gin.RouterGroup) {
	router := Router.Group("payPayment").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("payPayment")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.PayPaymentApi
	{
		router.POST("createPayPayment", apiV1.CreatePayPayment)             // 新建PayPayment
		router.DELETE("deletePayPayment", apiV1.DeletePayPayment)           // 删除PayPayment
		router.DELETE("deletePayPaymentByIds", apiV1.DeletePayPaymentByIds) // 批量删除PayPayment
		router.PUT("updatePayPayment", apiV1.UpdatePayPayment)              // 更新PayPayment
		router.POST("quickEdit", apiV1.QuickEdit)                           // 快速编辑
	}
	{
		routerNoRecord.GET("findPayPayment", apiV1.FindPayPayment)       // 根据ID获取PayPayment
		routerNoRecord.GET("getPayPaymentList", apiV1.GetPayPaymentList) // 获取PayPayment列表
		routerNoRecord.GET("excelList", apiV1.ExcelList)                 // 分页导出excel PayPayment列表
	}
}
