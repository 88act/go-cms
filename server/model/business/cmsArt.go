// 模板CmsArt
// Author 10512203@qq.com
package business

import (
	"go-cms/global"
	"go-cms/model/common/request"
)

// 表格字段
const Field_CmsArt_mini = "" // "id,created_at,updated_at,user_id,cat_id,cat_id_sys,type,title,tag_list,image,be_top,total_whole,total_share,total_click,status,

// CmsArt 结构体
type CmsArt struct {
	global.BaseModel
	Pid          int64  `json:"pid" form:"pid" cn:"父id"  gorm:"column:pid;comment:父id:章节集合;type:bigint"`
	UserId       int64  `json:"userId" form:"userId" cn:"用户id"  gorm:"column:user_id;comment:用户id;type:bigint"`
	Type         int    `json:"type" form:"type" cn:"文章类型"   gorm:"column:type;comment:文章类型;type:smallint"`
	Title        string `json:"title" form:"title" cn:"文章标题"  gorm:"column:title;comment:文章标题;type:varchar(200);"`
	Desc         string `json:"desc" form:"desc" cn:"文章摘要"  gorm:"column:desc;comment:文章摘要;type:varchar(500);"`
	TagList      string `json:"tagList" form:"tagList" cn:"标签列表"  gorm:"column:tag_list;comment:标签列表;type:varchar(500);"`
	Source       string `json:"source" form:"source" cn:"来源"  gorm:"column:source;comment:来源;type:varchar(255);"`
	Image        string `json:"image"  form:"image"  cn:"插图" gorm:"column:image;comment:插图;type:单图片"`
	FileList     string `json:"fileList"  form:"fileList"  cn:"媒体列表" gorm:"column:file_list;comment:媒体列表;type:多图片"`
	Link         string `json:"link" form:"link" cn:"链接地址"  gorm:"column:link;comment:链接地址;type:varchar(255);"`
	TotalWhole   int    `json:"totalWhole" form:"totalWhole" cn:"综合指数"   gorm:"column:total_whole;comment:综合指数;type:int"`
	TotalShare   int    `json:"totalShare" form:"totalShare" cn:"总分享"   gorm:"column:total_share;comment:总分享;type:int"`
	TotalFav     int    `json:"totalFav" form:"totalFav" cn:"总收藏"   gorm:"column:total_fav;comment:总收藏;type:int"`
	TotalDiscuss int    `json:"totalDiscuss" form:"totalDiscuss" cn:"总评论"   gorm:"column:total_discuss;comment:总评论;type:int"`
	TotalClick   int    `json:"totalClick" form:"totalClick" cn:"总点击"   gorm:"column:total_click;comment:总点击;type:int"`
	TotalStar    int    `json:"totalStar" form:"totalStar" cn:"总星"   gorm:"column:total_star;comment:总星;type:int"`
	TotalGood    int    `json:"totalGood" form:"totalGood" cn:"总赞"   gorm:"column:total_good;comment:总赞;type:int"`
	TotalPoor    int    `json:"totalPoor" form:"totalPoor" cn:"总踩"   gorm:"column:total_poor;comment:总踩;type:int"`
	Status       int    `json:"status" form:"status" cn:"状态"   gorm:"column:status;comment:状态:0未审核 1审核通过 2未通过审核 3草稿;type:smallint"`
	VerifyMsg    string `json:"verifyMsg" form:"verifyMsg" cn:"审核信息"  gorm:"column:verify_msg;comment:审核信息;type:varchar(255);"`
}

// TableName CmsArt 表名
func (CmsArt) TableName() string {
	return "cms_art"
}

type CmsArtSearch struct {
	request.PageInfo
	global.BaseModel
	Pid          *int64 `json:"pid"  form:"pid" `
	UserId       *int64 `json:"userId"  form:"userId" `
	Type         *int   `json:"type"   form:"type" `
	Title        string `json:"title"   form:"title" `
	Desc         string `json:"desc"   form:"desc" `
	TagList      string `json:"tagList"   form:"tagList" `
	Source       string `json:"source"   form:"source" `
	Image        string `json:"image" form:"image" `
	FileList     string `json:"fileList" form:"fileList" `
	Link         string `json:"link"   form:"link" `
	TotalWhole   *int   `json:"totalWhole"   form:"totalWhole" `
	TotalShare   *int   `json:"totalShare"   form:"totalShare" `
	TotalFav     *int   `json:"totalFav"   form:"totalFav" `
	TotalDiscuss *int   `json:"totalDiscuss"   form:"totalDiscuss" `
	TotalClick   *int   `json:"totalClick"   form:"totalClick" `
	TotalStar    *int   `json:"totalStar"   form:"totalStar" `
	TotalGood    *int   `json:"totalGood"   form:"totalGood" `
	TotalPoor    *int   `json:"totalPoor"   form:"totalPoor" `
	Status       *int   `json:"status"   form:"status" `
	VerifyMsg    string `json:"verifyMsg"   form:"verifyMsg" `
}
