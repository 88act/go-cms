package business

import (
	v1 "go-cms/api/v1"
	"go-cms/middleware"

	"github.com/gin-gonic/gin"
)

type CmsArtRouter struct {
}

// InitCmsArtRouter 初始化 CmsArt 路由信息
func (s *CmsArtRouter) InitCmsArtRouter(Router *gin.RouterGroup) {
	router := Router.Group("cmsArt").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("cmsArt")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.CmsArtApi
	{
		router.POST("createCmsArt", apiV1.CreateCmsArt)           // 新建CmsArt
		router.POST("deleteCmsArt", apiV1.DeleteCmsArt)           // 删除CmsArt
		router.POST("deleteCmsArtByIds", apiV1.DeleteCmsArtByIds) // 批量删除CmsArt
		router.POST("updateCmsArt", apiV1.UpdateCmsArt)           // 更新CmsArt
		router.POST("quickEdit", apiV1.QuickEdit)                 // 快速编辑
	}
	{
		routerNoRecord.GET("findCmsArt", apiV1.FindCmsArt)       // 根据ID获取CmsArt
		routerNoRecord.GET("getCmsArtList", apiV1.GetCmsArtList) // 获取CmsArt列表
		routerNoRecord.GET("excelList", apiV1.ExcelList)         // 分页导出excel CmsArt列表
	}
}
