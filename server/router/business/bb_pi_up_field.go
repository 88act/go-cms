package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type BbPiUpFieldRouter struct {
}

// InitBbPiUpFieldRouter 初始化 BbPiUpField 路由信息
func (s *BbPiUpFieldRouter) InitBbPiUpFieldRouter(Router *gin.RouterGroup) {
	router := Router.Group("bbPiUpField").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("bbPiUpField")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.BbPiUpFieldApi
	{
		router.POST("createBbPiUpField", apiV1.CreateBbPiUpField)   // 新建BbPiUpField
		router.DELETE("deleteBbPiUpField", apiV1.DeleteBbPiUpField) // 删除BbPiUpField
		router.DELETE("deleteBbPiUpFieldByIds", apiV1.DeleteBbPiUpFieldByIds) // 批量删除BbPiUpField
		router.PUT("updateBbPiUpField", apiV1.UpdateBbPiUpField)    // 更新BbPiUpField
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findBbPiUpField", apiV1.FindBbPiUpField)        // 根据ID获取BbPiUpField
		routerNoRecord.GET("getBbPiUpFieldList", apiV1.GetBbPiUpFieldList)  // 获取BbPiUpField列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel BbPiUpField列表
	}
}
