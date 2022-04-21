// 自动生成模板BbPiTreatmentDiagnose
package business

import (
	"go-cms/global"
)

// BbPiTreatmentDiagnose 结构体
// 如果含有time.Time 请自行import time包
type BbPiTreatmentDiagnose struct {
     BbPiTreatmentDiagnoseMini
}

// BbPiTreatmentDiagnoseMini 结构体
type BbPiTreatmentDiagnoseMini struct {
      global.GO_MODEL
      Status  *int `json:"status" cn:"上传" form:"status" gorm:"column:status;comment:上传:0未上传 1已上传 2上传失败;type:smallint"`
      Jgdm  string `json:"jgdm" cn:"机构标识" form:"jgdm" gorm:"column:jgdm;comment:机构标识;type:varchar(200);"`
      Fwwddm  string `json:"fwwddm" cn:"服务网点代码" form:"fwwddm" gorm:"column:fwwddm;comment:服务网点代码;type:varchar(200);"`
      Cfbh  string `json:"cfbh" cn:"处方编号" form:"cfbh" gorm:"column:cfbh;comment:处方编号;type:varchar(200);"`
      Cfmxid  string `json:"cfmxid" cn:"处方明细ID" form:"cfmxid" gorm:"column:cfmxid;comment:处方明细ID;type:varchar(200);"`
      Zdxxid  string `json:"zdxxid" cn:"诊断信息 ID" form:"zdxxid" gorm:"column:zdxxid;comment:诊断信息 ID;type:varchar(200);"`
      Kh  string `json:"kh" cn:"卡号" form:"kh" gorm:"column:kh;comment:卡号;type:varchar(200);"`
      Klx  string `json:"klx" cn:"卡类型" form:"klx" gorm:"column:klx;comment:卡类型;type:varchar(200);"`
      Mzh  string `json:"mzh" cn:"门诊号" form:"mzh" gorm:"column:mzh;comment:门诊号;type:varchar(200);"`
      Xm  string `json:"xm" cn:"姓名" form:"xm" gorm:"column:xm;comment:姓名;type:varchar(200);"`
      Xbdm  string `json:"xbdm" cn:"性别" form:"xbdm" gorm:"column:xbdm;comment:性别;type:varchar(200);"`
      Csrq  string `json:"csrq" cn:"出生日期" form:"csrq" gorm:"column:csrq;comment:出生日期;type:varchar(200);"`
      Nls  string `json:"nls" cn:"年龄岁" form:"nls" gorm:"column:nls;comment:年龄岁;type:varchar(200);"`
      Nly  string `json:"nly" cn:"年龄月" form:"nly" gorm:"column:nly;comment:年龄月;type:varchar(200);"`
      Jzrqsj  string `json:"jzrqsj" cn:"就诊时间" form:"jzrqsj" gorm:"column:jzrqsj;comment:就诊时间;type:varchar(200);"`
      Zdlxbm  string `json:"zdlxbm" cn:"诊断类型编码" form:"zdlxbm" gorm:"column:zdlxbm;comment:诊断类型编码;type:varchar(200);"`
      Xyzdbm  string `json:"xyzdbm" cn:"西医诊断编码" form:"xyzdbm" gorm:"column:xyzdbm;comment:西医诊断编码;type:varchar(200);"`
      Xyzdmc  string `json:"xyzdmc" cn:"西医诊断名称" form:"xyzdmc" gorm:"column:xyzdmc;comment:西医诊断名称;type:varchar(200);"`
      Xyzdbmlx  string `json:"xyzdbmlx" cn:"西医诊断编码类型" form:"xyzdbmlx" gorm:"column:xyzdbmlx;comment:西医诊断编码类型;type:varchar(200);"`
      Zyzdbm  string `json:"zyzdbm" cn:"中医诊断编码" form:"zyzdbm" gorm:"column:zyzdbm;comment:中医诊断编码;type:varchar(200);"`
      Zyzdmc  string `json:"zyzdmc" cn:"中医诊断名称" form:"zyzdmc" gorm:"column:zyzdmc;comment:中医诊断名称;type:varchar(200);"`
      Zyzdbmlx  string `json:"zyzdbmlx" cn:"中医诊断编码类型" form:"zyzdbmlx" gorm:"column:zyzdbmlx;comment:中医诊断编码类型;type:varchar(200);"`
      Zdsm  string `json:"zdsm" cn:"诊断说明" form:"zdsm" gorm:"column:zdsm;comment:诊断说明;type:varchar(200);"`
      Zdbz  string `json:"zdbz" cn:"诊断标志" form:"zdbz" gorm:"column:zdbz;comment:诊断标志;type:varchar(200);"`
      Zdysgh  string `json:"zdysgh" cn:"诊断医生工号" form:"zdysgh" gorm:"column:zdysgh;comment:诊断医生工号;type:varchar(200);"`
      Zdysxm  string `json:"zdysxm" cn:"诊断医生姓名" form:"zdysxm" gorm:"column:zdysxm;comment:诊断医生姓名;type:varchar(200);"`
      Zdsj  string `json:"zdsj" cn:"诊断时间" form:"zdsj" gorm:"column:zdsj;comment:诊断时间;type:varchar(200);"`
      Yl1  string `json:"yl1" cn:"预留一" form:"yl1" gorm:"column:yl1;comment:预留一;type:varchar(200);"`
      Sjscsj  string `json:"sjscsj" cn:"数据生成时间" form:"sjscsj" gorm:"column:sjscsj;comment:数据生成时间;type:varchar(200);"`
      Tbrqsj  string `json:"tbrqsj" cn:"填报日期" form:"tbrqsj" gorm:"column:tbrqsj;comment:填报日期;type:varchar(200);"`
      Cxbz  string `json:"cxbz" cn:"撤销标志" form:"cxbz" gorm:"column:cxbz;comment:撤销标志;type:varchar(200);"`
      Mj  string `json:"mj" cn:"密级" form:"mj" gorm:"column:mj;comment:密级;type:varchar(200);"`

}


// TableName BbPiTreatmentDiagnose 表名
func (BbPiTreatmentDiagnose) TableName() string {
  return "bb_pi_treatment_diagnose"
}

