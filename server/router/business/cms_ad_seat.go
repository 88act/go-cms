package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type CmsAdSeatRouter struct {
}

// InitCmsAdSeatRouter 初始化 CmsAdSeat 路由信息
func (s *CmsAdSeatRouter) InitCmsAdSeatRouter(Router *gin.RouterGroup) {
	router := Router.Group("cmsAdSeat").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("cmsAdSeat")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.CmsAdSeatApi
	{
		router.POST("createCmsAdSeat", apiV1.CreateCmsAdSeat)   // 新建CmsAdSeat
		router.DELETE("deleteCmsAdSeat", apiV1.DeleteCmsAdSeat) // 删除CmsAdSeat
		router.DELETE("deleteCmsAdSeatByIds", apiV1.DeleteCmsAdSeatByIds) // 批量删除CmsAdSeat
		router.PUT("updateCmsAdSeat", apiV1.UpdateCmsAdSeat)    // 更新CmsAdSeat
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findCmsAdSeat", apiV1.FindCmsAdSeat)        // 根据ID获取CmsAdSeat
		routerNoRecord.GET("getCmsAdSeatList", apiV1.GetCmsAdSeatList)  // 获取CmsAdSeat列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel CmsAdSeat列表
	}
}
