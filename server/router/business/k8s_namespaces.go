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
	router := Router.Group("k8sNamespaces").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("k8sNamespaces")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.K8sNamespacesApi
	{
		router.POST("createK8sNamespaces", apiV1.CreateK8sNamespaces)   // 新建K8sNamespaces
		router.DELETE("deleteK8sNamespaces", apiV1.DeleteK8sNamespaces) // 删除K8sNamespaces
		router.DELETE("deleteK8sNamespacesByIds", apiV1.DeleteK8sNamespacesByIds) // 批量删除K8sNamespaces
		router.PUT("updateK8sNamespaces", apiV1.UpdateK8sNamespaces)    // 更新K8sNamespaces
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findK8sNamespaces", apiV1.FindK8sNamespaces)        // 根据ID获取K8sNamespaces
		routerNoRecord.GET("getK8sNamespacesList", apiV1.GetK8sNamespacesList)  // 获取K8sNamespaces列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel K8sNamespaces列表
	}
}
