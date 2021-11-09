package business

import (
	"github.com/88act/go-cms/server/api/v1"
	"github.com/88act/go-cms/server/middleware"
	"github.com/gin-gonic/gin"
)

type CmsAdRouter struct {
}

// InitCmsAdRouter 初始化 CmsAd 路由信息
func (s *CmsAdRouter) InitCmsAdRouter(Router *gin.RouterGroup) {
	cmsAdRouter := Router.Group("cmsAd").Use(middleware.OperationRecord())
	cmsAdRouterWithoutRecord := Router.Group("cmsAd")
	var cmsAdApi = v1.ApiGroupApp.BusinessApiGroup.CmsAdApi
	{
		cmsAdRouter.POST("createCmsAd", cmsAdApi.CreateCmsAd)   // 新建CmsAd
		cmsAdRouter.DELETE("deleteCmsAd", cmsAdApi.DeleteCmsAd) // 删除CmsAd
		cmsAdRouter.DELETE("deleteCmsAdByIds", cmsAdApi.DeleteCmsAdByIds) // 批量删除CmsAd
		cmsAdRouter.PUT("updateCmsAd", cmsAdApi.UpdateCmsAd)    // 更新CmsAd
	    cmsAdRouter.POST("quickEdit", cmsAdApi.QuickEdit)  // 快速编辑
	}
	{
		cmsAdRouterWithoutRecord.GET("findCmsAd", cmsAdApi.FindCmsAd)        // 根据ID获取CmsAd
		cmsAdRouterWithoutRecord.GET("getCmsAdList", cmsAdApi.GetCmsAdList)  // 获取CmsAd列表
	}
}
