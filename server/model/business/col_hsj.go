// 自动生成模板ColHsj
package business

import (
	"time"

	"go-cms/global"
)

// ColHsj 结构体
// 如果含有time.Time 请自行import time包
type ColHsj struct {
	ColHsjMini
	Content string `json:"content" cn:"内容" form:"content" gorm:"column:content;comment:内容;type:text;"`
}

// ColHsjMini 结构体
type ColHsjMini struct {
	global.GO_MODEL
	ColId   *int       `json:"colId" cn:"采集id" form:"colId" gorm:"column:col_id;comment:采集id;type:int"`
	Url     string     `json:"url" cn:"原URL" form:"url" gorm:"column:url;comment:原URL;type:varchar(255);"`
	Title   string     `json:"title" cn:"标题" form:"title" gorm:"column:title;comment:标题;type:varchar(200);"`
	PubTime *time.Time `json:"pubTime" cn:"发布时间" form:"pubTime" gorm:"column:pub_time;comment:发布时间;type:datetime"`
	PubUnit string     `json:"pubUnit" cn:"发布单位" form:"pubUnit" gorm:"column:pub_unit;comment:发布单位;type:varchar(200);"`
	Log     string     `json:"log" cn:"经度" form:"log" gorm:"column:log;comment:经度;type:varchar(100);"`
	Lat     string     `json:"lat" cn:"纬度" form:"lat" gorm:"column:lat;comment:纬度;type:varchar(100);"`
	Status  *int       `json:"status" cn:"状态" form:"status" gorm:"column:status;comment:状态:0未审核 1有效  2未通过审核 3草稿;type:smallint"`
}

// TableName ColHsj 表名
func (ColHsj) TableName() string {
	return "col_hsj"
}
