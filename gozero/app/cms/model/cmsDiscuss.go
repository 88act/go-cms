// 自动生成模板CmsDiscuss
package model

import (
	. "go-cms/common/baseModel"
)

// CmsDiscuss 结构体
type CmsDiscuss struct {
	BaseModel
	Pid       int64  `db:"pid" json:"pid" gorm:"column:pid"`                     //父ID
	ArtId     int64  `db:"art_id" json:"artId" gorm:"column:art_id"`             //id
	UserId    int64  `db:"user_id" json:"userId" gorm:"column:user_id"`          //用户id
	UserIdAt  int64  `db:"user_id_at" json:"userIdAt" gorm:"column:user_id_at"`  //用户At
	Title     string `db:"title" json:"title" gorm:"column:title"`               //标题
	Detail    string `db:"detail" json:"detail" gorm:"column:detail"`            //内容
	FileList  string `db:"file_list" json:"fileList" gorm:"column:file_list"`    //图片
	TotalGood int    `db:"total_good" json:"totalGood" gorm:"column:total_good"` //总赞
	TotalPoor int    `db:"total_poor" json:"totalPoor" gorm:"column:total_poor"` //总踩
	Status    int    `db:"status" json:"status" gorm:"column:status"`            //状态

}

// TableName CmsDiscuss 表名
func (CmsDiscuss) TableName() string {
	return "cms_discuss"
}

// CmsDiscussSearch  查询
type CmsDiscussSearch struct {
	BaseModel
	PageInfo
	Pid       *int64 `json:"pid"  form:"pid" `
	ArtId     *int64 `json:"artId"  form:"artId" `
	UserId    *int64 `json:"userId"  form:"userId" `
	UserIdAt  *int64 `json:"userIdAt"  form:"userIdAt" `
	Title     string `json:"title"   form:"title" `
	Detail    string `json:"detail"   form:"detail" `
	FileList  string `json:"fileList" form:"fileList" `
	TotalGood *int   `json:"totalGood"   form:"totalGood" `
	TotalPoor *int   `json:"totalPoor"   form:"totalPoor" `
	Status    *int   `json:"status"   form:"status" `
}
