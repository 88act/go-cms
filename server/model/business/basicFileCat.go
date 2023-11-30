// 模板BasicFileCat
// Author 10512203@qq.com
package business

import (
	"go-cms/global"
	"go-cms/model/common/request"
)

// 表格字段
const Field_BasicFileCat_mini = "" // "id,created_at,updated_at,pid,user_id,name,sort,status,

// BasicFileCat 结构体
type BasicFileCat struct {
	global.BaseModel
	Pid    int    `json:"pid" form:"pid" cn:"父ID"   gorm:"column:pid;comment:父ID;type:int"`
	UserId int64  `json:"userId" form:"userId" cn:"用户id"  gorm:"column:user_id;comment:用户id;type:bigint"`
	Name   string `json:"name" form:"name" cn:"名称"  gorm:"column:name;comment:名称;type:varchar(200);"`
	Sort   int    `json:"sort" form:"sort" cn:"排序"   gorm:"column:sort;comment:排序;type:int"`
	Status int    `json:"status" form:"status" cn:"状态"   gorm:"column:status;comment:状态:0未审核 1审核  2未通过审核 3 草稿;type:smallint"`
}

// TableName BasicFileCat 表名
func (BasicFileCat) TableName() string {
	return "basic_file_cat"
}

type BasicFileCatSearch struct {
	request.PageInfo
	global.BaseModel
	Pid    *int   `json:"pid"   form:"pid" `
	UserId *int64 `json:"userId"  form:"userId" `
	Name   string `json:"name"   form:"name" `
	Sort   *int   `json:"sort"   form:"sort" `
	Status *int   `json:"status"   form:"status" `
}
