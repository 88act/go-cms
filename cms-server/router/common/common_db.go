package commom

import (
	v1 "github.com/88act/go-cms/server/api/v1"
	"github.com/88act/go-cms/server/middleware"
	"github.com/gin-gonic/gin"
)

type CommonDbRouter struct {
}

// InitCommDbRouter 初始化 QuickEdit 路由信息
func (s *CommonDbRouter) InitCommonDbRouter(Router *gin.RouterGroup) {
	var dbApi = v1.ApiGroupApp.CommonApiGroup.CommonDbApi
	commonDbRouter := Router.Group("commDb").Use(middleware.OperationRecord())
	//commDbRouterWithoutRecord := Router.Group("commDb")
	{
		commonDbRouter.POST("quickEdit", dbApi.QuickEdit) // 更新 QuickEdit
	}

	{
		//commFileRouterWithoutRecord.GET("get", basicFileApi.FindBasicFile)                           // 根据ID获取BasicFile
		//commFileRouterWithoutRecord.POST("upload", v1.ApiGroupApp.CommonApiGroup.FileApi.UploadFile) // 获取BasicFile列表
	}

}
