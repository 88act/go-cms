package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type CmsDiscussRouter struct {
}

// InitCmsDiscussRouter 初始化 CmsDiscuss 路由信息
func (s *CmsDiscussRouter) InitCmsDiscussRouter(Router *gin.RouterGroup) {
	router := Router.Group("cmsDiscuss").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("cmsDiscuss")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.CmsDiscussApi
	{
		router.POST("createCmsDiscuss", apiV1.CreateCmsDiscuss)   // 新建CmsDiscuss
		router.POST("deleteCmsDiscuss", apiV1.DeleteCmsDiscuss) // 删除CmsDiscuss
		router.POST("deleteCmsDiscussByIds", apiV1.DeleteCmsDiscussByIds) // 批量删除CmsDiscuss
		router.POST("updateCmsDiscuss", apiV1.UpdateCmsDiscuss)    // 更新CmsDiscuss
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findCmsDiscuss", apiV1.FindCmsDiscuss)        // 根据ID获取CmsDiscuss
		routerNoRecord.GET("getCmsDiscussList", apiV1.GetCmsDiscussList)  // 获取CmsDiscuss列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel CmsDiscuss列表
	}
}
