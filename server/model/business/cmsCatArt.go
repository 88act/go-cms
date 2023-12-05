// 模板CmsCatArt
// Author 10512203@qq.com
package business

import (
	"go-cms/global"
	"go-cms/model/common/request"
)

// 表格字段
const Field_CmsCatArt_mini = "" // "id,created_at,updated_at,user_id,cat_id,art_id,status,

// CmsCatArt 结构体
type CmsCatArt struct {
	global.BaseModel
	UserId int64 `json:"userId" form:"userId" cn:"用户id"  gorm:"column:user_id;comment:用户id;type:bigint"`
	CatId  int64 `json:"catId" form:"catId" cn:"栏目id"  gorm:"column:cat_id;comment:栏目id;type:bigint"`
	ArtId  int64 `json:"artId" form:"artId" cn:"文章id"  gorm:"column:art_id;comment:文章id;type:bigint"`
	BeHot  bool  `json:"beHot" form:"beHot" cn:"热门"   gorm:"column:be_hot;comment:beHot;type:tinyint"`
	Sort   int   `json:"sort" form:"sort" cn:"排序"   gorm:"column:sort;comment:排序;type:int"`
	Status int   `json:"status" form:"status" cn:"状态"   gorm:"column:status;comment:状态:0未审核 1审核 2未通过审核 3草稿;type:smallint"`
	//
	Title string `json:"title" form:"title" cn:"文章标题"  gorm:"column:title;<-:false"`
}

// TableName CmsCatArt 表名
func (CmsCatArt) TableName() string {
	return "cms_cat_art"
}

type CmsCatArtSearch struct {
	request.PageInfo
	global.BaseModel
	UserId *int64 `json:"userId"  form:"userId" `
	CatId  *int64 `json:"catId"  form:"catId" `
	ArtId  *int64 `json:"artId"  form:"artId" `
	BeHot  *bool  `json:"beHot" form:"beHot"`
	Sort   *int   `json:"sort" form:"sort" `
	Status *int   `json:"status"   form:"status" `
}
