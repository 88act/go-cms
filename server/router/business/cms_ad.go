package business

import (
	"github.com/gin-gonic/gin"
	"go-cms/api/v1"
	"go-cms/middleware"
)

type CmsAdRouter struct {
}

// InitCmsAdRouter 初始化 CmsAd 路由信息
func (s *CmsAdRouter) InitCmsAdRouter(Router *gin.RouterGroup) {
	router := Router.Group("cmsAd").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("cmsAd")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.CmsAdApi
	{
		router.POST("createCmsAd", apiV1.CreateCmsAd)             // 新建CmsAd
		router.DELETE("deleteCmsAd", apiV1.DeleteCmsAd)           // 删除CmsAd
		router.DELETE("deleteCmsAdByIds", apiV1.DeleteCmsAdByIds) // 批量删除CmsAd
		router.PUT("updateCmsAd", apiV1.UpdateCmsAd)              // 更新CmsAd
		router.POST("quickEdit", apiV1.QuickEdit)                 // 快速编辑
	}
	{
		routerNoRecord.GET("findCmsAd", apiV1.FindCmsAd)       // 根据ID获取CmsAd
		routerNoRecord.GET("getCmsAdList", apiV1.GetCmsAdList) // 获取CmsAd列表
		routerNoRecord.GET("excelList", apiV1.ExcelList)       // 分页导出excel CmsAd列表
	}
}
