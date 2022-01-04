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
	router := Router.Group("k8sDeployments").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("k8sDeployments")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.K8sDeploymentsApi
	{
		router.POST("createK8sDeployments", apiV1.CreateK8sDeployments)   // 新建K8sDeployments
		router.DELETE("deleteK8sDeployments", apiV1.DeleteK8sDeployments) // 删除K8sDeployments
		router.DELETE("deleteK8sDeploymentsByIds", apiV1.DeleteK8sDeploymentsByIds) // 批量删除K8sDeployments
		router.PUT("updateK8sDeployments", apiV1.UpdateK8sDeployments)    // 更新K8sDeployments
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("findK8sDeployments", apiV1.FindK8sDeployments)        // 根据ID获取K8sDeployments
		routerNoRecord.GET("getK8sDeploymentsList", apiV1.GetK8sDeploymentsList)  // 获取K8sDeployments列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel K8sDeployments列表
	}
}
