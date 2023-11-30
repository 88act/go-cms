// 模板CmsBlog
// Author 10512203@qq.com
package business

import (
	"go-cms/global"
	"go-cms/model/common/request"
)

// 表格字段
const Field_CmsBlog_mini = "" // "id,created_at,updated_at,user_id,title,desc,image,file_list,total_whole,status,

// CmsBlog 结构体
type CmsBlog struct {
	global.BaseModel
	UserId     int64  `json:"userId" form:"userId" cn:"用户id"  gorm:"column:user_id;comment:用户id;type:bigint"`
	Title      string `json:"title" form:"title" cn:"标题"  gorm:"column:title;comment:标题;type:varchar(150);"`
	Desc       string `json:"desc" form:"desc" cn:"简介"  gorm:"column:desc;comment:简介;type:varchar(500);"`
	Image      string `json:"image"  form:"image"  cn:"缩略图" gorm:"column:image;comment:缩略图;type:单图片"`
	FileList   string `json:"fileList"  form:"fileList"  cn:"文件列表" gorm:"column:file_list;comment:文件列表;type:多图片"`
	Email      string `json:"email" form:"email" cn:"邮件"  gorm:"column:email;comment:邮件;type:varchar(60);"`
	AreaId     int    `json:"areaId" form:"areaId" cn:"地区id"   gorm:"column:area_id;comment:地区id;type:int"`
	Address    string `json:"address" form:"address" cn:"地址"  gorm:"column:address;comment:地址;type:varchar(500);"`
	TotalWhole int    `json:"totalWhole" form:"totalWhole" cn:"综合指数"   gorm:"column:total_whole;comment:综合指数;type:int"`
	TotalClick int    `json:"totalClick" form:"totalClick" cn:"总点击"   gorm:"column:total_click;comment:总点击;type:int"`
	TotalFan   int    `json:"totalFan" form:"totalFan" cn:"总粉丝"   gorm:"column:total_fan;comment:总粉丝;type:int"`
	TotalGood  int    `json:"totalGood" form:"totalGood" cn:"总赞"   gorm:"column:total_good;comment:总赞;type:int"`
	TotalPoor  int    `json:"totalPoor" form:"totalPoor" cn:"总踩"   gorm:"column:total_poor;comment:总踩;type:int"`
	Status     int    `json:"status" form:"status" cn:"状态"   gorm:"column:status;comment:状态:0未审核 1审核通过 2未通过审核 3黑名单;type:smallint"`
}

// TableName CmsBlog 表名
func (CmsBlog) TableName() string {
	return "cms_blog"
}

type CmsBlogSearch struct {
	request.PageInfo
	global.BaseModel
	UserId     *int64 `json:"userId"  form:"userId" `
	Title      string `json:"title"   form:"title" `
	Desc       string `json:"desc"   form:"desc" `
	Image      string `json:"image" form:"image" `
	FileList   string `json:"fileList" form:"fileList" `
	Email      string `json:"email"   form:"email" `
	AreaId     *int   `json:"areaId"   form:"areaId" `
	Address    string `json:"address"   form:"address" `
	TotalWhole *int   `json:"totalWhole"   form:"totalWhole" `
	TotalClick *int   `json:"totalClick"   form:"totalClick" `
	TotalFan   *int   `json:"totalFan"   form:"totalFan" `
	TotalGood  *int   `json:"totalGood"   form:"totalGood" `
	TotalPoor  *int   `json:"totalPoor"   form:"totalPoor" `
	Status     *int   `json:"status"   form:"status" `
}
