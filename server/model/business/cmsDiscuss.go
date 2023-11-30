// 模板CmsDiscuss
// Author 10512203@qq.com
package business

import (
	"go-cms/global"
	"go-cms/model/common/request"
)

// 表格字段
const Field_CmsDiscuss_mini = "" // "id,created_at,updated_at,art_id,user_id,user_id_at,title,file_list,total_good,total_poor,status,

// CmsDiscuss 结构体
type CmsDiscuss struct {
	global.BaseModel
	Pid       int64  `json:"pid" form:"pid" cn:"父ID"  gorm:"column:pid;comment:父ID;type:bigint"`
	ArtId     int64  `json:"artId" form:"artId" cn:"id"  gorm:"column:art_id;comment:id;type:bigint"`
	UserId    int64  `json:"userId" form:"userId" cn:"用户id"  gorm:"column:user_id;comment:用户id;type:bigint"`
	UserIdAt  int64  `json:"userIdAt" form:"userIdAt" cn:"用户At"  gorm:"column:user_id_at;comment:用户At;type:bigint"`
	Title     string `json:"title" form:"title" cn:"标题"  gorm:"column:title;comment:标题;type:varchar(100);"`
	Detail    string `json:"detail" form:"detail" cn:"内容"  gorm:"column:detail;comment:内容;type:varchar(4000);"`
	FileList  string `json:"fileList"  form:"fileList"  cn:"图片" gorm:"column:file_list;comment:图片;type:多图片"`
	TotalGood int    `json:"totalGood" form:"totalGood" cn:"总赞"   gorm:"column:total_good;comment:总赞;type:int"`
	TotalPoor int    `json:"totalPoor" form:"totalPoor" cn:"总踩"   gorm:"column:total_poor;comment:总踩;type:int"`
	Status    int    `json:"status" form:"status" cn:"状态"   gorm:"column:status;comment:状态:0未审核 1审核通过 2未通过审核;type:smallint"`
}

// TableName CmsDiscuss 表名
func (CmsDiscuss) TableName() string {
	return "cms_discuss"
}

type CmsDiscussSearch struct {
	request.PageInfo
	global.BaseModel
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
