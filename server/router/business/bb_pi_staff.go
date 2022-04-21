package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type BbPiStaffRouter struct {
}

// InitBbPiStaffRouter 初始化 BbPiStaff 路由信息
func (s *BbPiStaffRouter) InitBbPiStaffRouter(Router *gin.RouterGroup) {
	router := Router.Group("bbPiStaff").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("bbPiStaff")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.BbPiStaffApi
	{
		router.POST("createBbPiStaff", apiV1.CreateBbPiStaff)   // 新建BbPiStaff
		router.DELETE("deleteBbPiStaff", apiV1.DeleteBbPiStaff) // 删除BbPiStaff
		router.DELETE("deleteBbPiStaffByIds", apiV1.DeleteBbPiStaffByIds) // 批量删除BbPiStaff
		router.PUT("updateBbPiStaff", apiV1.UpdateBbPiStaff)    // 更新BbPiStaff
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findBbPiStaff", apiV1.FindBbPiStaff)        // 根据ID获取BbPiStaff
		routerNoRecord.GET("getBbPiStaffList", apiV1.GetBbPiStaffList)  // 获取BbPiStaff列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel BbPiStaff列表
	}
}
