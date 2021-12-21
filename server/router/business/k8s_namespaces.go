package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type K8sNamespacesRouter struct {
}

// InitK8sNamespacesRouter 初始化 K8sNamespaces 路由信息
func (s *K8sNamespacesRouter) InitK8sNamespacesRouter(Router *gin.RouterGroup) {
	k8sNamespacesRouter := Router.Group("k8sNamespaces").Use(middleware.OperationRecord())
	k8sNamespacesRouterWithoutRecord := Router.Group("k8sNamespaces")
	var k8sNamespacesApi = v1.ApiGroupApp.BusinessApiGroup.K8sNamespacesApi
	{
		k8sNamespacesRouter.POST("createK8sNamespaces", k8sNamespacesApi.CreateK8sNamespaces)   // 新建K8sNamespaces
		k8sNamespacesRouter.DELETE("deleteK8sNamespaces", k8sNamespacesApi.DeleteK8sNamespaces) // 删除K8sNamespaces
		k8sNamespacesRouter.DELETE("deleteK8sNamespacesByIds", k8sNamespacesApi.DeleteK8sNamespacesByIds) // 批量删除K8sNamespaces
		k8sNamespacesRouter.PUT("updateK8sNamespaces", k8sNamespacesApi.UpdateK8sNamespaces)    // 更新K8sNamespaces
	    k8sNamespacesRouter.POST("quickEdit", k8sNamespacesApi.QuickEdit)  // 快速编辑
		k8sNamespacesRouterWithoutRecord.GET("excelList", k8sNamespacesApi.ExcelList)  // 分页导出excel K8sNamespaces列表
	}
	{
		k8sNamespacesRouterWithoutRecord.GET("findK8sNamespaces", k8sNamespacesApi.FindK8sNamespaces)        // 根据ID获取K8sNamespaces
		k8sNamespacesRouterWithoutRecord.GET("getK8sNamespacesList", k8sNamespacesApi.GetK8sNamespacesList)  // 获取K8sNamespaces列表
	    
	}
}
