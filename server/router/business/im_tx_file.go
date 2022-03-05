package business

import (
	"github.com/gin-gonic/gin"
	"go-cms/api/v1"
	"go-cms/middleware"
)

type ImTxFileRouter struct {
}

// InitImTxFileRouter 初始化 ImTxFile 路由信息
func (s *ImTxFileRouter) InitImTxFileRouter(Router *gin.RouterGroup) {
	router := Router.Group("imTxFile").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("imTxFile")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.ImTxFileApi
	{
		router.POST("createImTxFile", apiV1.CreateImTxFile)             // 新建ImTxFile
		router.DELETE("deleteImTxFile", apiV1.DeleteImTxFile)           // 删除ImTxFile
		router.DELETE("deleteImTxFileByIds", apiV1.DeleteImTxFileByIds) // 批量删除ImTxFile
		router.PUT("updateImTxFile", apiV1.UpdateImTxFile)              // 更新ImTxFile
		router.POST("quickEdit", apiV1.QuickEdit)                       // 快速编辑
	}
	{
		routerNoRecord.GET("findImTxFile", apiV1.FindImTxFile)       // 根据ID获取ImTxFile
		routerNoRecord.GET("getImTxFileList", apiV1.GetImTxFileList) // 获取ImTxFile列表
		routerNoRecord.GET("excelList", apiV1.ExcelList)             // 分页导出excel ImTxFile列表
	}
}
