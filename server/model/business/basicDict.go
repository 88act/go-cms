// 模板BasicDict
// Author 10512203@qq.com
package business

import (
	"go-cms/global"
	"go-cms/model/common/request"
)

// 表格字段
const Field_BasicDict_mini = "" // "id,created_at,updated_at,type,pid,module,label,value,icon,remark,sort,status,key,

// BasicDict 结构体
type BasicDict struct {
	global.BaseModel
	Type   int    `json:"type" form:"type" cn:"类别"   gorm:"column:type;comment:类别:0=普通 1=系统;type:smallint"`
	Pid    int64  `json:"pid" form:"pid" cn:"父id"  gorm:"column:pid;comment:父id;type:bigint"`
	Module string `json:"module" form:"module" cn:"模块名"  gorm:"column:module;comment:模块名;type:varchar(100);"`
	Label  string `json:"label" form:"label" cn:"名称"  gorm:"column:label;comment:名称;type:varchar(200);"`
	Value  string `json:"value" form:"value" cn:"key"  gorm:"column:value;comment:key;type:varchar(200);"`
	Icon   string `json:"icon" form:"icon" cn:"图标"  gorm:"column:icon;comment:图标;type:varchar(256);"`
	Remark string `json:"remark" form:"remark" cn:"备注"  gorm:"column:remark;comment:备注;type:varchar(500);"`
	Sort   int    `json:"sort" form:"sort" cn:"排序"   gorm:"column:sort;comment:排序;type:int"`
	Status int    `json:"status" form:"status" cn:"状态"   gorm:"column:status;comment:状态;type:smallint"`
	Key    string `json:"key" form:"key" cn:"key字段"  gorm:"column:key;comment:;type:varchar(20);"`
}

// TableName BasicDict 表名
func (BasicDict) TableName() string {
	return "basic_dict"
}

type BasicDictSearch struct {
	request.PageInfo
	global.BaseModel
	Type   *int   `json:"type"   form:"type" `
	Pid    *int64 `json:"pid"  form:"pid" `
	Module string `json:"module"   form:"module" `
	Label  string `json:"label"   form:"label" `
	Value  string `json:"value"   form:"value" `
	Icon   string `json:"icon"   form:"icon" `
	Remark string `json:"remark"   form:"remark" `
	Sort   *int   `json:"sort"   form:"sort" `
	Status *int   `json:"status"   form:"status" `
	Key    string `json:"key"   form:"key" `
}
