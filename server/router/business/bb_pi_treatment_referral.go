package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type BbPiTreatmentReferralRouter struct {
}

// InitBbPiTreatmentReferralRouter 初始化 BbPiTreatmentReferral 路由信息
func (s *BbPiTreatmentReferralRouter) InitBbPiTreatmentReferralRouter(Router *gin.RouterGroup) {
	router := Router.Group("bbPiTreatmentReferral").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("bbPiTreatmentReferral")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.BbPiTreatmentReferralApi
	{
		router.POST("createBbPiTreatmentReferral", apiV1.CreateBbPiTreatmentReferral)   // 新建BbPiTreatmentReferral
		router.DELETE("deleteBbPiTreatmentReferral", apiV1.DeleteBbPiTreatmentReferral) // 删除BbPiTreatmentReferral
		router.DELETE("deleteBbPiTreatmentReferralByIds", apiV1.DeleteBbPiTreatmentReferralByIds) // 批量删除BbPiTreatmentReferral
		router.PUT("updateBbPiTreatmentReferral", apiV1.UpdateBbPiTreatmentReferral)    // 更新BbPiTreatmentReferral
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findBbPiTreatmentReferral", apiV1.FindBbPiTreatmentReferral)        // 根据ID获取BbPiTreatmentReferral
		routerNoRecord.GET("getBbPiTreatmentReferralList", apiV1.GetBbPiTreatmentReferralList)  // 获取BbPiTreatmentReferral列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel BbPiTreatmentReferral列表
	}
}
