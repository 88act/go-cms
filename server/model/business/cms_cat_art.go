// 自动生成模板CmsCatArt
package business

import (
	"go-cms/global"
)

// CmsCatArt 结构体
// 如果含有time.Time 请自行import time包
type CmsCatArt struct {
	global.GO_MODEL
	ArticleId *bool `json:"articleId" form:"articleId" gorm:"column:article_id;comment:文章id;type:tinyint"`
	CatId     *int  `json:"catId" form:"catId" gorm:"column:cat_id;comment:栏目id;type:int"`
	Status    *int  `json:"status" form:"status" gorm:"column:status;comment:状态:0未审核 1审核 2未通过审核 3草稿;type:smallint"`
}

// TableName CmsCatArt 表名
func (CmsCatArt) TableName() string {
	return "cms_cat_art"
}
