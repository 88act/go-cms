package business

import (
	v1 "go-cms/api/v1"
	"go-cms/middleware"

	"github.com/gin-gonic/gin"
)

type ImTximRouter struct {
}

// InitImTximRouter 初始化 ImTxim 路由信息
func (s *ImTximRouter) InitImTximRouter(Router *gin.RouterGroup) {
	router := Router.Group("imTxim").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("imTxim")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.ImTximApi2
	{
		router.POST("createImTxim", apiV1.CreateImTxim)             // 新建ImTxim
		router.DELETE("deleteImTxim", apiV1.DeleteImTxim)           // 删除ImTxim
		router.DELETE("deleteImTximByIds", apiV1.DeleteImTximByIds) // 批量删除ImTxim
		router.PUT("updateImTxim", apiV1.UpdateImTxim)              // 更新ImTxim
		router.POST("quickEdit", apiV1.QuickEdit)                   // 快速编辑
	}
	{
		routerNoRecord.GET("findImTxim", apiV1.FindImTxim)       // 根据ID获取ImTxim
		routerNoRecord.GET("getImTximList", apiV1.GetImTximList) // 获取ImTxim列表
		routerNoRecord.GET("excelList", apiV1.ExcelList)         // 分页导出excel ImTxim列表
		//手工修改
		routerNoRecord.GET("startOrStopCollect", apiV1.StartOrStopCollect)
	}
}
