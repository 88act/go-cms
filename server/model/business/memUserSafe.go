// 模板MemUserSafe
// Author 10512203@qq.com
package business

import (
	"go-cms/global"
	"go-cms/model/common/request"
)

// 表格字段
const Field_MemUserSafe_mini = "" // "id,created_at,updated_at,user_id,safe_type,val_old,val_new,ip,ip_addr,status,

// MemUserSafe 结构体
type MemUserSafe struct {
	global.BaseModel
	SafeType int    `json:"safeType" form:"safeType" cn:"类型"   gorm:"column:safe_type;comment:类型:1改密码 2改手机号 3改用户名 4实名认证;type:smallint"`
	ValOld   string `json:"valOld" form:"valOld" cn:"旧值"  gorm:"column:val_old;comment:旧值;type:varchar(500);"`
	ValNew   string `json:"valNew" form:"valNew" cn:"新值"  gorm:"column:val_new;comment:新值;type:varchar(500);"`
	Ip       int    `json:"ip" form:"ip" cn:"ip"   gorm:"column:ip;comment:ip;type:int"`
	IpAddr   string `json:"ipAddr" form:"ipAddr" cn:"ip城市"  gorm:"column:ip_addr;comment:ip城市;type:varchar(255);"`
	Status   int    `json:"status" form:"status" cn:"状态"   gorm:"column:status;comment:状态:0=不正常,1=正常;type:smallint"`
}

// TableName MemUserSafe 表名
func (MemUserSafe) TableName() string {
	return "mem_user_safe"
}

type MemUserSafeSearch struct {
	request.PageInfo
	global.BaseModel
	SafeType *int   `json:"safeType"   form:"safeType" `
	ValOld   string `json:"valOld"   form:"valOld" `
	ValNew   string `json:"valNew"   form:"valNew" `
	Ip       *int   `json:"ip"   form:"ip" `
	IpAddr   string `json:"ipAddr"   form:"ipAddr" `
	Status   *int   `json:"status"   form:"status" `
}
