// 自动生成模板BbPiDepartment
package business

import (
	"go-cms/global"
)

// BbPiDepartment 结构体
// 如果含有time.Time 请自行import time包
type BbPiDepartment struct {
     BbPiDepartmentMini
}

// BbPiDepartmentMini 结构体
type BbPiDepartmentMini struct {
      global.GO_MODEL
      Status  *int `json:"status" cn:"上传" form:"status" gorm:"column:status;comment:上传:0未上传 1已上传 2上传失败;type:smallint"`
      Jgdm  string `json:"jgdm" cn:"机构标识" form:"jgdm" gorm:"column:jgdm;comment:机构标识;type:varchar(200);"`
      Ksbm  string `json:"ksbm" cn:"科室编码" form:"ksbm" gorm:"column:ksbm;comment:科室编码;type:varchar(200);"`
      Ksmc  string `json:"ksmc" cn:"科室名称" form:"ksmc" gorm:"column:ksmc;comment:科室名称;type:varchar(200);"`
      Bzksdm  string `json:"bzksdm" cn:"标准科室代码" form:"bzksdm" gorm:"column:bzksdm;comment:标准科室代码;type:varchar(200);"`
      Ybjdm  string `json:"ybjdm" cn:"医保局代码" form:"ybjdm" gorm:"column:ybjdm;comment:医保局代码;type:varchar(200);"`
      Ksjs  string `json:"ksjs" cn:"科室简介" form:"ksjs" gorm:"column:ksjs;comment:科室简介;type:varchar(200);"`
      Sjscsj  string `json:"sjscsj" cn:"数据生成日期时间" form:"sjscsj" gorm:"column:sjscsj;comment:数据生成日期时间;type:varchar(200);"`
      Tbrqsj  string `json:"tbrqsj" cn:"填报日期" form:"tbrqsj" gorm:"column:tbrqsj;comment:填报日期;type:varchar(200);"`
      Cxbz  string `json:"cxbz" cn:"撤销标志" form:"cxbz" gorm:"column:cxbz;comment:撤销标志;type:varchar(200);"`

}


// TableName BbPiDepartment 表名
func (BbPiDepartment) TableName() string {
  return "bb_pi_department"
}

