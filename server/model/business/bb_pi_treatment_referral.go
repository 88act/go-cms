// 自动生成模板BbPiTreatmentReferral
package business

import (
	"go-cms/global"
)

// BbPiTreatmentReferral 结构体
// 如果含有time.Time 请自行import time包
type BbPiTreatmentReferral struct {
     BbPiTreatmentReferralMini
}

// BbPiTreatmentReferralMini 结构体
type BbPiTreatmentReferralMini struct {
      global.GO_MODEL
      Status  *int `json:"status" cn:"上传" form:"status" gorm:"column:status;comment:上传:0未上传 1已上传 2上传失败;type:smallint"`
      Jgdm  string `json:"jgdm" cn:"机构标识" form:"jgdm" gorm:"column:jgdm;comment:机构标识;type:varchar(200);"`
      Fwwddm  string `json:"fwwddm" cn:"服务网点代码" form:"fwwddm" gorm:"column:fwwddm;comment:服务网点代码;type:varchar(200);"`
      Zzjlid  string `json:"zzjlid" cn:"转诊记录" form:"zzjlid" gorm:"column:zzjlid;comment:转诊记录;type:varchar(200);"`
      Kh  string `json:"kh" cn:"卡号" form:"kh" gorm:"column:kh;comment:卡号;type:varchar(200);"`
      Klx  string `json:"klx" cn:"卡类型" form:"klx" gorm:"column:klx;comment:卡类型;type:varchar(200);"`
      Mzh  string `json:"mzh" cn:"门诊号" form:"mzh" gorm:"column:mzh;comment:门诊号;type:varchar(200);"`
      Xm  string `json:"xm" cn:"姓名" form:"xm" gorm:"column:xm;comment:姓名;type:varchar(200);"`
      Xb  string `json:"xb" cn:"性别" form:"xb" gorm:"column:xb;comment:性别;type:varchar(200);"`
      Csrq  string `json:"csrq" cn:"出生日期" form:"csrq" gorm:"column:csrq;comment:出生日期;type:varchar(200);"`
      Nls  string `json:"nls" cn:"年龄岁" form:"nls" gorm:"column:nls;comment:年龄岁;type:varchar(200);"`
      Nly  string `json:"nly" cn:"年龄月" form:"nly" gorm:"column:nly;comment:年龄月;type:varchar(200);"`
      Jzrqsj  string `json:"jzrqsj" cn:"就诊时间" form:"jzrqsj" gorm:"column:jzrqsj;comment:就诊时间;type:varchar(200);"`
      Fzysgh  string `json:"fzysgh" cn:"负责医生工号" form:"fzysgh" gorm:"column:fzysgh;comment:负责医生工号;type:varchar(200);"`
      Fzysxm  string `json:"fzysxm" cn:"负责医生姓名" form:"fzysxm" gorm:"column:fzysxm;comment:负责医生姓名;type:varchar(200);"`
      Fzksbm  string `json:"fzksbm" cn:"负责科室编码" form:"fzksbm" gorm:"column:fzksbm;comment:负责科室编码;type:varchar(200);"`
      Fzksmc  string `json:"fzksmc" cn:"负责科室名称" form:"fzksmc" gorm:"column:fzksmc;comment:负责科室名称;type:varchar(200);"`
      Zzyy  string `json:"zzyy" cn:"转诊原因" form:"zzyy" gorm:"column:zzyy;comment:转诊原因;type:varchar(200);"`
      Zzsj  string `json:"zzsj" cn:"转诊时间" form:"zzsj" gorm:"column:zzsj;comment:转诊时间;type:varchar(200);"`
      Zxjgdm  string `json:"zxjgdm" cn:"转向机构代码" form:"zxjgdm" gorm:"column:zxjgdm;comment:转向机构代码;type:varchar(200);"`
      Zxjgmc  string `json:"zxjgmc" cn:"转向机构名称" form:"zxjgmc" gorm:"column:zxjgmc;comment:转向机构名称;type:varchar(200);"`
      Zxksdm  string `json:"zxksdm" cn:"转向科室代码" form:"zxksdm" gorm:"column:zxksdm;comment:转向科室代码;type:varchar(200);"`
      Zxksmc  string `json:"zxksmc" cn:"转向科室名称" form:"zxksmc" gorm:"column:zxksmc;comment:转向科室名称;type:varchar(200);"`
      Zxysgh  string `json:"zxysgh" cn:"转向医生工号" form:"zxysgh" gorm:"column:zxysgh;comment:转向医生工号;type:varchar(200);"`
      Zxysxm  string `json:"zxysxm" cn:"转向医生姓名" form:"zxysxm" gorm:"column:zxysxm;comment:转向医生姓名;type:varchar(200);"`
      Zyjws  string `json:"zyjws" cn:"主要既往史" form:"zyjws" gorm:"column:zyjws;comment:主要既往史;type:varchar(200);"`
      Zljg  string `json:"zljg" cn:"治疗经过" form:"zljg" gorm:"column:zljg;comment:治疗经过;type:varchar(200);"`
      Xybzlfa  string `json:"xybzlfa" cn:"下一步治疗方案" form:"xybzlfa" gorm:"column:xybzlfa;comment:下一步治疗方案;type:varchar(200);"`
      Zzbz  string `json:"zzbz" cn:"转诊标志" form:"zzbz" gorm:"column:zzbz;comment:转诊标志;type:varchar(200);"`
      Sjscsj  string `json:"sjscsj" cn:"数据生成时间" form:"sjscsj" gorm:"column:sjscsj;comment:数据生成时间;type:varchar(200);"`
      Tbrqsj  string `json:"tbrqsj" cn:"填报日期" form:"tbrqsj" gorm:"column:tbrqsj;comment:填报日期;type:varchar(200);"`
      Mj  string `json:"mj" cn:"密级" form:"mj" gorm:"column:mj;comment:密级;type:varchar(200);"`
      Cxbz  string `json:"cxbz" cn:"撤销标志" form:"cxbz" gorm:"column:cxbz;comment:撤销标志;type:varchar(200);"`

}


// TableName BbPiTreatmentReferral 表名
func (BbPiTreatmentReferral) TableName() string {
  return "bb_pi_treatment_referral"
}

