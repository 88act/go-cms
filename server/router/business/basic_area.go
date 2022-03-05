package business

import (
	"github.com/gin-gonic/gin"
	"go-cms/api/v1"
	"go-cms/middleware"
)

type BasicAreaRouter struct {
}

// InitBasicAreaRouter 初始化 BasicArea 路由信息
func (s *BasicAreaRouter) InitBasicAreaRouter(Router *gin.RouterGroup) {
	router := Router.Group("basicArea").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("basicArea")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.BasicAreaApi
	{
		router.POST("createBasicArea", apiV1.CreateBasicArea)             // 新建BasicArea
		router.DELETE("deleteBasicArea", apiV1.DeleteBasicArea)           // 删除BasicArea
		router.DELETE("deleteBasicAreaByIds", apiV1.DeleteBasicAreaByIds) // 批量删除BasicArea
		router.PUT("updateBasicArea", apiV1.UpdateBasicArea)              // 更新BasicArea
		router.POST("quickEdit", apiV1.QuickEdit)                         // 快速编辑
	}
	{
		routerNoRecord.GET("findBasicArea", apiV1.FindBasicArea)       // 根据ID获取BasicArea
		routerNoRecord.GET("getBasicAreaList", apiV1.GetBasicAreaList) // 获取BasicArea列表
		routerNoRecord.GET("excelList", apiV1.ExcelList)               // 分页导出excel BasicArea列表
	}
}
