package business

import (
	v1 "go-cms/api/v1"
	"go-cms/middleware"

	"github.com/gin-gonic/gin"
)

type ActFavRouter struct {
}

// InitActFavRouter 初始化 ActFav 路由信息
func (s *ActFavRouter) InitActFavRouter(Router *gin.RouterGroup) {
	router := Router.Group("actFav").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("actFav")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.ActFavApi
	{
		router.POST("createActFav", apiV1.CreateActFav)             // 新建ActFav
		router.DELETE("deleteActFav", apiV1.DeleteActFav)           // 删除ActFav
		router.DELETE("deleteActFavByIds", apiV1.DeleteActFavByIds) // 批量删除ActFav
		router.PUT("updateActFav", apiV1.UpdateActFav)              // 更新ActFav
		router.POST("quickEdit", apiV1.QuickEdit)                   // 快速编辑
	}
	{
		routerNoRecord.GET("findActFav", apiV1.FindActFav)       // 根据ID获取ActFav
		routerNoRecord.GET("getActFavList", apiV1.GetActFavList) // 获取ActFav列表
		routerNoRecord.GET("excelList", apiV1.ExcelList)         // 分页导出excel ActFav列表
	}
}
