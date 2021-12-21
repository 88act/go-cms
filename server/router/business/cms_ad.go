package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
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
		cmsAdRouterWithoutRecord.GET("excelList", cmsAdApi.ExcelList)  // 分页导出excel CmsAd列表
	}
	{
		cmsAdRouterWithoutRecord.GET("findCmsAd", cmsAdApi.FindCmsAd)        // 根据ID获取CmsAd
		cmsAdRouterWithoutRecord.GET("getCmsAdList", cmsAdApi.GetCmsAdList)  // 获取CmsAd列表
	    
	}
}
