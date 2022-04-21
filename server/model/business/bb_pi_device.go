// 自动生成模板BbPiDevice
package business

import (
	"go-cms/global"
)

// BbPiDevice 结构体
// 如果含有time.Time 请自行import time包
type BbPiDevice struct {
     BbPiDeviceMini
}

// BbPiDeviceMini 结构体
type BbPiDeviceMini struct {
      global.GO_MODEL
      Status  *int `json:"status" cn:"上传" form:"status" gorm:"column:status;comment:上传:0未上传 1已上传 2上传失败;type:smallint"`
      Jgdm  string `json:"jgdm" cn:"机构标识" form:"jgdm" gorm:"column:jgdm;comment:机构标识;type:varchar(200);"`
      Sbdh  string `json:"sbdh" cn:"设备代号" form:"sbdh" gorm:"column:sbdh;comment:设备代号;type:varchar(200);"`
      Sbmc  string `json:"sbmc" cn:"设备名称" form:"sbmc" gorm:"column:sbmc;comment:设备名称;type:varchar(200);"`
      Tpsbts  string `json:"tpsbts" cn:"同批设备台数" form:"tpsbts" gorm:"column:tpsbts;comment:同批设备台数;type:varchar(200);"`
      Cd  string `json:"cd" cn:"产地" form:"cd" gorm:"column:cd;comment:产地;type:varchar(200);"`
      Sccj  string `json:"sccj" cn:"生产厂家" form:"sccj" gorm:"column:sccj;comment:生产厂家;type:varchar(200);"`
      Sbxh  string `json:"sbxh" cn:"设备型号" form:"sbxh" gorm:"column:sbxh;comment:设备型号;type:varchar(200);"`
      Sbcs  string `json:"sbcs" cn:"设备参数" form:"sbcs" gorm:"column:sbcs;comment:设备参数;type:varchar(200);"`
      Sblx  string `json:"sblx" cn:"设备类型" form:"sblx" gorm:"column:sblx;comment:设备类型;type:varchar(200);"`
      Gmrq  string `json:"gmrq" cn:"购买日期" form:"gmrq" gorm:"column:gmrq;comment:购买日期;type:varchar(200);"`
      Yt  string `json:"yt" cn:"用途" form:"yt" gorm:"column:yt;comment:用途;type:varchar(200);"`
      Sbjzje  string `json:"sbjzje" cn:"设备价值金额" form:"sbjzje" gorm:"column:sbjzje;comment:设备价值金额;type:varchar(200);"`
      Gjsxqk  string `json:"gjsxqk" cn:"购进时新旧情况" form:"gjsxqk" gorm:"column:gjsxqk;comment:购进时新旧情况;type:varchar(200);"`
      Llsjsm  string `json:"llsjsm" cn:"理论设计寿命" form:"llsjsm" gorm:"column:llsjsm;comment:理论设计寿命;type:varchar(200);"`
      Syqk  string `json:"syqk" cn:"使用情况" form:"syqk" gorm:"column:syqk;comment:使用情况;type:varchar(200);"`
      Bz  string `json:"bz" cn:"备注" form:"bz" gorm:"column:bz;comment:备注;type:varchar(200);"`
      Sjscsj  string `json:"sjscsj" cn:"数据生成时间" form:"sjscsj" gorm:"column:sjscsj;comment:数据生成时间;type:varchar(200);"`
      Tbrqsj  string `json:"tbrqsj" cn:"填报日期" form:"tbrqsj" gorm:"column:tbrqsj;comment:填报日期;type:varchar(200);"`
      Cxbz  string `json:"cxbz" cn:"撤销标志" form:"cxbz" gorm:"column:cxbz;comment:撤销标志;type:varchar(200);"`

}


// TableName BbPiDevice 表名
func (BbPiDevice) TableName() string {
  return "bb_pi_device"
}

