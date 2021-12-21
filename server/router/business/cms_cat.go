package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type CmsCatRouter struct {
}

// InitCmsCatRouter 初始化 CmsCat 路由信息
func (s *CmsCatRouter) InitCmsCatRouter(Router *gin.RouterGroup) {
	cmsCatRouter := Router.Group("cmsCat").Use(middleware.OperationRecord())
	cmsCatRouterWithoutRecord := Router.Group("cmsCat")
	var cmsCatApi = v1.ApiGroupApp.BusinessApiGroup.CmsCatApi
	{
		cmsCatRouter.POST("createCmsCat", cmsCatApi.CreateCmsCat)   // 新建CmsCat
		cmsCatRouter.DELETE("deleteCmsCat", cmsCatApi.DeleteCmsCat) // 删除CmsCat
		cmsCatRouter.DELETE("deleteCmsCatByIds", cmsCatApi.DeleteCmsCatByIds) // 批量删除CmsCat
		cmsCatRouter.PUT("updateCmsCat", cmsCatApi.UpdateCmsCat)    // 更新CmsCat
	    cmsCatRouter.POST("quickEdit", cmsCatApi.QuickEdit)  // 快速编辑
		cmsCatRouterWithoutRecord.GET("excelList", cmsCatApi.ExcelList)  // 分页导出excel CmsCat列表
	}
	{
		cmsCatRouterWithoutRecord.GET("findCmsCat", cmsCatApi.FindCmsCat)        // 根据ID获取CmsCat
		cmsCatRouterWithoutRecord.GET("getCmsCatList", cmsCatApi.GetCmsCatList)  // 获取CmsCat列表
	    
	}
}
