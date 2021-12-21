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
	k8sPodsRouter := Router.Group("k8sPods").Use(middleware.OperationRecord())
	k8sPodsRouterWithoutRecord := Router.Group("k8sPods")
	var k8sPodsApi = v1.ApiGroupApp.BusinessApiGroup.K8sPodsApi
	{
		k8sPodsRouter.POST("createK8sPods", k8sPodsApi.CreateK8sPods)   // 新建K8sPods
		k8sPodsRouter.DELETE("deleteK8sPods", k8sPodsApi.DeleteK8sPods) // 删除K8sPods
		k8sPodsRouter.DELETE("deleteK8sPodsByIds", k8sPodsApi.DeleteK8sPodsByIds) // 批量删除K8sPods
		k8sPodsRouter.PUT("updateK8sPods", k8sPodsApi.UpdateK8sPods)    // 更新K8sPods
	    k8sPodsRouter.POST("quickEdit", k8sPodsApi.QuickEdit)  // 快速编辑
		k8sPodsRouterWithoutRecord.GET("excelList", k8sPodsApi.ExcelList)  // 分页导出excel K8sPods列表
	}
	{
		k8sPodsRouterWithoutRecord.GET("findK8sPods", k8sPodsApi.FindK8sPods)        // 根据ID获取K8sPods
		k8sPodsRouterWithoutRecord.GET("getK8sPodsList", k8sPodsApi.GetK8sPodsList)  // 获取K8sPods列表
	    
	}
}
