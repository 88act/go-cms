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
	k8sNodesRouter := Router.Group("k8sNodes").Use(middleware.OperationRecord())
	k8sNodesRouterWithoutRecord := Router.Group("k8sNodes")
	var k8sNodesApi = v1.ApiGroupApp.BusinessApiGroup.K8sNodesApi
	{
		k8sNodesRouter.POST("createK8sNodes", k8sNodesApi.CreateK8sNodes)   // 新建K8sNodes
		k8sNodesRouter.DELETE("deleteK8sNodes", k8sNodesApi.DeleteK8sNodes) // 删除K8sNodes
		k8sNodesRouter.DELETE("deleteK8sNodesByIds", k8sNodesApi.DeleteK8sNodesByIds) // 批量删除K8sNodes
		k8sNodesRouter.PUT("updateK8sNodes", k8sNodesApi.UpdateK8sNodes)    // 更新K8sNodes
	    k8sNodesRouter.POST("quickEdit", k8sNodesApi.QuickEdit)  // 快速编辑
		k8sNodesRouterWithoutRecord.GET("excelList", k8sNodesApi.ExcelList)  // 分页导出excel K8sNodes列表
	}
	{
		k8sNodesRouterWithoutRecord.GET("findK8sNodes", k8sNodesApi.FindK8sNodes)        // 根据ID获取K8sNodes
		k8sNodesRouterWithoutRecord.GET("getK8sNodesList", k8sNodesApi.GetK8sNodesList)  // 获取K8sNodes列表
	    
	}
}
