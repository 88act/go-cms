package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type BbPiContactsRouter struct {
}

// InitBbPiContactsRouter 初始化 BbPiContacts 路由信息
func (s *BbPiContactsRouter) InitBbPiContactsRouter(Router *gin.RouterGroup) {
	router := Router.Group("bbPiContacts").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("bbPiContacts")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.BbPiContactsApi
	{
		router.POST("createBbPiContacts", apiV1.CreateBbPiContacts)   // 新建BbPiContacts
		router.DELETE("deleteBbPiContacts", apiV1.DeleteBbPiContacts) // 删除BbPiContacts
		router.DELETE("deleteBbPiContactsByIds", apiV1.DeleteBbPiContactsByIds) // 批量删除BbPiContacts
		router.PUT("updateBbPiContacts", apiV1.UpdateBbPiContacts)    // 更新BbPiContacts
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findBbPiContacts", apiV1.FindBbPiContacts)        // 根据ID获取BbPiContacts
		routerNoRecord.GET("getBbPiContactsList", apiV1.GetBbPiContactsList)  // 获取BbPiContacts列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel BbPiContacts列表
	}
}
