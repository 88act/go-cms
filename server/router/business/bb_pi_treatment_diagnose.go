package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type BbPiTreatmentDiagnoseRouter struct {
}

// InitBbPiTreatmentDiagnoseRouter 初始化 BbPiTreatmentDiagnose 路由信息
func (s *BbPiTreatmentDiagnoseRouter) InitBbPiTreatmentDiagnoseRouter(Router *gin.RouterGroup) {
	router := Router.Group("bbPiTreatmentDiagnose").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("bbPiTreatmentDiagnose")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.BbPiTreatmentDiagnoseApi
	{
		router.POST("createBbPiTreatmentDiagnose", apiV1.CreateBbPiTreatmentDiagnose)   // 新建BbPiTreatmentDiagnose
		router.DELETE("deleteBbPiTreatmentDiagnose", apiV1.DeleteBbPiTreatmentDiagnose) // 删除BbPiTreatmentDiagnose
		router.DELETE("deleteBbPiTreatmentDiagnoseByIds", apiV1.DeleteBbPiTreatmentDiagnoseByIds) // 批量删除BbPiTreatmentDiagnose
		router.PUT("updateBbPiTreatmentDiagnose", apiV1.UpdateBbPiTreatmentDiagnose)    // 更新BbPiTreatmentDiagnose
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findBbPiTreatmentDiagnose", apiV1.FindBbPiTreatmentDiagnose)        // 根据ID获取BbPiTreatmentDiagnose
		routerNoRecord.GET("getBbPiTreatmentDiagnoseList", apiV1.GetBbPiTreatmentDiagnoseList)  // 获取BbPiTreatmentDiagnose列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel BbPiTreatmentDiagnose列表
	}
}
