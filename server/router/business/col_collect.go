package business

import (
	v1 "go-cms/api/v1"
	"go-cms/middleware"

	"github.com/gin-gonic/gin"
)

type ColCollectRouter struct {
}

// InitColCollectRouter 初始化 ColCollect 路由信息
func (s *ColCollectRouter) InitColCollectRouter(Router *gin.RouterGroup) {
	router := Router.Group("colCollect").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("colCollect")
	//	//手工修改
	var colCollectApi = v1.ApiGroupApp.BusinessApiGroup.ColCollectApi2
	{
		router.POST("createColCollect", colCollectApi.CreateColCollect)             // 新建ColCollect
		router.DELETE("deleteColCollect", colCollectApi.DeleteColCollect)           // 删除ColCollect
		router.DELETE("deleteColCollectByIds", colCollectApi.DeleteColCollectByIds) // 批量删除ColCollect
		router.PUT("updateColCollect", colCollectApi.UpdateColCollect)              // 更新ColCollect
		router.POST("quickEdit", colCollectApi.QuickEdit)                           // 快速编辑
		routerNoRecord.GET("excelList", colCollectApi.ExcelList)
		//手工修改
		routerNoRecord.GET("startOrStopCollect", colCollectApi.StartOrStopCollect)
		// 分页导出excel ColCollect列表
	}
	{
		routerNoRecord.GET("findColCollect", colCollectApi.FindColCollect)       // 根据ID获取ColCollect
		routerNoRecord.GET("getColCollectList", colCollectApi.GetColCollectList) // 获取ColCollect列表

	}
}
