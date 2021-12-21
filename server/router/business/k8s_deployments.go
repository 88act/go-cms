package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type K8sDeploymentsRouter struct {
}

// InitK8sDeploymentsRouter 初始化 K8sDeployments 路由信息
func (s *K8sDeploymentsRouter) InitK8sDeploymentsRouter(Router *gin.RouterGroup) {
	k8sDeploymentsRouter := Router.Group("k8sDeployments").Use(middleware.OperationRecord())
	k8sDeploymentsRouterWithoutRecord := Router.Group("k8sDeployments")
	var k8sDeploymentsApi = v1.ApiGroupApp.BusinessApiGroup.K8sDeploymentsApi
	{
		k8sDeploymentsRouter.POST("createK8sDeployments", k8sDeploymentsApi.CreateK8sDeployments)   // 新建K8sDeployments
		k8sDeploymentsRouter.DELETE("deleteK8sDeployments", k8sDeploymentsApi.DeleteK8sDeployments) // 删除K8sDeployments
		k8sDeploymentsRouter.DELETE("deleteK8sDeploymentsByIds", k8sDeploymentsApi.DeleteK8sDeploymentsByIds) // 批量删除K8sDeployments
		k8sDeploymentsRouter.PUT("updateK8sDeployments", k8sDeploymentsApi.UpdateK8sDeployments)    // 更新K8sDeployments
	    k8sDeploymentsRouter.POST("quickEdit", k8sDeploymentsApi.QuickEdit)  // 快速编辑
		k8sDeploymentsRouterWithoutRecord.GET("excelList", k8sDeploymentsApi.ExcelList)  // 分页导出excel K8sDeployments列表
	}
	{
		k8sDeploymentsRouterWithoutRecord.GET("findK8sDeployments", k8sDeploymentsApi.FindK8sDeployments)        // 根据ID获取K8sDeployments
		k8sDeploymentsRouterWithoutRecord.GET("getK8sDeploymentsList", k8sDeploymentsApi.GetK8sDeploymentsList)  // 获取K8sDeployments列表
	    
	}
}
