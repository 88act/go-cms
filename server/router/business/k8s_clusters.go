package business

import (
	"github.com/gin-gonic/gin"
	"go-cms/api/v1"
	"go-cms/middleware"
)

type K8sClustersRouter struct {
}

// InitK8sClustersRouter 初始化 K8sClusters 路由信息
func (s *K8sClustersRouter) InitK8sClustersRouter(Router *gin.RouterGroup) {
	router := Router.Group("k8sClusters").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("k8sClusters")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.K8sClustersApi
	{
		router.POST("createK8sClusters", apiV1.CreateK8sClusters)             // 新建K8sClusters
		router.DELETE("deleteK8sClusters", apiV1.DeleteK8sClusters)           // 删除K8sClusters
		router.DELETE("deleteK8sClustersByIds", apiV1.DeleteK8sClustersByIds) // 批量删除K8sClusters
		router.PUT("updateK8sClusters", apiV1.UpdateK8sClusters)              // 更新K8sClusters
		router.POST("quickEdit", apiV1.QuickEdit)                             // 快速编辑
	}
	{
		routerNoRecord.GET("findK8sClusters", apiV1.FindK8sClusters)       // 根据ID获取K8sClusters
		routerNoRecord.GET("getK8sClustersList", apiV1.GetK8sClustersList) // 获取K8sClusters列表
		routerNoRecord.GET("excelList", apiV1.ExcelList)                   // 分页导出excel K8sClusters列表
	}
}
