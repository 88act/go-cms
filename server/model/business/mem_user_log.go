// 自动生成模板MemUserLog
package business

import (
	"go-cms/global"
)

// MemUserLog 结构体
// 如果含有time.Time 请自行import time包
type MemUserLog struct {
	MemUserLogMini
}

// MemUserLogMini 结构体
type MemUserLogMini struct {
	global.GO_MODEL
	UserId *int   `json:"userId" cn:"用户id" form:"userId" gorm:"column:user_id;comment:用户id;type:bigint"`
	Type   *int   `json:"type" cn:"类型" form:"type" gorm:"column:type;comment:类型:1登录 2退出;type:smallint"`
	Status *int   `json:"status" cn:"状态" form:"status" gorm:"column:status;comment:状态:0=不正常,1=正常;type:smallint"`
	Ip     string `json:"ip" cn:"ip" form:"ip" gorm:"column:ip;comment:ip;type:varchar(50);"`
}

// TableName MemUserLog 表名
func (MemUserLog) TableName() string {
	return "mem_user_log"
}
