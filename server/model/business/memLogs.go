// 模板MemLogs
// Author 10512203@qq.com
package business

import (
	"go-cms/global"
	"go-cms/model/common/request"
)

// 表格字段
const Field_MemLogs_mini = "" // "id,created_at,updated_at,cu_id,user_id,type,remark,ip,ip_addr,status,

// MemLogs 结构体
type MemLogs struct {
	global.BaseModel

	UserId int64  `json:"userId" form:"userId" cn:"商户id"  gorm:"column:user_id;comment:商户id;type:bigint"`
	Type   int    `json:"type" form:"type" cn:"类型"   gorm:"column:type;comment:类型: 1登录 2退出 3增加用户4 增加任务;type:smallint"`
	Remark string `json:"remark" form:"remark" cn:"备注"  gorm:"column:remark;comment:备注;type:text;"`
	Ip     int    `json:"ip" form:"ip" cn:"ip"   gorm:"column:ip;comment:ip;type:int"`
	IpAddr string `json:"ipAddr" form:"ipAddr" cn:"ip城市"  gorm:"column:ip_addr;comment:ip城市;type:varchar(255);"`
	Status int    `json:"status" form:"status" cn:"状态"   gorm:"column:status;comment:状态:0=不正常,1=正常;type:smallint"`
}

// TableName MemLogs 表名
func (MemLogs) TableName() string {
	return "mem_logs"
}

type MemLogsSearch struct {
	request.PageInfo
	global.BaseModel

	UserId *int64 `json:"userId"  form:"userId" `
	Type   *int   `json:"type"   form:"type" `
	Remark string `json:"remark"   form:"remark" `
	Ip     *int   `json:"ip"   form:"ip" `
	IpAddr string `json:"ipAddr"   form:"ipAddr" `
	Status *int   `json:"status"   form:"status" `
}
