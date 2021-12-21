// 自动生成模板CmsCat
package business

import (
	"go-cms/global"
)

// CmsCat 结构体
// 如果含有time.Time 请自行import time包
type CmsCat struct {
     CmsCatMini
      Desc  string `json:"desc" cn:"描述" form:"desc" gorm:"column:desc;comment:描述;type:text;"`
      Keywords  string `json:"keywords" cn:"关键词" form:"keywords" gorm:"column:keywords;comment:关键词;type:longtext;"`
      Alias  string `json:"alias" cn:"别名" form:"alias" gorm:"column:alias;comment:别名;type:varchar(256);"`
}

// CmsCatMini 结构体
type CmsCatMini struct {
      global.GO_MODEL
      Pid  *int `json:"pid" cn:"父ID" form:"pid" gorm:"column:pid;comment:父ID;type:int"`
      BeSys  *bool `json:"beSys" cn:"系统分类" form:"beSys" gorm:"column:be_sys;comment:系统分类: 0=否 1=是;type:tinyint"`
      GroupId  *int `json:"groupId" cn:"群组id" form:"groupId" gorm:"column:group_id;comment:群组id;type:bigint"`
      MediaType  *int `json:"mediaType" cn:"文章类型" form:"mediaType" gorm:"column:media_type;comment:文章类型;type:smallint"`
      Name  string `json:"name" cn:"名称" form:"name" gorm:"column:name;comment:名称:（栏目/专题）;type:varchar(100);"`
      Thumb  string `json:"thumb" cn:"配图" form:"thumb" gorm:"column:thumb;comment:配图;type:单图片"`
      Sort  *int `json:"sort" cn:"排序" form:"sort" gorm:"column:sort;comment:排序;type:int"`
      BeNav  *bool `json:"beNav" cn:"是否导航" form:"beNav" gorm:"column:be_nav;comment:是否导航;type:tinyint"`
      Status  *int `json:"status" cn:"状态" form:"status" gorm:"column:status;comment:状态:0未审核 1审核  2未通过审核 3 草稿;type:smallint"`

}


// TableName CmsCat 表名
func (CmsCat) TableName() string {
  return "cms_cat"
}

