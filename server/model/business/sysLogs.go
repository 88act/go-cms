// 模板SysLogs
// Author 10512203@qq.com
package business

import (
	"go-cms/global"
	"go-cms/model/common/request"
)

// 表格字段
const Field_SysLogs_mini = "" // "id,created_at,updated_at,cu_id,user_id,log_type,method,path,latency,agent,error_message,body,resp,ip,ip_addr,status,

// SysLogs 结构体
type SysLogs struct {
	global.BaseModel
	UserId       int64  `json:"userId" form:"userId" cn:"用户id"  gorm:"column:user_id;comment:用户id;type:bigint"`
	LogType      int    `json:"logType" form:"logType" cn:"log类型"   gorm:"column:log_type;comment:log类型;type:smallint"`
	Method       string `json:"method" form:"method" cn:"请求方法"  gorm:"column:method;comment:请求方法;type:varchar(255);"`
	Path         string `json:"path" form:"path" cn:"请求路径"  gorm:"column:path;comment:请求路径;type:varchar(255);"`
	Latency      int64  `json:"latency" form:"latency" cn:"延迟"  gorm:"column:latency;comment:延迟;type:bigint"`
	Agent        string `json:"agent" form:"agent" cn:"代理"  gorm:"column:agent;comment:代理;type:varchar(255);"`
	ErrorMessage string `json:"errorMessage" form:"errorMessage" cn:"错误信息"  gorm:"column:error_message;comment:错误信息;type:varchar(255);"`
	Body         string `json:"body" form:"body" cn:"请求Body"  gorm:"column:body;comment:请求Body;type:text;"`
	Resp         string `json:"resp" form:"resp" cn:"响应Body"  gorm:"column:resp;comment:响应Body;type:text;"`
	Ip           string `json:"ip" form:"ip" cn:"请求ip"  gorm:"column:ip;comment:请求ip;type:varchar(100);"`
	IpAddr       string `json:"ipAddr" form:"ipAddr" cn:"地址"  gorm:"column:ip_addr;comment:地址;type:varchar(255);"`
	Status       int    `json:"status" form:"status" cn:"请求状态"   gorm:"column:status;comment:请求状态;type:int"`
}

// TableName SysLogs 表名
func (SysLogs) TableName() string {
	return "sys_logs"
}

type SysLogsSearch struct {
	request.PageInfo
	global.BaseModel
	UserId       *int64 `json:"userId"  form:"userId" `
	LogType      *int   `json:"logType"   form:"logType" `
	Method       string `json:"method"   form:"method" `
	Path         string `json:"path"   form:"path" `
	Latency      *int64 `json:"latency"  form:"latency" `
	Agent        string `json:"agent"   form:"agent" `
	ErrorMessage string `json:"errorMessage"   form:"errorMessage" `
	Body         string `json:"body"   form:"body" `
	Resp         string `json:"resp"   form:"resp" `
	Ip           string `json:"ip"   form:"ip" `
	IpAddr       string `json:"ipAddr"   form:"ipAddr" `
	Status       *int   `json:"status"   form:"status" `
}
