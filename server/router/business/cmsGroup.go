package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type CmsGroupRouter struct {
}

// InitCmsGroupRouter 初始化 CmsGroup 路由信息
func (s *CmsGroupRouter) InitCmsGroupRouter(Router *gin.RouterGroup) {
	router := Router.Group("cmsGroup").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("cmsGroup")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.CmsGroupApi
	{
		router.POST("createCmsGroup", apiV1.CreateCmsGroup)   // 新建CmsGroup
		router.POST("deleteCmsGroup", apiV1.DeleteCmsGroup) // 删除CmsGroup
		router.POST("deleteCmsGroupByIds", apiV1.DeleteCmsGroupByIds) // 批量删除CmsGroup
		router.POST("updateCmsGroup", apiV1.UpdateCmsGroup)    // 更新CmsGroup
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findCmsGroup", apiV1.FindCmsGroup)        // 根据ID获取CmsGroup
		routerNoRecord.GET("getCmsGroupList", apiV1.GetCmsGroupList)  // 获取CmsGroup列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel CmsGroup列表
	}
}
