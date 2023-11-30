package client

import (
	v1 "go-cms/api/v1"

	"github.com/gin-gonic/gin"
)

type userRouter struct {
}

// InitCommDbRouter 初始化 QuickEdit 路由信息
func (s *userRouter) InitUserRouter(Router *gin.RouterGroup) {
	var dbApi = v1.ApiGroupApp.CommonApiGroup.CommonDbApi
	userRouter := Router.Group("user")
	//commDbRouterWithoutRecord := Router.Group("commDb")
	{
		userRouter.POST("quickEdit", dbApi.QuickEdit) // 更新 QuickEdit
		//userRouter.GET("test_db_get", dbApi.Test_db_get)
		//userRouter.GET("test_db", dbApi.Test_db)
	}

}
