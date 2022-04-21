package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type BbPiPersonRouter struct {
}

// InitBbPiPersonRouter 初始化 BbPiPerson 路由信息
func (s *BbPiPersonRouter) InitBbPiPersonRouter(Router *gin.RouterGroup) {
	router := Router.Group("bbPiPerson").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("bbPiPerson")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.BbPiPersonApi
	{
		router.POST("createBbPiPerson", apiV1.CreateBbPiPerson)   // 新建BbPiPerson
		router.DELETE("deleteBbPiPerson", apiV1.DeleteBbPiPerson) // 删除BbPiPerson
		router.DELETE("deleteBbPiPersonByIds", apiV1.DeleteBbPiPersonByIds) // 批量删除BbPiPerson
		router.PUT("updateBbPiPerson", apiV1.UpdateBbPiPerson)    // 更新BbPiPerson
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findBbPiPerson", apiV1.FindBbPiPerson)        // 根据ID获取BbPiPerson
		routerNoRecord.GET("getBbPiPersonList", apiV1.GetBbPiPersonList)  // 获取BbPiPerson列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel BbPiPerson列表
	}
}
