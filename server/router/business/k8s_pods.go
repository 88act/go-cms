package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type K8sPodsRouter struct {
}

// InitK8sPodsRouter 初始化 K8sPods 路由信息
func (s *K8sPodsRouter) InitK8sPodsRouter(Router *gin.RouterGroup) {
	router := Router.Group("k8sPods").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("k8sPods")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.K8sPodsApi
	{
		router.POST("createK8sPods", apiV1.CreateK8sPods)   // 新建K8sPods
		router.DELETE("deleteK8sPods", apiV1.DeleteK8sPods) // 删除K8sPods
		router.DELETE("deleteK8sPodsByIds", apiV1.DeleteK8sPodsByIds) // 批量删除K8sPods
		router.PUT("updateK8sPods", apiV1.UpdateK8sPods)    // 更新K8sPods
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findK8sPods", apiV1.FindK8sPods)        // 根据ID获取K8sPods
		routerNoRecord.GET("getK8sPodsList", apiV1.GetK8sPodsList)  // 获取K8sPods列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel K8sPods列表
	}
}
