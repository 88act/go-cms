// 自动生成模板BasicArea
package business

import (
	"go-cms/global"
)

// BasicArea 结构体
// 如果含有time.Time 请自行import time包
type BasicArea struct {
	BasicAreaMini
}

// BasicAreaMini 结构体
type BasicAreaMini struct {
	global.GO_MODEL
	Areaname   string `json:"areaname" cn:"名称" form:"areaname" gorm:"column:areaname;comment:名称;type:varchar(50);"`
	Parentid   *int   `json:"parentid" cn:"父栏目" form:"parentid" gorm:"column:parentid;comment:父栏目;type:int"`
	Shortname  string `json:"shortname" cn:"简称" form:"shortname" gorm:"column:shortname;comment:简称;type:varchar(50);"`
	Areacode   string `json:"areacode" cn:"电话区号" form:"areacode" gorm:"column:areacode;comment:电话区号;type:varchar(10);"`
	Zipcode    string `json:"zipcode" cn:"邮编" form:"zipcode" gorm:"column:zipcode;comment:邮编;type:varchar(10);"`
	Pinyin     string `json:"pinyin" cn:"拼音" form:"pinyin" gorm:"column:pinyin;comment:拼音;type:varchar(100);"`
	Py         string `json:"py" cn:"拼音简写" form:"py" gorm:"column:py;comment:拼音简写;type:varchar(10);"`
	Lng        string `json:"lng" cn:"经度" form:"lng" gorm:"column:lng;comment:经度;type:varchar(20);"`
	Lat        string `json:"lat" cn:"纬度" form:"lat" gorm:"column:lat;comment:纬度;type:varchar(20);"`
	Level      *int   `json:"level" cn:"级别" form:"level" gorm:"column:level;comment:级别;type:smallint"`
	Position   string `json:"position" cn:"position字段" form:"position" gorm:"column:position;comment:;type:varchar(255);"`
	Mergername string `json:"mergername" cn:"组合名称" form:"mergername" gorm:"column:mergername;comment:组合名称;type:varchar(255);"`
	Sort       *int   `json:"sort" cn:"排序" form:"sort" gorm:"column:sort;comment:排序;type:int"`
}

// TableName BasicArea 表名
func (BasicArea) TableName() string {
	return "basic_area"
}
