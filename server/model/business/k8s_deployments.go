// 自动生成模板K8sDeployments
package business

import (
	"go-cms/global"
)

// K8sDeployments 结构体
// 如果含有time.Time 请自行import time包
type K8sDeployments struct {
	K8sDeploymentsMini
}

// K8sDeploymentsMini 结构体
type K8sDeploymentsMini struct {
	global.GO_MODEL
	Namespace  string `json:"namespace" cn:"命名空间" form:"namespace" gorm:"column:namespace;comment:命名空间;type:varchar(191);"`
	Deployment string `json:"deployment" cn:"应用" form:"deployment" gorm:"column:deployment;comment:应用;type:varchar(191);"`
	Replicas   *int   `json:"replicas" cn:"实例数" form:"replicas" gorm:"column:replicas;comment:实例数;type:int"`
	CreateTime string `json:"createTime" cn:"创建时间" form:"createTime" gorm:"column:create_time;comment:创建时间;type:varchar(191);"`
}

// TableName K8sDeployments 表名
func (K8sDeployments) TableName() string {
	return "k8s_deployments"
}
