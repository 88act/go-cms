// 自动生成模板CmsCat
package model

import (
	. "go-cms/common/baseModel"
)

// CmsCat 结构体
type CmsCat struct {
	BaseModel
	Pid      int    `db:"pid" json:"pid" gorm:"column:pid"`                  //父ID
	UserId   int64  `db:"user_id" json:"userId" gorm:"column:user_id"`       //用户id
	BeSys    bool   `db:"be_sys" json:"beSys" gorm:"column:be_sys"`          //系统分类
	ObjId    int64  `db:"obj_id" json:"objId" gorm:"column:obj_id"`          //专题id
	Type     int    `db:"type" json:"type" gorm:"column:type"`               //栏目类型
	Title    string `db:"title" json:"title" gorm:"column:title"`            //标题
	Desc     string `db:"desc" json:"desc" gorm:"column:desc"`               //摘要
	TagList  string `db:"tag_list" json:"tagList" gorm:"column:tag_list"`    //标签列表
	Image    string `db:"image" json:"image" gorm:"column:image"`            //插图
	FileList string `db:"file_list" json:"fileList" gorm:"column:file_list"` //媒体列表
	BeTop    bool   `db:"be_top" json:"beTop" gorm:"column:be_top"`          //置顶
	BeNav    bool   `db:"be_nav" json:"beNav" gorm:"column:be_nav"`          //是否导航
	Sort     int    `db:"sort" json:"sort" gorm:"column:sort"`               //排序
	Status   int    `db:"status" json:"status" gorm:"column:status"`         //状态

}

// TableName CmsCat 表名
func (CmsCat) TableName() string {
	return "cms_cat"
}

// CmsCatSearch  查询
type CmsCatSearch struct {
	BaseModel
	PageInfo
	Pid      *int   `json:"pid"   form:"pid" `
	UserId   *int64 `json:"userId"  form:"userId" `
	BeSys    *bool  `json:"beSys" form:"beSys" `
	ObjId    *int64 `json:"objId"  form:"objId" `
	Type     *int   `json:"type"   form:"type" `
	Title    string `json:"title"   form:"title" `
	Desc     string `json:"desc"   form:"desc" `
	TagList  string `json:"tagList"   form:"tagList" `
	Image    string `json:"image" form:"image" `
	FileList string `json:"fileList" form:"fileList" `
	BeTop    *bool  `json:"beTop" form:"beTop" `
	BeNav    *bool  `json:"beNav" form:"beNav" `
	Sort     *int   `json:"sort"   form:"sort" `
	Status   *int   `json:"status"   form:"status" `
}
