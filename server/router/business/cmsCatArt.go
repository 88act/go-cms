package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type CmsCatArtRouter struct {
}

// InitCmsCatArtRouter 初始化 CmsCatArt 路由信息
func (s *CmsCatArtRouter) InitCmsCatArtRouter(Router *gin.RouterGroup) {
	router := Router.Group("cmsCatArt").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("cmsCatArt")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.CmsCatArtApi
	{
		router.POST("createCmsCatArt", apiV1.CreateCmsCatArt)   // 新建CmsCatArt
		router.POST("deleteCmsCatArt", apiV1.DeleteCmsCatArt) // 删除CmsCatArt
		router.POST("deleteCmsCatArtByIds", apiV1.DeleteCmsCatArtByIds) // 批量删除CmsCatArt
		router.POST("updateCmsCatArt", apiV1.UpdateCmsCatArt)    // 更新CmsCatArt
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findCmsCatArt", apiV1.FindCmsCatArt)        // 根据ID获取CmsCatArt
		routerNoRecord.GET("getCmsCatArtList", apiV1.GetCmsCatArtList)  // 获取CmsCatArt列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel CmsCatArt列表
	}
}
