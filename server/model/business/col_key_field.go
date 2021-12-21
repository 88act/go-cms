// 自动生成模板ColKeyField
package business

import (
	"go-cms/global"
)

// ColKeyField 结构体
// 如果含有time.Time 请自行import time包
type ColKeyField struct {
	ColKeyFieldMini
}

// ColKeyFieldMini 结构体
type ColKeyFieldMini struct {
	global.GO_MODEL
	ColId   *int   `json:"colId" cn:"采集id" form:"colId" gorm:"column:col_id;comment:采集id;type:int"`
	ToKey   string `json:"toKey" cn:"标识" form:"toKey" gorm:"column:to_key;comment:标识;type:varchar(200);"`
	ToField string `json:"toField" cn:"字段" form:"toField" gorm:"column:to_field;comment:字段;type:varchar(100);"`
	Status  *int   `json:"status" cn:"状态" form:"status" gorm:"column:status;comment:状态:0未审核 1有效  2未通过审核 3草稿;type:smallint"`
}

// TableName ColKeyField 表名
func (ColKeyField) TableName() string {
	return "col_key_field"
}
