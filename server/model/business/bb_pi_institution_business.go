// 自动生成模板BbPiInstitutionBusiness
package business

import (
	"go-cms/global"
)

// BbPiInstitutionBusiness 结构体
// 如果含有time.Time 请自行import time包
type BbPiInstitutionBusiness struct {
     BbPiInstitutionBusinessMini
}

// BbPiInstitutionBusinessMini 结构体
type BbPiInstitutionBusinessMini struct {
      global.GO_MODEL
      Status  *int `json:"status" cn:"上传" form:"status" gorm:"column:status;comment:上传:0未上传 1已上传 2上传失败;type:smallint"`
      Jgdm  string `json:"jgdm" cn:"机构标识" form:"jgdm" gorm:"column:jgdm;comment:机构标识;type:varchar(200);"`
      Nf  string `json:"nf" cn:"月份" form:"nf" gorm:"column:nf;comment:月份;type:varchar(200);"`
      Pcjgsl  string `json:"pcjgsl" cn:"派出机构数量" form:"pcjgsl" gorm:"column:pcjgsl;comment:派出机构数量;type:varchar(200);"`
      Zgzs  string `json:"zgzs" cn:"职工总数" form:"zgzs" gorm:"column:zgzs;comment:职工总数;type:varchar(200);"`
      Khffryzs  string `json:"khffryzs" cn:"客户服务人员总数" form:"khffryzs" gorm:"column:khffryzs;comment:客户服务人员总数;type:varchar(200);"`
      Rjzzyspbs  string `json:"rjzzyspbs" cn:"日均坐诊医生数" form:"rjzzyspbs" gorm:"column:rjzzyspbs;comment:日均坐诊医生数;type:varchar(200);"`
      Ywyfmj  string `json:"ywyfmj" cn:"业务用房面积" form:"ywyfmj" gorm:"column:ywyfmj;comment:业务用房面积;type:varchar(200);"`
      Zsr  string `json:"zsr" cn:"总收入" form:"zsr" gorm:"column:zsr;comment:总收入;type:varchar(200);"`
      Zzc  string `json:"zzc" cn:"总支出" form:"zzc" gorm:"column:zzc;comment:总支出;type:varchar(200);"`
      Zzch  string `json:"zzch" cn:"总资产" form:"zzch" gorm:"column:zzch;comment:总资产;type:varchar(200);"`
      Ldzc  string `json:"ldzc" cn:"流动资产" form:"ldzc" gorm:"column:ldzc;comment:流动资产;type:varchar(200);"`
      Dwtz  string `json:"dwtz" cn:"对外投资" form:"dwtz" gorm:"column:dwtz;comment:对外投资;type:varchar(200);"`
      Gdzc  string `json:"gdzc" cn:"固定资产" form:"gdzc" gorm:"column:gdzc;comment:固定资产;type:varchar(200);"`
      Wxzcjkbf  string `json:"wxzcjkbf" cn:"无形资产及开办费" form:"wxzcjkbf" gorm:"column:wxzcjkbf;comment:无形资产及开办费;type:varchar(200);"`
      Fz  string `json:"fz" cn:"负债" form:"fz" gorm:"column:fz;comment:负债;type:varchar(200);"`
      Jzc  string `json:"jzc" cn:"净资产" form:"jzc" gorm:"column:jzc;comment:净资产;type:varchar(200);"`
      Sjscsj  string `json:"sjscsj" cn:"数据生成日期时间" form:"sjscsj" gorm:"column:sjscsj;comment:数据生成日期时间;type:varchar(200);"`
      Tbrqsj  string `json:"tbrqsj" cn:"填报日期" form:"tbrqsj" gorm:"column:tbrqsj;comment:填报日期;type:varchar(200);"`
      Cxbz  string `json:"cxbz" cn:"撤销标志" form:"cxbz" gorm:"column:cxbz;comment:撤销标志;type:varchar(200);"`

}


// TableName BbPiInstitutionBusiness 表名
func (BbPiInstitutionBusiness) TableName() string {
  return "bb_pi_institution_business"
}

