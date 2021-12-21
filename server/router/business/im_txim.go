package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type ImTximRouter struct {
}

// InitImTximRouter 初始化 ImTxim 路由信息
func (s *ImTximRouter) InitImTximRouter(Router *gin.RouterGroup) {
	imTximRouter := Router.Group("imTxim").Use(middleware.OperationRecord())
	imTximRouterWithoutRecord := Router.Group("imTxim")
	var imTximApi = v1.ApiGroupApp.BusinessApiGroup.ImTximApi
	{
		imTximRouter.POST("createImTxim", imTximApi.CreateImTxim)   // 新建ImTxim
		imTximRouter.DELETE("deleteImTxim", imTximApi.DeleteImTxim) // 删除ImTxim
		imTximRouter.DELETE("deleteImTximByIds", imTximApi.DeleteImTximByIds) // 批量删除ImTxim
		imTximRouter.PUT("updateImTxim", imTximApi.UpdateImTxim)    // 更新ImTxim
	    imTximRouter.POST("quickEdit", imTximApi.QuickEdit)  // 快速编辑
		imTximRouterWithoutRecord.GET("excelList", imTximApi.ExcelList)  // 分页导出excel ImTxim列表
	}
	{
		imTximRouterWithoutRecord.GET("findImTxim", imTximApi.FindImTxim)        // 根据ID获取ImTxim
		imTximRouterWithoutRecord.GET("getImTximList", imTximApi.GetImTximList)  // 获取ImTxim列表
	    
	}
}
