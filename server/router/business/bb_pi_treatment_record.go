package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type BbPiTreatmentRecordRouter struct {
}

// InitBbPiTreatmentRecordRouter 初始化 BbPiTreatmentRecord 路由信息
func (s *BbPiTreatmentRecordRouter) InitBbPiTreatmentRecordRouter(Router *gin.RouterGroup) {
	router := Router.Group("bbPiTreatmentRecord").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("bbPiTreatmentRecord")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.BbPiTreatmentRecordApi
	{
		router.POST("createBbPiTreatmentRecord", apiV1.CreateBbPiTreatmentRecord)   // 新建BbPiTreatmentRecord
		router.DELETE("deleteBbPiTreatmentRecord", apiV1.DeleteBbPiTreatmentRecord) // 删除BbPiTreatmentRecord
		router.DELETE("deleteBbPiTreatmentRecordByIds", apiV1.DeleteBbPiTreatmentRecordByIds) // 批量删除BbPiTreatmentRecord
		router.PUT("updateBbPiTreatmentRecord", apiV1.UpdateBbPiTreatmentRecord)    // 更新BbPiTreatmentRecord
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findBbPiTreatmentRecord", apiV1.FindBbPiTreatmentRecord)        // 根据ID获取BbPiTreatmentRecord
		routerNoRecord.GET("getBbPiTreatmentRecordList", apiV1.GetBbPiTreatmentRecordList)  // 获取BbPiTreatmentRecord列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel BbPiTreatmentRecord列表
	}
}
