// 模板CmsDetail
// Author 10512203@qq.com
package business

import (
	"go-cms/global"
	"go-cms/model/common/request"
)

// 表格字段
const Field_CmsDetail_mini = "" // "id,created_at,updated_at,art_id,detail,

// CmsDetail 结构体
type CmsDetail struct {
	global.BaseModel
	ArtId  int64  `json:"artId" form:"artId" cn:"文章id"  gorm:"column:art_id;comment:文章id;type:bigint"`
	Detail string `json:"detail" form:"detail" cn:"详细"  gorm:"column:detail;comment:详细;type:text;"`
}

// TableName CmsDetail 表名
func (CmsDetail) TableName() string {
	return "cms_detail"
}

type CmsDetailSearch struct {
	request.PageInfo
	global.BaseModel
	ArtId  *int64 `json:"artId"  form:"artId" `
	Detail string `json:"detail"   form:"detail" `
}
