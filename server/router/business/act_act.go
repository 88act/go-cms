package business

import (
	v1 "go-cms/api/v1"
	"go-cms/middleware"

	"github.com/gin-gonic/gin"
)

type ActActRouter struct {
}

// InitActActRouter 初始化 ActAct 路由信息
func (s *ActActRouter) InitActActRouter(Router *gin.RouterGroup) {
	router := Router.Group("actAct").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("actAct")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.ActActApi
	{
		router.POST("createActAct", apiV1.CreateActAct)             // 新建ActAct
		router.DELETE("deleteActAct", apiV1.DeleteActAct)           // 删除ActAct
		router.DELETE("deleteActActByIds", apiV1.DeleteActActByIds) // 批量删除ActAct
		router.PUT("updateActAct", apiV1.UpdateActAct)              // 更新ActAct
		router.POST("quickEdit", apiV1.QuickEdit)                   // 快速编辑
	}
	{
		routerNoRecord.GET("findActAct", apiV1.FindActAct)       // 根据ID获取ActAct
		routerNoRecord.GET("getActActList", apiV1.GetActActList) // 获取ActAct列表
		routerNoRecord.GET("excelList", apiV1.ExcelList)         // 分页导出excel ActAct列表
	}
}
