package business

import (
	v1 "go-cms/api/v1"
	"go-cms/middleware"

	"github.com/gin-gonic/gin"
)

type BasicFileRouter struct {
}

// InitBasicFileRouter 初始化 BasicFile 路由信息
func (s *BasicFileRouter) InitBasicFileRouter(Router *gin.RouterGroup) {
	router := Router.Group("basicFile").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("basicFile")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.BasicFileApi
	{
		router.POST("createBasicFile", apiV1.CreateBasicFile)           // 新建BasicFile
		router.POST("deleteBasicFile", apiV1.DeleteBasicFile)           // 删除BasicFile
		router.POST("deleteBasicFileByIds", apiV1.DeleteBasicFileByIds) // 批量删除BasicFile
		router.POST("updateBasicFile", apiV1.UpdateBasicFile)           // 更新BasicFile
		router.POST("quickEdit", apiV1.QuickEdit)                       // 快速编辑
	}
	{
		routerNoRecord.GET("findBasicFile", apiV1.FindBasicFile)       // 根据ID获取BasicFile
		routerNoRecord.GET("getBasicFileList", apiV1.GetBasicFileList) // 获取BasicFile列表
		routerNoRecord.GET("excelList", apiV1.ExcelList)               // 分页导出excel BasicFile列表
	}
}
