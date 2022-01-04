package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type K8sNodesRouter struct {
}

// InitK8sNodesRouter 初始化 K8sNodes 路由信息
func (s *K8sNodesRouter) InitK8sNodesRouter(Router *gin.RouterGroup) {
	router := Router.Group("k8sNodes").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("k8sNodes")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.K8sNodesApi
	{
		router.POST("createK8sNodes", apiV1.CreateK8sNodes)   // 新建K8sNodes
		router.DELETE("deleteK8sNodes", apiV1.DeleteK8sNodes) // 删除K8sNodes
		router.DELETE("deleteK8sNodesByIds", apiV1.DeleteK8sNodesByIds) // 批量删除K8sNodes
		router.PUT("updateK8sNodes", apiV1.UpdateK8sNodes)    // 更新K8sNodes
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findK8sNodes", apiV1.FindK8sNodes)        // 根据ID获取K8sNodes
		routerNoRecord.GET("getK8sNodesList", apiV1.GetK8sNodesList)  // 获取K8sNodes列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel K8sNodes列表
	}
}
