package system

import (
	v1 "github.com/88act/go-cms/server/api/v1"
	"github.com/88act/go-cms/server/middleware"
	"github.com/gin-gonic/gin"
)

type CasbinRouter struct {
}

func (s *CasbinRouter) InitCasbinRouter(Router *gin.RouterGroup) {
	casbinRouter := Router.Group("casbin").Use(middleware.OperationRecord())
	casbinRouterWithoutRecord := Router.Group("casbin")
	var casbinApi = v1.ApiGroupApp.SystemApiGroup.CasbinApi
	{
		casbinRouter.POST("updateCasbin", casbinApi.UpdateCasbin)
	}
	{
		casbinRouterWithoutRecord.POST("getPolicyPathByAuthorityId", casbinApi.GetPolicyPathByAuthorityId)
	}
}
