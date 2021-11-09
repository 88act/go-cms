package business

import (
	v1 "github.com/88act/go-cms/server/api/v1"
	"github.com/88act/go-cms/server/middleware"
	"github.com/gin-gonic/gin"
)

type BasicFileRouter struct {
}

// InitBasicFileRouter 初始化 BasicFile 路由信息
func (s *BasicFileRouter) InitBasicFileRouter(Router *gin.RouterGroup) {
	basicFileRouter := Router.Group("basicFile").Use(middleware.OperationRecord())
	basicFileRouterWithoutRecord := Router.Group("basicFile")
	var basicFileApi = v1.ApiGroupApp.BusinessApiGroup.BasicFileApi
	{
		basicFileRouter.POST("createBasicFile", basicFileApi.CreateBasicFile)             // 新建BasicFile
		basicFileRouter.DELETE("deleteBasicFile", basicFileApi.DeleteBasicFile)           // 删除BasicFile
		basicFileRouter.DELETE("deleteBasicFileByIds", basicFileApi.DeleteBasicFileByIds) // 批量删除BasicFile
		basicFileRouter.PUT("updateBasicFile", basicFileApi.UpdateBasicFile)              // 更新BasicFile
		basicFileRouter.POST("quickEdit", basicFileApi.QuickEdit)                         // 快速编辑
	}
	{
		basicFileRouterWithoutRecord.GET("findBasicFile", basicFileApi.FindBasicFile)       // 根据ID获取BasicFile
		basicFileRouterWithoutRecord.GET("getBasicFileList", basicFileApi.GetBasicFileList) // 获取BasicFile列表
	}
}
