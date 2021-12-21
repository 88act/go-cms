// 自动生成模板ColCollect
package business

import (
	"go-cms/global"
)

// ColCollect 结构体
// 如果含有time.Time 请自行import time包
type ColCollect struct {
	ColCollectMini
	UrlPage string `json:"urlPage" cn:"翻页路径" form:"urlPage" gorm:"column:url_page;comment:翻页路径;type:varchar(256);"`
	ToTable string `json:"toTable" cn:"对应表" form:"toTable" gorm:"column:to_table;comment:对应表;type:varchar(100);"`
	NowId   string `json:"nowId" cn:"当前id" form:"nowId" gorm:"column:now_id;comment:当前id;type:varchar(100);"`
	PageNow *int   `json:"pageNow" cn:"当前页码" form:"pageNow" gorm:"column:page_now;comment:当前页码;type:int"`
}

// ColCollectMini 结构体
type ColCollectMini struct {
	global.GO_MODEL
	Userid    *int   `json:"userid" cn:"用户id" form:"userid" gorm:"column:userid;comment:id;type:int"`
	Name      string `json:"name" cn:"名称" form:"name" gorm:"column:name;comment:名称;type:varchar(200);"`
	Url       string `json:"url" cn:"路径" form:"url" gorm:"column:url;comment:路径;type:varchar(256);"`
	PageStart *int   `json:"pageStart" cn:"开始页码" form:"pageStart" gorm:"column:page_start;comment:开始页码;type:int"`
	PageEnd   *int   `json:"pageEnd" cn:"结束页码" form:"pageEnd" gorm:"column:page_end;comment:结束页码;type:int"`
	StatusRun *int   `json:"statusRun" cn:"运行状态" form:"statusRun" gorm:"column:status_run;comment:运行状态:0停止 1采集中;type:smallint"`
	Status    *int   `json:"status" cn:"状态" form:"status" gorm:"column:status;comment:状态:0未审核 1有效  2未通过审核 3草稿;type:smallint"`
}

// TableName ColCollect 表名
func (ColCollect) TableName() string {
	return "col_collect"
}
