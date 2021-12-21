// 自动生成模板K8sNodes
package business

import (
	"go-cms/global"
)

// K8sNodes 结构体
// 如果含有time.Time 请自行import time包
type K8sNodes struct {
	K8sNodesMini
}

// K8sNodesMini 结构体
type K8sNodesMini struct {
	global.GO_MODEL
	NodeName   string `json:"nodeName" cn:"节点名称" form:"nodeName" gorm:"column:node_name;comment:节点名称;type:varchar(191);"`
	Ip         string `json:"ip" cn:"ip地址" form:"ip" gorm:"column:ip;comment:ip地址;type:varchar(191);"`
	Status     string `json:"status" cn:"节点状态" form:"status" gorm:"column:status;comment:节点状态;type:varchar(191);"`
	Roles      string `json:"roles" cn:"节点角色" form:"roles" gorm:"column:roles;comment:节点角色;type:varchar(191);"`
	CreateTime string `json:"createTime" cn:"创建时间" form:"createTime" gorm:"column:create_time;comment:创建时间;type:varchar(191);"`
	Version    string `json:"version" cn:"节点版本" form:"version" gorm:"column:version;comment:节点版本;type:varchar(191);"`
	Resource   string `json:"resource" cn:"节点资源" form:"resource" gorm:"column:resource;comment:节点资源;type:varchar(191);"`
	Label      string `json:"label" cn:"节点标签" form:"label" gorm:"column:label;comment:节点标签;type:varchar(191);"`
}

// TableName K8sNodes 表名
func (K8sNodes) TableName() string {
	return "k8s_nodes"
}
