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
	cmsAdSeatRouter := Router.Group("cmsAdSeat").Use(middleware.OperationRecord())
	cmsAdSeatRouterWithoutRecord := Router.Group("cmsAdSeat")
	var cmsAdSeatApi = v1.ApiGroupApp.BusinessApiGroup.CmsAdSeatApi
	{
		cmsAdSeatRouter.POST("createCmsAdSeat", cmsAdSeatApi.CreateCmsAdSeat)   // 新建CmsAdSeat
		cmsAdSeatRouter.DELETE("deleteCmsAdSeat", cmsAdSeatApi.DeleteCmsAdSeat) // 删除CmsAdSeat
		cmsAdSeatRouter.DELETE("deleteCmsAdSeatByIds", cmsAdSeatApi.DeleteCmsAdSeatByIds) // 批量删除CmsAdSeat
		cmsAdSeatRouter.PUT("updateCmsAdSeat", cmsAdSeatApi.UpdateCmsAdSeat)    // 更新CmsAdSeat
	    cmsAdSeatRouter.POST("quickEdit", cmsAdSeatApi.QuickEdit)  // 快速编辑
		cmsAdSeatRouterWithoutRecord.GET("excelList", cmsAdSeatApi.ExcelList)  // 分页导出excel CmsAdSeat列表
	}
	{
		cmsAdSeatRouterWithoutRecord.GET("findCmsAdSeat", cmsAdSeatApi.FindCmsAdSeat)        // 根据ID获取CmsAdSeat
		cmsAdSeatRouterWithoutRecord.GET("getCmsAdSeatList", cmsAdSeatApi.GetCmsAdSeatList)  // 获取CmsAdSeat列表
	    
	}
}
