package business

import (
	v1 "go-cms/api/v1"
	"go-cms/middleware"

	"github.com/gin-gonic/gin"
)

type BbPiTreatmentOrderRouter struct {
}

// InitBbPiTreatmentOrderRouter 初始化 BbPiTreatmentOrder 路由信息
func (s *BbPiTreatmentOrderRouter) InitBbPiTreatmentOrderRouter(Router *gin.RouterGroup) {
	router := Router.Group("bbPiTreatmentOrder").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("bbPiTreatmentOrder")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.BbPiTreatmentOrderApi
	{
		router.POST("createBbPiTreatmentOrder", apiV1.CreateBbPiTreatmentOrder)             // 新建BbPiTreatmentOrder
		router.DELETE("deleteBbPiTreatmentOrder", apiV1.DeleteBbPiTreatmentOrder)           // 删除BbPiTreatmentOrder
		router.DELETE("deleteBbPiTreatmentOrderByIds", apiV1.DeleteBbPiTreatmentOrderByIds) // 批量删除BbPiTreatmentOrder
		router.PUT("updateBbPiTreatmentOrder", apiV1.UpdateBbPiTreatmentOrder)              // 更新BbPiTreatmentOrder
		router.POST("quickEdit", apiV1.QuickEdit)                                           // 快速编辑
	}
	{
		routerNoRecord.GET("findBbPiTreatmentOrder", apiV1.FindBbPiTreatmentOrder)       // 根据ID获取BbPiTreatmentOrder
		routerNoRecord.GET("getBbPiTreatmentOrderList", apiV1.GetBbPiTreatmentOrderList) // 获取BbPiTreatmentOrder列表
		routerNoRecord.GET("excelList", apiV1.ExcelList)                                 // 分页导出excel BbPiTreatmentOrder列表
	}
}
