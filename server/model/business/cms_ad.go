// 自动生成模板CmsAd
package business

import (
	"go-cms/global"
	"time"
)

// CmsAd 结构体
// 如果含有time.Time 请自行import time包
type CmsAd struct {
	CmsAdMini
	Detail    string `json:"detail" cn:"广告内容" form:"detail" gorm:"column:detail;comment:广告内容;type:varchar(2000);"`
	ExpDetail string `json:"expDetail" cn:"过期内容" form:"expDetail" gorm:"column:exp_detail;comment:过期内容;type:varchar(2000);"`
}

// CmsAdMini 结构体
type CmsAdMini struct {
	global.GO_MODEL
	SeatId     *int       `json:"seatId" cn:"广告位置" form:"seatId" gorm:"column:seat_id;comment:广告位置;type:int"`
	Name       string     `json:"name" cn:"广告名称" form:"name" gorm:"column:name;comment:广告名称;type:varchar(500);"`
	MediaType  *int       `json:"mediaType" cn:"广告类型" form:"mediaType" gorm:"column:media_type;comment:广告类型:0=文本，1=图片，2=视频，3=flash，4=音频;type:smallint"`
	BeTarget   *bool      `json:"beTarget" cn:"是否新窗" form:"beTarget" gorm:"column:be_target;comment:是否新窗: 0=新窗体 1=本页;type:tinyint"`
	Image      string     `json:"image" cn:"广告图片" form:"image" gorm:"column:image;comment:广告图片;type:单图片"`
	Link       string     `json:"link" cn:"广告链接" form:"link" gorm:"column:link;comment:广告链接;type:varchar(255);"`
	StartTime  *time.Time `json:"startTime" cn:"投放时间" form:"startTime" gorm:"column:start_time;comment:投放时间;type:datetime"`
	EndTime    *time.Time `json:"endTime" cn:"结束时间" form:"endTime" gorm:"column:end_time;comment:结束时间;type:datetime"`
	TotalClick *int       `json:"totalClick" cn:"点击量" form:"totalClick" gorm:"column:total_click;comment:点击量;type:int"`
	Sort       *int       `json:"sort" cn:"排序" form:"sort" gorm:"column:sort;comment:排序;type:int"`
	Status     *int       `json:"status" cn:"状态" form:"status" gorm:"column:status;comment:状态:0未审核 1审核 2未通过审核 3 草稿;type:smallint"`
}

// TableName CmsAd 表名
func (CmsAd) TableName() string {
	return "cms_ad"
}
