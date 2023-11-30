// 模板CmsTag
// Author 10512203@qq.com
package business

import (
	"go-cms/global"
	"go-cms/model/common/request"
)

// 表格字段
const Field_CmsTag_mini = "" // "id,created_at,updated_at,user_id,title,sort,status,

// CmsTag 结构体
type CmsTag struct {
	global.BaseModel
	UserId int64  `json:"userId" form:"userId" cn:"用户id"  gorm:"column:user_id;comment:用户id;type:bigint"`
	Title  string `json:"title" form:"title" cn:"名称"  gorm:"column:title;comment:名称;type:varchar(100);"`
	Sort   int    `json:"sort" form:"sort" cn:"排序"   gorm:"column:sort;comment:排序;type:int"`
	Status int    `json:"status" form:"status" cn:"状态"   gorm:"column:status;comment:状态:0未审核 1审核 2未通过审核 3 草稿;type:smallint"`
}

// TableName CmsTag 表名
func (CmsTag) TableName() string {
	return "cms_tag"
}

type CmsTagSearch struct {
	request.PageInfo
	global.BaseModel
	UserId *int64 `json:"userId"  form:"userId" `
	Title  string `json:"title"   form:"title" `
	Sort   *int   `json:"sort"   form:"sort" `
	Status *int   `json:"status"   form:"status" `
}
