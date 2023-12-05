package commom

import (
	v1 "go-cms/api/v1"

	"github.com/gin-gonic/gin"
)

type CommonFileRouter struct {
}

// InitCommonFileRouter 初始化 路由信息
func (s *CommonFileRouter) InitCommonFileRouter(Router *gin.RouterGroup) {

}

// 私有路由
func (s *CommonDbRouter) InitCommonFileRouterPrv(Router *gin.RouterGroup) {
	commonFileRouterWithoutRecord := Router.Group("commFile")
	var basicFileApi = v1.ApiGroupApp.BusinessApiGroup.BasicFileApi

	var fileApi = v1.ApiGroupApp.CommonApiGroup.CommonFileApi
	{
		commonFileRouterWithoutRecord.GET("get", basicFileApi.FindBasicFile) // 根据ID获取BasicFile
		commonFileRouterWithoutRecord.POST("upload", fileApi.UploadFile)
		commonFileRouterWithoutRecord.GET("getFileByKey", fileApi.GetFileByKey) // 根据sha1 key 获取BasicFile
	}

}
