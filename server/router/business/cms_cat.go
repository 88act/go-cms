package business

import (
	"github.com/gin-gonic/gin"
	"go-cms/api/v1"
	"go-cms/middleware"
)

type CmsCatRouter struct {
}

// InitCmsCatRouter 初始化 CmsCat 路由信息
func (s *CmsCatRouter) InitCmsCatRouter(Router *gin.RouterGroup) {
	router := Router.Group("cmsCat").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("cmsCat")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.CmsCatApi
	{
		router.POST("createCmsCat", apiV1.CreateCmsCat)             // 新建CmsCat
		router.DELETE("deleteCmsCat", apiV1.DeleteCmsCat)           // 删除CmsCat
		router.DELETE("deleteCmsCatByIds", apiV1.DeleteCmsCatByIds) // 批量删除CmsCat
		router.PUT("updateCmsCat", apiV1.UpdateCmsCat)              // 更新CmsCat
		router.POST("quickEdit", apiV1.QuickEdit)                   // 快速编辑
	}
	{
		routerNoRecord.GET("findCmsCat", apiV1.FindCmsCat)       // 根据ID获取CmsCat
		routerNoRecord.GET("getCmsCatList", apiV1.GetCmsCatList) // 获取CmsCat列表
		routerNoRecord.GET("excelList", apiV1.ExcelList)         // 分页导出excel CmsCat列表
	}
}
