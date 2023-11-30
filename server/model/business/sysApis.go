// 模板SysApis
// Author 10512203@qq.com
package business

import (
	"go-cms/global"
	"go-cms/model/common/request"
)

// 表格字段
const Field_SysApis_mini = "" // "id,created_at,updated_at,path,desc,model,api_group,method,status,cu_id,

// SysApis 结构体
type SysApis struct {
	global.BaseModel
	Path     string `json:"path" form:"path" cn:"api路径"  gorm:"column:path;comment:api路径;type:varchar(256);"`
	Desc     string `json:"desc" form:"desc" cn:"api中文描述"  gorm:"column:desc;comment:api中文描述;type:varchar(256);"`
	Model    string `json:"model" form:"model" cn:"所属模块"  gorm:"column:model;comment:所属模块;type:varchar(100);"`
	ApiGroup string `json:"apiGroup" form:"apiGroup" cn:"api组"  gorm:"column:api_group;comment:api组;type:varchar(256);"`
	Method   string `json:"method" form:"method" cn:"方法"  gorm:"column:method;comment:方法;type:varchar(256);"`
	Status   int    `json:"status" form:"status" cn:"状态"   gorm:"column:status;comment:状态;type:smallint"`
}

// TableName SysApis 表名
func (SysApis) TableName() string {
	return "sys_apis"
}

type SysApisSearch struct {
	request.PageInfo
	global.BaseModel
	Path     string `json:"path"   form:"path" `
	Desc     string `json:"desc"   form:"desc" `
	Model    string `json:"model"   form:"model" `
	ApiGroup string `json:"apiGroup"   form:"apiGroup" `
	Method   string `json:"method"   form:"method" `
	Status   *int   `json:"status"   form:"status" `
}
