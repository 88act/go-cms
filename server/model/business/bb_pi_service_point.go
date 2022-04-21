// 自动生成模板BbPiServicePoint
package business

import (
	"go-cms/global"
)

// BbPiServicePoint 结构体
// 如果含有time.Time 请自行import time包
type BbPiServicePoint struct {
     BbPiServicePointMini
}

// BbPiServicePointMini 结构体
type BbPiServicePointMini struct {
      global.GO_MODEL
      Status  *int `json:"status" cn:"上传" form:"status" gorm:"column:status;comment:上传:0未上传 1已上传 2上传失败;type:smallint"`
      Jgdm  string `json:"jgdm" cn:"机构标识" form:"jgdm" gorm:"column:jgdm;comment:机构标识;type:varchar(200);"`
      Zzjgdm  string `json:"zzjgdm" cn:"统一社会信用代码" form:"zzjgdm" gorm:"column:zzjgdm;comment:统一社会信用代码;type:varchar(200);"`
      Fwwddm  string `json:"fwwddm" cn:"服务网点代码" form:"fwwddm" gorm:"column:fwwddm;comment:服务网点代码;type:varchar(200);"`
      Fwdmc  string `json:"fwdmc" cn:"服务点名称" form:"fwdmc" gorm:"column:fwdmc;comment:服务点名称;type:varchar(200);"`
      Xzqhdm  string `json:"xzqhdm" cn:"行政区划代码" form:"xzqhdm" gorm:"column:xzqhdm;comment:行政区划代码;type:varchar(200);"`
      Fwdlx  string `json:"fwdlx" cn:"服务点类型" form:"fwdlx" gorm:"column:fwdlx;comment:服务点类型;type:varchar(200);"`
      Fwdclrq  string `json:"fwdclrq" cn:"服务点成立日期" form:"fwdclrq" gorm:"column:fwdclrq;comment:服务点成立日期;type:varchar(200);"`
      Dwlsgxdm  string `json:"dwlsgxdm" cn:"单位隶属关系代码" form:"dwlsgxdm" gorm:"column:dwlsgxdm;comment:单位隶属关系代码;type:varchar(200);"`
      Fwdflgllbdm  string `json:"fwdflgllbdm" cn:"服务点分类管理类别代码" form:"fwdflgllbdm" gorm:"column:fwdflgllbdm;comment:服务点分类管理类别代码;type:varchar(200);"`
      Fwdfldm  string `json:"fwdfldm" cn:"服务点分类代码" form:"fwdfldm" gorm:"column:fwdfldm;comment:服务点分类代码;type:varchar(200);"`
      Jjlxdm  string `json:"jjlxdm" cn:"经济类型代码" form:"jjlxdm" gorm:"column:jjlxdm;comment:经济类型代码;type:varchar(200);"`
      Dz  string `json:"dz" cn:"地址" form:"dz" gorm:"column:dz;comment:地址;type:varchar(200);"`
      Fwdyyjb  string `json:"fwdyyjb" cn:"服务点医院级别" form:"fwdyyjb" gorm:"column:fwdyyjb;comment:服务点医院级别;type:varchar(200);"`
      Fwdyydj  string `json:"fwdyydj" cn:"服务点医院等级" form:"fwdyydj" gorm:"column:fwdyydj;comment:服务点医院等级;type:varchar(200);"`
      Xkzhm  string `json:"xkzhm" cn:"许可证号码" form:"xkzhm" gorm:"column:xkzhm;comment:许可证号码;type:varchar(200);"`
      Xkxmmc  string `json:"xkxmmc" cn:"许可项目名称" form:"xkxmmc" gorm:"column:xkxmmc;comment:许可项目名称;type:varchar(200);"`
      Xkzyxq  string `json:"xkzyxq" cn:"许可证有效期" form:"xkzyxq" gorm:"column:xkzyxq;comment:许可证有效期;type:varchar(200);"`
      Kbzjjes  string `json:"kbzjjes" cn:"开办资金金额数" form:"kbzjjes" gorm:"column:kbzjjes;comment:开办资金金额数;type:varchar(200);"`
      Frdb  string `json:"frdb" cn:"法人代表" form:"frdb" gorm:"column:frdb;comment:法人代表;type:varchar(200);"`
      Fwdszdmzzzdfbz  string `json:"fwdszdmzzzdfbz" cn:"服务点地方标志" form:"fwdszdmzzzdfbz" gorm:"column:fwdszdmzzzdfbz;comment:服务点地方标志;type:varchar(200);"`
      Sffzjg  string `json:"sffzjg" cn:"是否分支机构" form:"sffzjg" gorm:"column:sffzjg;comment:是否分支机构;type:varchar(200);"`

}


// TableName BbPiServicePoint 表名
func (BbPiServicePoint) TableName() string {
  return "bb_pi_service_point"
}

