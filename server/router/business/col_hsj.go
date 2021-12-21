package business

import (
	"go-cms/middleware"

	v1 "go-cms/api/v1"

	"github.com/gin-gonic/gin"
)

type ColHsjRouter struct {
}

// InitColHsjRouter 初始化 ColHsj 路由信息
func (s *ColHsjRouter) InitColHsjRouter(Router *gin.RouterGroup) {
	colHsjRouter := Router.Group("colHsj").Use(middleware.OperationRecord())
	colHsjRouterWithoutRecord := Router.Group("colHsj")
	var colHsjApi = v1.ApiGroupApp.BusinessApiGroup.ColHsjApi
	{
		colHsjRouter.POST("createColHsj", colHsjApi.CreateColHsj)             // 新建ColHsj
		colHsjRouter.DELETE("deleteColHsj", colHsjApi.DeleteColHsj)           // 删除ColHsj
		colHsjRouter.DELETE("deleteColHsjByIds", colHsjApi.DeleteColHsjByIds) // 批量删除ColHsj
		colHsjRouter.PUT("updateColHsj", colHsjApi.UpdateColHsj)              // 更新ColHsj
		colHsjRouter.POST("quickEdit", colHsjApi.QuickEdit)                   // 快速编辑
	}
	{
		colHsjRouterWithoutRecord.GET("findColHsj", colHsjApi.FindColHsj)       // 根据ID获取ColHsj
		colHsjRouterWithoutRecord.GET("getColHsjList", colHsjApi.GetColHsjList) // 获取ColHsj列表
		colHsjRouterWithoutRecord.GET("excelList", colHsjApi.ExcelList)         // 分页导出excel ColHsj列表
	}
}
