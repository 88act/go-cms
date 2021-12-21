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
	colCollectRouter := Router.Group("colCollect").Use(middleware.OperationRecord())
	colCollectRouterWithoutRecord := Router.Group("colCollect")
	var colCollectApi = v1.ApiGroupApp.BusinessApiGroup.ColCollectApi2
	{
		colCollectRouter.POST("createColCollect", colCollectApi.CreateColCollect)             // 新建ColCollect
		colCollectRouter.DELETE("deleteColCollect", colCollectApi.DeleteColCollect)           // 删除ColCollect
		colCollectRouter.DELETE("deleteColCollectByIds", colCollectApi.DeleteColCollectByIds) // 批量删除ColCollect
		colCollectRouter.PUT("updateColCollect", colCollectApi.UpdateColCollect)              // 更新ColCollect
		colCollectRouter.POST("quickEdit", colCollectApi.QuickEdit)                           // 快速编辑
		colCollectRouter.GET("startOrStopCollect", colCollectApi.StartOrStopCollect)
		colCollectRouterWithoutRecord.GET("excelList", colCollectApi.ExcelList) // 分页导出excel ColCollect列表
	}
	{
		colCollectRouterWithoutRecord.GET("findColCollect", colCollectApi.FindColCollect)       // 根据ID获取ColCollect
		colCollectRouterWithoutRecord.GET("getColCollectList", colCollectApi.GetColCollectList) // 获取ColCollect列表

	}
}
