// 自动生成模板CmsArt
package model

import (
	. "go-cms/common/baseModel"
)

// CmsArt 结构体
type CmsArt struct {
	BaseModel
	Pid          int64  `db:"pid" json:"pid" gorm:"column:pid"`                              //父id
	UserId       int64  `db:"user_id" json:"userId" gorm:"column:user_id"`                   //用户id
	CatId        int64  `db:"cat_id" json:"catId" gorm:"column:cat_id"`                      //类别
	CatIdSys     int64  `db:"cat_id_sys" json:"catIdSys" gorm:"column:cat_id_sys"`           //系统类别
	Type         int    `db:"type" json:"type" gorm:"column:type"`                           //文章类型
	Title        string `db:"title" json:"title" gorm:"column:title"`                        //文章标题
	Desc         string `db:"desc" json:"desc" gorm:"column:desc"`                           //文章摘要
	TagList      string `db:"tag_list" json:"tagList" gorm:"column:tag_list"`                //标签列表
	Source       string `db:"source" json:"source" gorm:"column:source"`                     //来源
	Image        string `db:"image" json:"image" gorm:"column:image"`                        //插图
	FileList     string `db:"file_list" json:"fileList" gorm:"column:file_list"`             //媒体列表
	Link         string `db:"link" json:"link" gorm:"column:link"`                           //链接地址
	BeTop        bool   `db:"be_top" json:"beTop" gorm:"column:be_top"`                      //置顶
	TotalWhole   int    `db:"total_whole" json:"totalWhole" gorm:"column:total_whole"`       //综合指数
	TotalShare   int    `db:"total_share" json:"totalShare" gorm:"column:total_share"`       //总分享
	TotalFav     int    `db:"total_fav" json:"totalFav" gorm:"column:total_fav"`             //总收藏
	TotalDiscuss int    `db:"total_discuss" json:"totalDiscuss" gorm:"column:total_discuss"` //总评论
	TotalClick   int    `db:"total_click" json:"totalClick" gorm:"column:total_click"`       //总点击
	TotalStar    int    `db:"total_star" json:"totalStar" gorm:"column:total_star"`          //总星
	TotalGood    int    `db:"total_good" json:"totalGood" gorm:"column:total_good"`          //总赞
	TotalPoor    int    `db:"total_poor" json:"totalPoor" gorm:"column:total_poor"`          //总踩
	Status       int    `db:"status" json:"status" gorm:"column:status"`                     //状态
	VerifyMsg    string `db:"verify_msg" json:"verifyMsg" gorm:"column:verify_msg"`          //审核信息

}

// TableName CmsArt 表名
func (CmsArt) TableName() string {
	return "cms_art"
}

// CmsArtSearch  查询
type CmsArtSearch struct {
	BaseModel
	PageInfo
	Pid          *int64 `json:"pid"  form:"pid" `
	UserId       *int64 `json:"userId"  form:"userId" `
	CatId        *int64 `json:"catId"  form:"catId" `
	CatIdSys     *int64 `json:"catIdSys"  form:"catIdSys" `
	Type         *int   `json:"type"   form:"type" `
	Title        string `json:"title"   form:"title" `
	Desc         string `json:"desc"   form:"desc" `
	TagList      string `json:"tagList"   form:"tagList" `
	Source       string `json:"source"   form:"source" `
	Image        string `json:"image" form:"image" `
	FileList     string `json:"fileList" form:"fileList" `
	Link         string `json:"link"   form:"link" `
	BeTop        *bool  `json:"beTop" form:"beTop" `
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
