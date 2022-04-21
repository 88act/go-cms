package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type BbPiInstitutionRouter struct {
}

// InitBbPiInstitutionRouter 初始化 BbPiInstitution 路由信息
func (s *BbPiInstitutionRouter) InitBbPiInstitutionRouter(Router *gin.RouterGroup) {
	router := Router.Group("bbPiInstitution").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("bbPiInstitution")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.BbPiInstitutionApi
	{
		router.POST("createBbPiInstitution", apiV1.CreateBbPiInstitution)   // 新建BbPiInstitution
		router.DELETE("deleteBbPiInstitution", apiV1.DeleteBbPiInstitution) // 删除BbPiInstitution
		router.DELETE("deleteBbPiInstitutionByIds", apiV1.DeleteBbPiInstitutionByIds) // 批量删除BbPiInstitution
		router.PUT("updateBbPiInstitution", apiV1.UpdateBbPiInstitution)    // 更新BbPiInstitution
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findBbPiInstitution", apiV1.FindBbPiInstitution)        // 根据ID获取BbPiInstitution
		routerNoRecord.GET("getBbPiInstitutionList", apiV1.GetBbPiInstitutionList)  // 获取BbPiInstitution列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel BbPiInstitution列表
	}
}
