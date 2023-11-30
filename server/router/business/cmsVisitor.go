package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type CmsVisitorRouter struct {
}

// InitCmsVisitorRouter 初始化 CmsVisitor 路由信息
func (s *CmsVisitorRouter) InitCmsVisitorRouter(Router *gin.RouterGroup) {
	router := Router.Group("cmsVisitor").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("cmsVisitor")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.CmsVisitorApi
	{
		router.POST("createCmsVisitor", apiV1.CreateCmsVisitor)   // 新建CmsVisitor
		router.POST("deleteCmsVisitor", apiV1.DeleteCmsVisitor) // 删除CmsVisitor
		router.POST("deleteCmsVisitorByIds", apiV1.DeleteCmsVisitorByIds) // 批量删除CmsVisitor
		router.POST("updateCmsVisitor", apiV1.UpdateCmsVisitor)    // 更新CmsVisitor
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findCmsVisitor", apiV1.FindCmsVisitor)        // 根据ID获取CmsVisitor
		routerNoRecord.GET("getCmsVisitorList", apiV1.GetCmsVisitorList)  // 获取CmsVisitor列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel CmsVisitor列表
	}
}
