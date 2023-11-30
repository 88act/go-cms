package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type CmsDetailRouter struct {
}

// InitCmsDetailRouter 初始化 CmsDetail 路由信息
func (s *CmsDetailRouter) InitCmsDetailRouter(Router *gin.RouterGroup) {
	router := Router.Group("cmsDetail").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("cmsDetail")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.CmsDetailApi
	{
		router.POST("createCmsDetail", apiV1.CreateCmsDetail)   // 新建CmsDetail
		router.POST("deleteCmsDetail", apiV1.DeleteCmsDetail) // 删除CmsDetail
		router.POST("deleteCmsDetailByIds", apiV1.DeleteCmsDetailByIds) // 批量删除CmsDetail
		router.POST("updateCmsDetail", apiV1.UpdateCmsDetail)    // 更新CmsDetail
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findCmsDetail", apiV1.FindCmsDetail)        // 根据ID获取CmsDetail
		routerNoRecord.GET("getCmsDetailList", apiV1.GetCmsDetailList)  // 获取CmsDetail列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel CmsDetail列表
	}
}
