// 模板CmsGroup
// Author 10512203@qq.com
package business

import (
	"go-cms/global"
	"go-cms/model/common/request"
)

// 表格字段
const Field_CmsGroup_mini = "" // "id,created_at,updated_at,user_id,type,title,image,status,

// CmsGroup 结构体
type CmsGroup struct {
	global.BaseModel
	UserId   int64  `json:"userId" form:"userId" cn:"用户id"  gorm:"column:user_id;comment:用户id;type:bigint"`
	Type     int    `json:"type" form:"type" cn:"群组类别"   gorm:"column:type;comment:群组类别: 1音乐 2舞蹈 3书法 4 棋艺;type:int"`
	Title    string `json:"title" form:"title" cn:"名称"  gorm:"column:title;comment:名称;type:varchar(100);"`
	Desc     string `json:"desc" form:"desc" cn:"内容"  gorm:"column:desc;comment:内容;type:varchar(2000);"`
	TagList  string `json:"tagList" form:"tagList" cn:"标签"  gorm:"column:tag_list;comment:标签;type:varchar(500);"`
	Image    string `json:"image"  form:"image"  cn:"插图" gorm:"column:image;comment:插图;type:单图片"`
	FileList string `json:"fileList"  form:"fileList"  cn:"文件" gorm:"column:file_list;comment:文件;type:多图片"`
	Status   int    `json:"status" form:"status" cn:"状态"   gorm:"column:status;comment:状态:0未审核 1审核 2未通过审核 3 草稿;type:smallint"`
}

// TableName CmsGroup 表名
func (CmsGroup) TableName() string {
	return "cms_group"
}

type CmsGroupSearch struct {
	request.PageInfo
	global.BaseModel
	UserId   *int64 `json:"userId"  form:"userId" `
	Type     *int   `json:"type"   form:"type" `
	Title    string `json:"title"   form:"title" `
	Desc     string `json:"desc"   form:"desc" `
	TagList  string `json:"tagList"   form:"tagList" `
	Image    string `json:"image" form:"image" `
	FileList string `json:"fileList" form:"fileList" `
	Status   *int   `json:"status"   form:"status" `
}
