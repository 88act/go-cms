package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type BbPiDepartmentRouter struct {
}

// InitBbPiDepartmentRouter 初始化 BbPiDepartment 路由信息
func (s *BbPiDepartmentRouter) InitBbPiDepartmentRouter(Router *gin.RouterGroup) {
	router := Router.Group("bbPiDepartment").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("bbPiDepartment")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.BbPiDepartmentApi
	{
		router.POST("createBbPiDepartment", apiV1.CreateBbPiDepartment)   // 新建BbPiDepartment
		router.DELETE("deleteBbPiDepartment", apiV1.DeleteBbPiDepartment) // 删除BbPiDepartment
		router.DELETE("deleteBbPiDepartmentByIds", apiV1.DeleteBbPiDepartmentByIds) // 批量删除BbPiDepartment
		router.PUT("updateBbPiDepartment", apiV1.UpdateBbPiDepartment)    // 更新BbPiDepartment
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findBbPiDepartment", apiV1.FindBbPiDepartment)        // 根据ID获取BbPiDepartment
		routerNoRecord.GET("getBbPiDepartmentList", apiV1.GetBbPiDepartmentList)  // 获取BbPiDepartment列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel BbPiDepartment列表
	}
}
