// 自动生成模板BbPiUpField
package business

import (
	"go-cms/global"
)

// BbPiUpField 结构体
// 如果含有time.Time 请自行import time包
type BbPiUpField struct {
     BbPiUpFieldMini
}

// BbPiUpFieldMini 结构体
type BbPiUpFieldMini struct {
      global.GO_MODEL
      Status  *int `json:"status" cn:"是否上传" form:"status" gorm:"column:status;comment:是否上传:0不传 1上传;type:smallint"`
      TabName  string `json:"tabName" cn:"表名" form:"tabName" gorm:"column:tab_name;comment:表名;type:varchar(200);"`
      TabNameCn  string `json:"tabNameCn" cn:"表名cn" form:"tabNameCn" gorm:"column:tab_name_cn;comment:表名cn;type:varchar(200);"`
      TabField  string `json:"tabField" cn:"字段名" form:"tabField" gorm:"column:tab_field;comment:字段名;type:varchar(200);"`
      TabFieldCn  string `json:"tabFieldCn" cn:"字段名cn" form:"tabFieldCn" gorm:"column:tab_field_cn;comment:字段名cn;type:varchar(200);"`
      Sort  *int `json:"sort" cn:"排序" form:"sort" gorm:"column:sort;comment:排序;type:int"`

}


// TableName BbPiUpField 表名
func (BbPiUpField) TableName() string {
  return "bb_pi_up_field"
}

