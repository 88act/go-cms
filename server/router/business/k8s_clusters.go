package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type K8sClustersRouter struct {
}

// InitK8sClustersRouter 初始化 K8sClusters 路由信息
func (s *K8sClustersRouter) InitK8sClustersRouter(Router *gin.RouterGroup) {
	k8sClustersRouter := Router.Group("k8sClusters").Use(middleware.OperationRecord())
	k8sClustersRouterWithoutRecord := Router.Group("k8sClusters")
	var k8sClustersApi = v1.ApiGroupApp.BusinessApiGroup.K8sClustersApi
	{
		k8sClustersRouter.POST("createK8sClusters", k8sClustersApi.CreateK8sClusters)   // 新建K8sClusters
		k8sClustersRouter.DELETE("deleteK8sClusters", k8sClustersApi.DeleteK8sClusters) // 删除K8sClusters
		k8sClustersRouter.DELETE("deleteK8sClustersByIds", k8sClustersApi.DeleteK8sClustersByIds) // 批量删除K8sClusters
		k8sClustersRouter.PUT("updateK8sClusters", k8sClustersApi.UpdateK8sClusters)    // 更新K8sClusters
	    k8sClustersRouter.POST("quickEdit", k8sClustersApi.QuickEdit)  // 快速编辑
		k8sClustersRouterWithoutRecord.GET("excelList", k8sClustersApi.ExcelList)  // 分页导出excel K8sClusters列表
	}
	{
		k8sClustersRouterWithoutRecord.GET("findK8sClusters", k8sClustersApi.FindK8sClusters)        // 根据ID获取K8sClusters
		k8sClustersRouterWithoutRecord.GET("getK8sClustersList", k8sClustersApi.GetK8sClustersList)  // 获取K8sClusters列表
	    
	}
}
