package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type ColKeyFieldRouter struct {
}

// InitColKeyFieldRouter 初始化 ColKeyField 路由信息
func (s *ColKeyFieldRouter) InitColKeyFieldRouter(Router *gin.RouterGroup) {
	colKeyFieldRouter := Router.Group("colKeyField").Use(middleware.OperationRecord())
	colKeyFieldRouterWithoutRecord := Router.Group("colKeyField")
	var colKeyFieldApi = v1.ApiGroupApp.BusinessApiGroup.ColKeyFieldApi
	{
		colKeyFieldRouter.POST("createColKeyField", colKeyFieldApi.CreateColKeyField)   // 新建ColKeyField
		colKeyFieldRouter.DELETE("deleteColKeyField", colKeyFieldApi.DeleteColKeyField) // 删除ColKeyField
		colKeyFieldRouter.DELETE("deleteColKeyFieldByIds", colKeyFieldApi.DeleteColKeyFieldByIds) // 批量删除ColKeyField
		colKeyFieldRouter.PUT("updateColKeyField", colKeyFieldApi.UpdateColKeyField)    // 更新ColKeyField
	    colKeyFieldRouter.POST("quickEdit", colKeyFieldApi.QuickEdit)  // 快速编辑
		colKeyFieldRouterWithoutRecord.GET("excelList", colKeyFieldApi.ExcelList)  // 分页导出excel ColKeyField列表
	}
	{
		colKeyFieldRouterWithoutRecord.GET("findColKeyField", colKeyFieldApi.FindColKeyField)        // 根据ID获取ColKeyField
		colKeyFieldRouterWithoutRecord.GET("getColKeyFieldList", colKeyFieldApi.GetColKeyFieldList)  // 获取ColKeyField列表
	    
	}
}
