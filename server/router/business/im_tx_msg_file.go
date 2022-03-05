package business

import (
	"github.com/gin-gonic/gin"
	"go-cms/api/v1"
	"go-cms/middleware"
)

type ImTxMsgFileRouter struct {
}

// InitImTxMsgFileRouter 初始化 ImTxMsgFile 路由信息
func (s *ImTxMsgFileRouter) InitImTxMsgFileRouter(Router *gin.RouterGroup) {
	router := Router.Group("imTxMsgFile").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("imTxMsgFile")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.ImTxMsgFileApi
	{
		router.POST("createImTxMsgFile", apiV1.CreateImTxMsgFile)             // 新建ImTxMsgFile
		router.DELETE("deleteImTxMsgFile", apiV1.DeleteImTxMsgFile)           // 删除ImTxMsgFile
		router.DELETE("deleteImTxMsgFileByIds", apiV1.DeleteImTxMsgFileByIds) // 批量删除ImTxMsgFile
		router.PUT("updateImTxMsgFile", apiV1.UpdateImTxMsgFile)              // 更新ImTxMsgFile
		router.POST("quickEdit", apiV1.QuickEdit)                             // 快速编辑
	}
	{
		routerNoRecord.GET("findImTxMsgFile", apiV1.FindImTxMsgFile)       // 根据ID获取ImTxMsgFile
		routerNoRecord.GET("getImTxMsgFileList", apiV1.GetImTxMsgFileList) // 获取ImTxMsgFile列表
		routerNoRecord.GET("excelList", apiV1.ExcelList)                   // 分页导出excel ImTxMsgFile列表
	}
}
