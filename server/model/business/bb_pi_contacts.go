// 自动生成模板BbPiContacts
package business

import (
	"go-cms/global"
)

// BbPiContacts 结构体
// 如果含有time.Time 请自行import time包
type BbPiContacts struct {
     BbPiContactsMini
}

// BbPiContactsMini 结构体
type BbPiContactsMini struct {
      global.GO_MODEL
      Status  *int `json:"status" cn:"上传" form:"status" gorm:"column:status;comment:上传:0未上传 1已上传 2上传失败;type:smallint"`
      Jgdm  string `json:"jgdm" cn:"机构标识" form:"jgdm" gorm:"column:jgdm;comment:机构标识;type:varchar(200);"`
      Bm  string `json:"bm" cn:"部门" form:"bm" gorm:"column:bm;comment:部门;type:varchar(200);"`
      Fzrxm  string `json:"fzrxm" cn:"负责人姓名" form:"fzrxm" gorm:"column:fzrxm;comment:负责人姓名;type:varchar(200);"`
      Lxdh  string `json:"lxdh" cn:"联系电话" form:"lxdh" gorm:"column:lxdh;comment:联系电话;type:varchar(200);"`
      Sjscsj  string `json:"sjscsj" cn:"数据生成日期时间" form:"sjscsj" gorm:"column:sjscsj;comment:数据生成日期时间;type:varchar(200);"`
      Tbrqsj  string `json:"tbrqsj" cn:"填报日期" form:"tbrqsj" gorm:"column:tbrqsj;comment:填报日期;type:varchar(200);"`
      Cxbz  string `json:"cxbz" cn:"撤销标志" form:"cxbz" gorm:"column:cxbz;comment:撤销标志;type:varchar(200);"`

}


// TableName BbPiContacts 表名
func (BbPiContacts) TableName() string {
  return "bb_pi_contacts"
}

