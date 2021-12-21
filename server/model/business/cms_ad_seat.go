// 自动生成模板CmsAdSeat
package business

import (
	"go-cms/global"
)

// CmsAdSeat 结构体
// 如果含有time.Time 请自行import time包
type CmsAdSeat struct {
     CmsAdSeatMini
      Desc  string `json:"desc" cn:"广告描述" form:"desc" gorm:"column:desc;comment:广告描述;type:varchar(500);"`
}

// CmsAdSeatMini 结构体
type CmsAdSeatMini struct {
      global.GO_MODEL
      Name  string `json:"name" cn:"位置名称" form:"name" gorm:"column:name;comment:位置名称;type:varchar(60);"`
      Width  *int `json:"width" cn:"广告位宽度" form:"width" gorm:"column:width;comment:广告位宽度;type:smallint"`
      Height  *int `json:"height" cn:"广告位高度" form:"height" gorm:"column:height;comment:广告位高度;type:smallint"`
      Style  string `json:"style" cn:"模板" form:"style" gorm:"column:style;comment:模板;type:text;"`
      Status  *int `json:"status" cn:"状态" form:"status" gorm:"column:status;comment:状态:0关闭1开启;type:smallint"`

}


// TableName CmsAdSeat 表名
func (CmsAdSeat) TableName() string {
  return "cms_ad_seat"
}

