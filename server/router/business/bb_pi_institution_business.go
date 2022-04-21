package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type BbPiInstitutionBusinessRouter struct {
}

// InitBbPiInstitutionBusinessRouter 初始化 BbPiInstitutionBusiness 路由信息
func (s *BbPiInstitutionBusinessRouter) InitBbPiInstitutionBusinessRouter(Router *gin.RouterGroup) {
	router := Router.Group("bbPiInstitutionBusiness").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("bbPiInstitutionBusiness")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.BbPiInstitutionBusinessApi
	{
		router.POST("createBbPiInstitutionBusiness", apiV1.CreateBbPiInstitutionBusiness)   // 新建BbPiInstitutionBusiness
		router.DELETE("deleteBbPiInstitutionBusiness", apiV1.DeleteBbPiInstitutionBusiness) // 删除BbPiInstitutionBusiness
		router.DELETE("deleteBbPiInstitutionBusinessByIds", apiV1.DeleteBbPiInstitutionBusinessByIds) // 批量删除BbPiInstitutionBusiness
		router.PUT("updateBbPiInstitutionBusiness", apiV1.UpdateBbPiInstitutionBusiness)    // 更新BbPiInstitutionBusiness
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findBbPiInstitutionBusiness", apiV1.FindBbPiInstitutionBusiness)        // 根据ID获取BbPiInstitutionBusiness
		routerNoRecord.GET("getBbPiInstitutionBusinessList", apiV1.GetBbPiInstitutionBusinessList)  // 获取BbPiInstitutionBusiness列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel BbPiInstitutionBusiness列表
	}
}
