package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type ImTxMsgRouter struct {
}

// InitImTxMsgRouter 初始化 ImTxMsg 路由信息
func (s *ImTxMsgRouter) InitImTxMsgRouter(Router *gin.RouterGroup) {
	router := Router.Group("imTxMsg").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("imTxMsg")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.ImTxMsgApi
	{
		router.POST("createImTxMsg", apiV1.CreateImTxMsg)   // 新建ImTxMsg
		router.DELETE("deleteImTxMsg", apiV1.DeleteImTxMsg) // 删除ImTxMsg
		router.DELETE("deleteImTxMsgByIds", apiV1.DeleteImTxMsgByIds) // 批量删除ImTxMsg
		router.PUT("updateImTxMsg", apiV1.UpdateImTxMsg)    // 更新ImTxMsg
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findImTxMsg", apiV1.FindImTxMsg)        // 根据ID获取ImTxMsg
		routerNoRecord.GET("getImTxMsgList", apiV1.GetImTxMsgList)  // 获取ImTxMsg列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel ImTxMsg列表
	}
}
