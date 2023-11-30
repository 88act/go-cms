package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type CmsBlogRouter struct {
}

// InitCmsBlogRouter 初始化 CmsBlog 路由信息
func (s *CmsBlogRouter) InitCmsBlogRouter(Router *gin.RouterGroup) {
	router := Router.Group("cmsBlog").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("cmsBlog")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.CmsBlogApi
	{
		router.POST("createCmsBlog", apiV1.CreateCmsBlog)   // 新建CmsBlog
		router.POST("deleteCmsBlog", apiV1.DeleteCmsBlog) // 删除CmsBlog
		router.POST("deleteCmsBlogByIds", apiV1.DeleteCmsBlogByIds) // 批量删除CmsBlog
		router.POST("updateCmsBlog", apiV1.UpdateCmsBlog)    // 更新CmsBlog
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findCmsBlog", apiV1.FindCmsBlog)        // 根据ID获取CmsBlog
		routerNoRecord.GET("getCmsBlogList", apiV1.GetCmsBlogList)  // 获取CmsBlog列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel CmsBlog列表
	}
}
