package client

import (
	v1 "go-cms/api/v1"

	"github.com/gin-gonic/gin"
)

type cmsRouter struct {
}

// InitCommDbRouter 初始化 QuickEdit 路由信息
func (s *cmsRouter) InitCmsRouter(Router *gin.RouterGroup) {
	var cmsApi = v1.ApiGroupApp.ClientGroup.CmsApi
	cmsRouter := Router.Group("")
	//commDbRouterWithoutRecord := Router.Group("commDb")
	{

		cmsRouter.GET("/", cmsApi.Index)           // 更新 QuickEdit
		cmsRouter.GET("/index.html", cmsApi.Index) // 更新 QuickEdit
		cmsRouter.GET("login.html", cmsApi.Login)  // 更新 QuickEdit
		//userRouter.GET("test_db_get", dbApi.Test_db_get)
		//userRouter.GET("test_db", dbApi.Test_db)
	}

}
