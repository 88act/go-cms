package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type CmsTagRouter struct {
}

// InitCmsTagRouter 初始化 CmsTag 路由信息
func (s *CmsTagRouter) InitCmsTagRouter(Router *gin.RouterGroup) {
	router := Router.Group("cmsTag").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("cmsTag")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.CmsTagApi
	{
		router.POST("createCmsTag", apiV1.CreateCmsTag)   // 新建CmsTag
		router.POST("deleteCmsTag", apiV1.DeleteCmsTag) // 删除CmsTag
		router.POST("deleteCmsTagByIds", apiV1.DeleteCmsTagByIds) // 批量删除CmsTag
		router.POST("updateCmsTag", apiV1.UpdateCmsTag)    // 更新CmsTag
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findCmsTag", apiV1.FindCmsTag)        // 根据ID获取CmsTag
		routerNoRecord.GET("getCmsTagList", apiV1.GetCmsTagList)  // 获取CmsTag列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel CmsTag列表
	}
}
