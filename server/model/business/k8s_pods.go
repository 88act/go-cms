// 自动生成模板K8sPods
package business

import (
	"go-cms/global"
)

// K8sPods 结构体
// 如果含有time.Time 请自行import time包
type K8sPods struct {
     K8sPodsMini
}

// K8sPodsMini 结构体
type K8sPodsMini struct {
      global.GO_MODEL
      PodName  string `json:"podName" cn:"容器名称" form:"podName" gorm:"column:pod_name;comment:容器名称;type:varchar(191);"`
      PodIp  string `json:"podIp" cn:"容器IP" form:"podIp" gorm:"column:pod_ip;comment:容器IP;type:varchar(191);"`
      HostIp  string `json:"hostIp" cn:"主机IP" form:"hostIp" gorm:"column:host_ip;comment:主机IP;type:varchar(191);"`
      Status  string `json:"status" cn:"状态" form:"status" gorm:"column:status;comment:状态;type:varchar(191);"`
      StartTime  string `json:"startTime" cn:"启动时间" form:"startTime" gorm:"column:start_time;comment:启动时间;type:varchar(191);"`
      RestartCount  *int `json:"restartCount" cn:"重启次数" form:"restartCount" gorm:"column:restart_count;comment:重启次数;type:int"`
      NameSpace  string `json:"nameSpace" cn:"命名空间" form:"nameSpace" gorm:"column:name_space;comment:命名空间;type:varchar(191);"`
      Line  *int `json:"line" cn:"line" form:"line" gorm:"column:line;comment:line;type:bigint"`

}


// TableName K8sPods 表名
func (K8sPods) TableName() string {
  return "k8s_pods"
}

