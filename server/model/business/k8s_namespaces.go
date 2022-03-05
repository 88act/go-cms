// 自动生成模板K8sNamespaces
package business

import (
	"go-cms/global"
)

// K8sNamespaces 结构体
// 如果含有time.Time 请自行import time包
type K8sNamespaces struct {
	K8sNamespacesMini
}

// K8sNamespacesMini 结构体
type K8sNamespacesMini struct {
	global.GO_MODEL
	Namespace  string `json:"namespace" cn:"命名空间" form:"namespace" gorm:"column:namespace;comment:命名空间;type:varchar(191);"`
	Status     string `json:"status" cn:"状态" form:"status" gorm:"column:status;comment:状态;type:varchar(191);"`
	CreateTime string `json:"createTime" cn:"创建时间" form:"createTime" gorm:"column:create_time;comment:创建时间;type:varchar(191);"`
}

// TableName K8sNamespaces 表名
func (K8sNamespaces) TableName() string {
	return "k8s_namespaces"
}
