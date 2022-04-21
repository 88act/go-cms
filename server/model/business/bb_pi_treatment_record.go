// 自动生成模板BbPiTreatmentRecord
package business

import (
	"go-cms/global"
)

// BbPiTreatmentRecord 结构体
// 如果含有time.Time 请自行import time包
type BbPiTreatmentRecord struct {
     BbPiTreatmentRecordMini
}

// BbPiTreatmentRecordMini 结构体
type BbPiTreatmentRecordMini struct {
      global.GO_MODEL
      Status  *int `json:"status" cn:"上传" form:"status" gorm:"column:status;comment:上传:0未上传 1已上传 2上传失败;type:smallint"`
      Jgdm  string `json:"jgdm" cn:"机构标识" form:"jgdm" gorm:"column:jgdm;comment:机构标识;type:varchar(200);"`
      Fwwddm  string `json:"fwwddm" cn:"服务网点代码" form:"fwwddm" gorm:"column:fwwddm;comment:服务网点代码;type:varchar(200);"`
      Jzlx  string `json:"jzlx" cn:"就诊类型" form:"jzlx" gorm:"column:jzlx;comment:就诊类型;type:varchar(200);"`
      Kh  string `json:"kh" cn:"卡号" form:"kh" gorm:"column:kh;comment:卡号;type:varchar(200);"`
      Klx  string `json:"klx" cn:"卡类型" form:"klx" gorm:"column:klx;comment:卡类型;type:varchar(200);"`
      Mzh  string `json:"mzh" cn:"门诊号" form:"mzh" gorm:"column:mzh;comment:门诊号;type:varchar(200);"`
      Ksbm  string `json:"ksbm" cn:"科室编码" form:"ksbm" gorm:"column:ksbm;comment:科室编码;type:varchar(200);"`
      Ksmc  string `json:"ksmc" cn:"科室名称" form:"ksmc" gorm:"column:ksmc;comment:科室名称;type:varchar(200);"`
      Xm  string `json:"xm" cn:"姓名" form:"xm" gorm:"column:xm;comment:姓名;type:varchar(200);"`
      Xbdm  string `json:"xbdm" cn:"性别代码" form:"xbdm" gorm:"column:xbdm;comment:性别代码;type:varchar(200);"`
      Csrq  string `json:"csrq" cn:"出生日期" form:"csrq" gorm:"column:csrq;comment:出生日期;type:varchar(200);"`
      Nls  string `json:"nls" cn:"年龄岁" form:"nls" gorm:"column:nls;comment:年龄岁;type:varchar(200);"`
      Nly  string `json:"nly" cn:"年龄月" form:"nly" gorm:"column:nly;comment:年龄月;type:varchar(200);"`
      Gmsbs  string `json:"gmsbs" cn:"过敏史标识" form:"gmsbs" gorm:"column:gmsbs;comment:过敏史标识;type:varchar(200);"`
      Gms  string `json:"gms" cn:"过敏史" form:"gms" gorm:"column:gms;comment:过敏史;type:varchar(200);"`
      Cblb  string `json:"cblb" cn:"参保类别" form:"cblb" gorm:"column:cblb;comment:参保类别;type:varchar(200);"`
      Jzrqsj  string `json:"jzrqsj" cn:"就诊日期时间" form:"jzrqsj" gorm:"column:jzrqsj;comment:就诊日期时间;type:varchar(200);"`
      Jzzdsm  string `json:"jzzdsm" cn:"就诊诊断说明" form:"jzzdsm" gorm:"column:jzzdsm;comment:就诊诊断说明;type:varchar(200);"`
      Czbzdm  string `json:"czbzdm" cn:"初诊标志代码" form:"czbzdm" gorm:"column:czbzdm;comment:初诊标志代码;type:varchar(200);"`
      Zs  string `json:"zs" cn:"主诉" form:"zs" gorm:"column:zs;comment:主诉;type:varchar(200);"`
      Xbs  string `json:"xbs" cn:"现病史" form:"xbs" gorm:"column:xbs;comment:现病史;type:varchar(200);"`
      Jws  string `json:"jws" cn:"既往史" form:"jws" gorm:"column:jws;comment:既往史;type:varchar(200);"`
      Fzjcxm  string `json:"fzjcxm" cn:"辅助检查项目" form:"fzjcxm" gorm:"column:fzjcxm;comment:辅助检查项目;type:varchar(200);"`
      Fzjcjg  string `json:"fzjcjg" cn:"辅助检查结果" form:"fzjcjg" gorm:"column:fzjcjg;comment:辅助检查结果;type:varchar(200);"`
      Mzzzzddm  string `json:"mzzzzddm" cn:"门诊症状-诊断代码" form:"mzzzzddm" gorm:"column:mzzzzddm;comment:门诊症状-诊断代码;type:varchar(200);"`
      Mzzzzdmc  string `json:"mzzzzdmc" cn:"门诊症状-诊断名称" form:"mzzzzdmc" gorm:"column:mzzzzdmc;comment:门诊症状-诊断名称;type:varchar(400);"`
      Mzzzzdbmlx  string `json:"mzzzzdbmlx" cn:"门诊症状诊断编码类型" form:"mzzzzdbmlx" gorm:"column:mzzzzdbmlx;comment:门诊症状诊断编码类型;type:varchar(200);"`
      Zzms  string `json:"zzms" cn:"症状描述" form:"zzms" gorm:"column:zzms;comment:症状描述;type:varchar(200);"`
      Bzyj  string `json:"bzyj" cn:"辨证依据" form:"bzyj" gorm:"column:bzyj;comment:辨证依据;type:varchar(200);"`
      Zzzf  string `json:"zzzf" cn:"治则治法" form:"zzzf" gorm:"column:zzzf;comment:治则治法;type:varchar(200);"`
      Jkzd  string `json:"jkzd" cn:"健康指导" form:"jkzd" gorm:"column:jkzd;comment:健康指导;type:varchar(200);"`
      Czjh  string `json:"czjh" cn:"处置计划" form:"czjh" gorm:"column:czjh;comment:处置计划;type:varchar(200);"`
      Yzysgh  string `json:"yzysgh" cn:"应诊医生工号" form:"yzysgh" gorm:"column:yzysgh;comment:应诊医生工号;type:varchar(200);"`
      Yzysqm  string `json:"yzysqm" cn:"应诊医师签名" form:"yzysqm" gorm:"column:yzysqm;comment:应诊医师签名;type:varchar(200);"`
      Czylwsjgdm  string `json:"czylwsjgdm" cn:"初诊医疗卫生机构代码" form:"czylwsjgdm" gorm:"column:czylwsjgdm;comment:初诊医疗卫生机构代码;type:varchar(200);"`
      Czylwsjgmc  string `json:"czylwsjgmc" cn:"初诊机构名称" form:"czylwsjgmc" gorm:"column:czylwsjgmc;comment:初诊机构名称;type:varchar(200);"`
      Jzjssj  string `json:"jzjssj" cn:"就诊结束时间" form:"jzjssj" gorm:"column:jzjssj;comment:就诊结束时间;type:varchar(200);"`
      Zzbz  string `json:"zzbz" cn:"转诊标志" form:"zzbz" gorm:"column:zzbz;comment:转诊标志;type:varchar(200);"`
      Hzmyd  string `json:"hzmyd" cn:"患者满意度" form:"hzmyd" gorm:"column:hzmyd;comment:患者满意度;type:varchar(200);"`
      Yl1  string `json:"yl1" cn:"预留一" form:"yl1" gorm:"column:yl1;comment:预留一;type:varchar(200);"`
      Yl2  string `json:"yl2" cn:"预留二" form:"yl2" gorm:"column:yl2;comment:预留二;type:varchar(200);"`
      Dzcfwjcfdz  string `json:"dzcfwjcfdz" cn:"电子处方地址" form:"dzcfwjcfdz" gorm:"column:dzcfwjcfdz;comment:电子处方地址;type:varchar(200);"`
      Ysdspwjcfdz  string `json:"ysdspwjcfdz" cn:"医生端视频地址" form:"ysdspwjcfdz" gorm:"column:ysdspwjcfdz;comment:医生端视频地址;type:varchar(200);"`
      Hzdspwjcfdz  string `json:"hzdspwjcfdz" cn:"患者端视频地址" form:"hzdspwjcfdz" gorm:"column:hzdspwjcfdz;comment:患者端视频地址;type:varchar(200);"`
      Jlypcfdz  string `json:"jlypcfdz" cn:"交流音频地址" form:"jlypcfdz" gorm:"column:jlypcfdz;comment:交流音频地址;type:varchar(200);"`
      Sjscsj  string `json:"sjscsj" cn:"数据生成日期时间" form:"sjscsj" gorm:"column:sjscsj;comment:数据生成日期时间;type:varchar(200);"`
      Tbrqsj  string `json:"tbrqsj" cn:"填报日期" form:"tbrqsj" gorm:"column:tbrqsj;comment:填报日期;type:varchar(200);"`
      Mj  string `json:"mj" cn:"密级" form:"mj" gorm:"column:mj;comment:密级;type:varchar(200);"`
      Cxbz  string `json:"cxbz" cn:"撤销标志" form:"cxbz" gorm:"column:cxbz;comment:撤销标志;type:varchar(200);"`

}


// TableName BbPiTreatmentRecord 表名
func (BbPiTreatmentRecord) TableName() string {
  return "bb_pi_treatment_record"
}

