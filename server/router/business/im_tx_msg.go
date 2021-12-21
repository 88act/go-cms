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
	imTxMsgRouter := Router.Group("imTxMsg").Use(middleware.OperationRecord())
	imTxMsgRouterWithoutRecord := Router.Group("imTxMsg")
	var imTxMsgApi = v1.ApiGroupApp.BusinessApiGroup.ImTxMsgApi
	{
		imTxMsgRouter.POST("createImTxMsg", imTxMsgApi.CreateImTxMsg)   // 新建ImTxMsg
		imTxMsgRouter.DELETE("deleteImTxMsg", imTxMsgApi.DeleteImTxMsg) // 删除ImTxMsg
		imTxMsgRouter.DELETE("deleteImTxMsgByIds", imTxMsgApi.DeleteImTxMsgByIds) // 批量删除ImTxMsg
		imTxMsgRouter.PUT("updateImTxMsg", imTxMsgApi.UpdateImTxMsg)    // 更新ImTxMsg
	    imTxMsgRouter.POST("quickEdit", imTxMsgApi.QuickEdit)  // 快速编辑
		imTxMsgRouterWithoutRecord.GET("excelList", imTxMsgApi.ExcelList)  // 分页导出excel ImTxMsg列表
	}
	{
		imTxMsgRouterWithoutRecord.GET("findImTxMsg", imTxMsgApi.FindImTxMsg)        // 根据ID获取ImTxMsg
		imTxMsgRouterWithoutRecord.GET("getImTxMsgList", imTxMsgApi.GetImTxMsgList)  // 获取ImTxMsg列表
	    
	}
}
