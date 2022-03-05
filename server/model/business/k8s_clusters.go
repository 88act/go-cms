// 自动生成模板K8sClusters
package business

import (
	"go-cms/global"
)

// K8sClusters 结构体
// 如果含有time.Time 请自行import time包
type K8sClusters struct {
	K8sClustersMini
}

// K8sClustersMini 结构体
type K8sClustersMini struct {
	global.GO_MODEL
	ClusterName    string   `json:"clusterName" cn:"集群名称" form:"clusterName" gorm:"column:cluster_name;comment:集群名称;type:varchar(191);"`
	KubeConfig     string   `json:"kubeConfig" cn:"Kubeconfig内容" form:"kubeConfig" gorm:"column:kube_config;comment:Kubeconfig内容;type:varchar(12800);"`
	ClusterVersion *float64 `json:"clusterVersion" cn:"集群版本" form:"clusterVersion" gorm:"column:cluster_version;comment:集群版本;type:float"`
}

// TableName K8sClusters 表名
func (K8sClusters) TableName() string {
	return "k8s_clusters"
}
