// 自动生成模板BbPiStaff
package business

import (
	"go-cms/global"
)

// BbPiStaff 结构体
// 如果含有time.Time 请自行import time包
type BbPiStaff struct {
     BbPiStaffMini
}

// BbPiStaffMini 结构体
type BbPiStaffMini struct {
      global.GO_MODEL
      Status  *int `json:"status" cn:"上传" form:"status" gorm:"column:status;comment:上传:0未上传 1已上传 2上传失败;type:smallint"`
      Jgdm  string `json:"jgdm" cn:"机构标识" form:"jgdm" gorm:"column:jgdm;comment:机构标识;type:varchar(200);"`
      Yhrygh  string `json:"yhrygh" cn:"医护人员账号" form:"yhrygh" gorm:"column:yhrygh;comment:医护人员账号;type:varchar(200);"`
      Yhryxm  string `json:"yhryxm" cn:"医护人员姓名" form:"yhryxm" gorm:"column:yhryxm;comment:医护人员姓名;type:varchar(200);"`
      Xb  string `json:"xb" cn:"性别" form:"xb" gorm:"column:xb;comment:性别;type:varchar(200);"`
      Csrq  string `json:"csrq" cn:"出生日期" form:"csrq" gorm:"column:csrq;comment:出生日期;type:varchar(200);"`
      Sfzh  string `json:"sfzh" cn:"身份证号" form:"sfzh" gorm:"column:sfzh;comment:身份证号;type:varchar(200);"`
      Zjlbdm  string `json:"zjlbdm" cn:"身份证件类别代码" form:"zjlbdm" gorm:"column:zjlbdm;comment:身份证件类别代码;type:varchar(200);"`
      Ksdm  string `json:"ksdm" cn:"所属科室代码" form:"ksdm" gorm:"column:ksdm;comment:所属科室代码;type:varchar(200);"`
      Zyjszwdm  string `json:"zyjszwdm" cn:"专业技术职务代码" form:"zyjszwdm" gorm:"column:zyjszwdm;comment:专业技术职务代码;type:varchar(200);"`
      Zyjszwlbdm  string `json:"zyjszwlbdm" cn:"专业技术职务类别代码" form:"zyjszwlbdm" gorm:"column:zyjszwlbdm;comment:专业技术职务类别代码;type:varchar(200);"`
      Zzlbmc  string `json:"zzlbmc" cn:"资质类别名称" form:"zzlbmc" gorm:"column:zzlbmc;comment:资质类别名称;type:varchar(200);"`
      Zgzsbh  string `json:"zgzsbh" cn:"资格证书编号" form:"zgzsbh" gorm:"column:zgzsbh;comment:资格证书编号;type:varchar(200);"`
      Zghdsj  string `json:"zghdsj" cn:"资格获得时间" form:"zghdsj" gorm:"column:zghdsj;comment:资格获得时间;type:varchar(200);"`
      Zyzsbm  string `json:"zyzsbm" cn:"执业证书编码" form:"zyzsbm" gorm:"column:zyzsbm;comment:执业证书编码;type:varchar(200);"`
      Fzrq  string `json:"fzrq" cn:"发证日期" form:"fzrq" gorm:"column:fzrq;comment:发证日期;type:varchar(200);"`
      Zydd  string `json:"zydd" cn:"执业地点" form:"zydd" gorm:"column:zydd;comment:执业地点;type:varchar(200);"`
      Zyfw  string `json:"zyfw" cn:"执业范围" form:"zyfw" gorm:"column:zyfw;comment:执业范围;type:varchar(200);"`
      Zyzyyljgdm  string `json:"zyzyyljgdm" cn:"主要执业医疗机构代码" form:"zyzyyljgdm" gorm:"column:zyzyyljgdm;comment:主要执业医疗机构代码;type:varchar(200);"`
      Zyzyyymc  string `json:"zyzyyymc" cn:"主要执业医院名称" form:"zyzyyymc" gorm:"column:zyzyyymc;comment:主要执业医院名称;type:varchar(200);"`
      Qkysbz  string `json:"qkysbz" cn:"全科医生标志" form:"qkysbz" gorm:"column:qkysbz;comment:全科医生标志;type:varchar(200);"`
      Sjhm  string `json:"sjhm" cn:"手机号码" form:"sjhm" gorm:"column:sjhm;comment:手机号码;type:varchar(200);"`
      Zc  string `json:"zc" cn:"专长" form:"zc" gorm:"column:zc;comment:专长;type:varchar(200);"`
      Mz  string `json:"mz" cn:"民族" form:"mz" gorm:"column:mz;comment:民族;type:varchar(200);"`
      Cjgzrq  string `json:"cjgzrq" cn:"参加工作日期" form:"cjgzrq" gorm:"column:cjgzrq;comment:参加工作日期;type:varchar(200);"`
      Zcsj  string `json:"zcsj" cn:"注册日期时间" form:"zcsj" gorm:"column:zcsj;comment:注册日期时间;type:varchar(200);"`
      Xl  string `json:"xl" cn:"学历" form:"xl" gorm:"column:xl;comment:学历;type:varchar(200);"`
      Grzpcfdz  string `json:"grzpcfdz" cn:"个人照片存放地址" form:"grzpcfdz" gorm:"column:grzpcfdz;comment:个人照片存放地址;type:varchar(200);"`
      Zgzcfdz  string `json:"zgzcfdz" cn:"资格证存放地址" form:"zgzcfdz" gorm:"column:zgzcfdz;comment:资格证存放地址;type:varchar(200);"`
      Zyzcfdz  string `json:"zyzcfdz" cn:"执业证存放地址" form:"zyzcfdz" gorm:"column:zyzcfdz;comment:执业证存放地址;type:varchar(200);"`
      Sjscsj  string `json:"sjscsj" cn:"数据生成日期时间" form:"sjscsj" gorm:"column:sjscsj;comment:数据生成日期时间;type:varchar(200);"`
      Tbrqsj  string `json:"tbrqsj" cn:"填报日期" form:"tbrqsj" gorm:"column:tbrqsj;comment:填报日期;type:varchar(200);"`
      Cxbz  string `json:"cxbz" cn:"撤销标志" form:"cxbz" gorm:"column:cxbz;comment:撤销标志;type:varchar(200);"`

}


// TableName BbPiStaff 表名
func (BbPiStaff) TableName() string {
  return "bb_pi_staff"
}

