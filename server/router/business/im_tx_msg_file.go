package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type ImTxMsgFileRouter struct {
}

// InitImTxMsgFileRouter 初始化 ImTxMsgFile 路由信息
func (s *ImTxMsgFileRouter) InitImTxMsgFileRouter(Router *gin.RouterGroup) {
	imTxMsgFileRouter := Router.Group("imTxMsgFile").Use(middleware.OperationRecord())
	imTxMsgFileRouterWithoutRecord := Router.Group("imTxMsgFile")
	var imTxMsgFileApi = v1.ApiGroupApp.BusinessApiGroup.ImTxMsgFileApi
	{
		imTxMsgFileRouter.POST("createImTxMsgFile", imTxMsgFileApi.CreateImTxMsgFile)   // 新建ImTxMsgFile
		imTxMsgFileRouter.DELETE("deleteImTxMsgFile", imTxMsgFileApi.DeleteImTxMsgFile) // 删除ImTxMsgFile
		imTxMsgFileRouter.DELETE("deleteImTxMsgFileByIds", imTxMsgFileApi.DeleteImTxMsgFileByIds) // 批量删除ImTxMsgFile
		imTxMsgFileRouter.PUT("updateImTxMsgFile", imTxMsgFileApi.UpdateImTxMsgFile)    // 更新ImTxMsgFile
	    imTxMsgFileRouter.POST("quickEdit", imTxMsgFileApi.QuickEdit)  // 快速编辑
		imTxMsgFileRouterWithoutRecord.GET("excelList", imTxMsgFileApi.ExcelList)  // 分页导出excel ImTxMsgFile列表
	}
	{
		imTxMsgFileRouterWithoutRecord.GET("findImTxMsgFile", imTxMsgFileApi.FindImTxMsgFile)        // 根据ID获取ImTxMsgFile
		imTxMsgFileRouterWithoutRecord.GET("getImTxMsgFileList", imTxMsgFileApi.GetImTxMsgFileList)  // 获取ImTxMsgFile列表
	    
	}
}
