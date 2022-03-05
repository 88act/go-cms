package business

import (
	"github.com/gin-gonic/gin"
	"go-cms/api/v1"
	"go-cms/middleware"
)

type CmsArticleRouter struct {
}

// InitCmsArticleRouter 初始化 CmsArticle 路由信息
func (s *CmsArticleRouter) InitCmsArticleRouter(Router *gin.RouterGroup) {
	router := Router.Group("cmsArticle").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("cmsArticle")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.CmsArticleApi
	{
		router.POST("createCmsArticle", apiV1.CreateCmsArticle)             // 新建CmsArticle
		router.DELETE("deleteCmsArticle", apiV1.DeleteCmsArticle)           // 删除CmsArticle
		router.DELETE("deleteCmsArticleByIds", apiV1.DeleteCmsArticleByIds) // 批量删除CmsArticle
		router.PUT("updateCmsArticle", apiV1.UpdateCmsArticle)              // 更新CmsArticle
		router.POST("quickEdit", apiV1.QuickEdit)                           // 快速编辑
	}
	{
		routerNoRecord.GET("findCmsArticle", apiV1.FindCmsArticle)       // 根据ID获取CmsArticle
		routerNoRecord.GET("getCmsArticleList", apiV1.GetCmsArticleList) // 获取CmsArticle列表
		routerNoRecord.GET("excelList", apiV1.ExcelList)                 // 分页导出excel CmsArticle列表
	}
}
