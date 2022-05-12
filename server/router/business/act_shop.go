package business

import (
	v1 "go-cms/api/v1"
	"go-cms/middleware"

	"github.com/gin-gonic/gin"
)

type ActShopRouter struct {
}

// InitActShopRouter 初始化 ActShop 路由信息
func (s *ActShopRouter) InitActShopRouter(Router *gin.RouterGroup) {
	router := Router.Group("actShop").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("actShop")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.ActShopApi
	{
		router.POST("createActShop", apiV1.CreateActShop)             // 新建ActShop
		router.DELETE("deleteActShop", apiV1.DeleteActShop)           // 删除ActShop
		router.DELETE("deleteActShopByIds", apiV1.DeleteActShopByIds) // 批量删除ActShop
		router.PUT("updateActShop", apiV1.UpdateActShop)              // 更新ActShop
		router.POST("quickEdit", apiV1.QuickEdit)                     // 快速编辑
	}
	{
		routerNoRecord.GET("findActShop", apiV1.FindActShop)       // 根据ID获取ActShop
		routerNoRecord.GET("getActShopList", apiV1.GetActShopList) // 获取ActShop列表
		routerNoRecord.GET("excelList", apiV1.ExcelList)           // 分页导出excel ActShop列表
	}
}
