package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type BbPiServicePointRouter struct {
}

// InitBbPiServicePointRouter 初始化 BbPiServicePoint 路由信息
func (s *BbPiServicePointRouter) InitBbPiServicePointRouter(Router *gin.RouterGroup) {
	router := Router.Group("bbPiServicePoint").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("bbPiServicePoint")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.BbPiServicePointApi
	{
		router.POST("createBbPiServicePoint", apiV1.CreateBbPiServicePoint)   // 新建BbPiServicePoint
		router.DELETE("deleteBbPiServicePoint", apiV1.DeleteBbPiServicePoint) // 删除BbPiServicePoint
		router.DELETE("deleteBbPiServicePointByIds", apiV1.DeleteBbPiServicePointByIds) // 批量删除BbPiServicePoint
		router.PUT("updateBbPiServicePoint", apiV1.UpdateBbPiServicePoint)    // 更新BbPiServicePoint
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findBbPiServicePoint", apiV1.FindBbPiServicePoint)        // 根据ID获取BbPiServicePoint
		routerNoRecord.GET("getBbPiServicePointList", apiV1.GetBbPiServicePointList)  // 获取BbPiServicePoint列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel BbPiServicePoint列表
	}
}
