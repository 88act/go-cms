// 自动生成模板BbPiPerson
package business

import (
	"go-cms/global"
)

// BbPiPerson 结构体
// 如果含有time.Time 请自行import time包
type BbPiPerson struct {
     BbPiPersonMini
}

// BbPiPersonMini 结构体
type BbPiPersonMini struct {
      global.GO_MODEL
      Userid  *int `json:"userid" cn:"userid" form:"userid" gorm:"column:userid;comment:userid;type:int"`
      Status  *int `json:"status" cn:"上传" form:"status" gorm:"column:status;comment:上传:0未上传 1已上传 2上传失败;type:smallint"`
      Jgdm  string `json:"jgdm" cn:"机构标识" form:"jgdm" gorm:"column:jgdm;comment:机构标识;type:varchar(200);"`
      Kh  string `json:"kh" cn:"卡号" form:"kh" gorm:"column:kh;comment:卡号;type:varchar(200);"`
      Klx  string `json:"klx" cn:"卡类型" form:"klx" gorm:"column:klx;comment:卡类型;type:varchar(200);"`
      Zjhm  string `json:"zjhm" cn:"身份证件号码" form:"zjhm" gorm:"column:zjhm;comment:身份证件号码;type:varchar(200);"`
      Zjlbdm  string `json:"zjlbdm" cn:"身份证件类别" form:"zjlbdm" gorm:"column:zjlbdm;comment:身份证件类别;type:varchar(200);"`
      Xm  string `json:"xm" cn:"姓名" form:"xm" gorm:"column:xm;comment:姓名;type:varchar(200);"`
      Xbdm  string `json:"xbdm" cn:"性别代码" form:"xbdm" gorm:"column:xbdm;comment:性别代码;type:varchar(200);"`
      Xbmc  string `json:"xbmc" cn:"性别名称" form:"xbmc" gorm:"column:xbmc;comment:性别名称;type:varchar(200);"`
      Hzsx  string `json:"hzsx" cn:"患者属性" form:"hzsx" gorm:"column:hzsx;comment:患者属性;type:varchar(200);"`
      Hyztdm  string `json:"hyztdm" cn:"婚姻状况代码" form:"hyztdm" gorm:"column:hyztdm;comment:婚姻状况代码;type:varchar(200);"`
      Hyztmc  string `json:"hyztmc" cn:"婚姻状态名称" form:"hyztmc" gorm:"column:hyztmc;comment:婚姻状态名称;type:varchar(200);"`
      Csrq  string `json:"csrq" cn:"出生日期" form:"csrq" gorm:"column:csrq;comment:出生日期;type:varchar(200);"`
      Mzdm  string `json:"mzdm" cn:"民族代码" form:"mzdm" gorm:"column:mzdm;comment:民族代码;type:varchar(200);"`
      Mzmc  string `json:"mzmc" cn:"民族名称" form:"mzmc" gorm:"column:mzmc;comment:民族名称;type:varchar(200);"`
      Gjdm  string `json:"gjdm" cn:"国籍代码" form:"gjdm" gorm:"column:gjdm;comment:国籍代码;type:varchar(200);"`
      Gjmc  string `json:"gjmc" cn:"国籍名称" form:"gjmc" gorm:"column:gjmc;comment:国籍名称;type:varchar(200);"`
      Jgnbdah  string `json:"jgnbdah" cn:"机构内部档案号" form:"jgnbdah" gorm:"column:jgnbdah;comment:机构内部档案号;type:varchar(200);"`
      Gddhhm  string `json:"gddhhm" cn:"固定电话" form:"gddhhm" gorm:"column:gddhhm;comment:固定电话;type:varchar(200);"`
      Sjhm  string `json:"sjhm" cn:"手机号码" form:"sjhm" gorm:"column:sjhm;comment:手机号码;type:varchar(200);"`
      Dzyj  string `json:"dzyj" cn:"电子邮件" form:"dzyj" gorm:"column:dzyj;comment:电子邮件;type:varchar(200);"`
      Whcddm  string `json:"whcddm" cn:"文化程度代码" form:"whcddm" gorm:"column:whcddm;comment:文化程度代码;type:varchar(200);"`
      Whcdmc  string `json:"whcdmc" cn:"文化程度名称" form:"whcdmc" gorm:"column:whcdmc;comment:文化程度名称;type:varchar(200);"`
      Zylbdm  string `json:"zylbdm" cn:"职业类别代码" form:"zylbdm" gorm:"column:zylbdm;comment:职业类别代码;type:varchar(200);"`
      Zylbmc  string `json:"zylbmc" cn:"职业类别名称" form:"zylbmc" gorm:"column:zylbmc;comment:职业类别名称;type:varchar(200);"`
      Csdxzqhm  string `json:"csdxzqhm" cn:"出生地行政区划码" form:"csdxzqhm" gorm:"column:csdxzqhm;comment:出生地行政区划码;type:varchar(200);"`
      Csd  string `json:"csd" cn:"出生地" form:"csd" gorm:"column:csd;comment:出生地;type:varchar(200);"`
      Jzdxzqhm  string `json:"jzdxzqhm" cn:"居住地行政区划码" form:"jzdxzqhm" gorm:"column:jzdxzqhm;comment:居住地行政区划码;type:varchar(200);"`
      Jzdz  string `json:"jzdz" cn:"居住地址" form:"jzdz" gorm:"column:jzdz;comment:居住地址;type:varchar(200);"`
      Hkdxzqhm  string `json:"hkdxzqhm" cn:"户口地行政区划码" form:"hkdxzqhm" gorm:"column:hkdxzqhm;comment:户口地行政区划码;type:varchar(200);"`
      Hkdz  string `json:"hkdz" cn:"户口地址" form:"hkdz" gorm:"column:hkdz;comment:户口地址;type:varchar(200);"`
      Lxrxm  string `json:"lxrxm" cn:"联系人姓名" form:"lxrxm" gorm:"column:lxrxm;comment:联系人姓名;type:varchar(200);"`
      Lxrdh  string `json:"lxrdh" cn:"联系人电话" form:"lxrdh" gorm:"column:lxrdh;comment:联系人电话;type:varchar(200);"`
      Abo  string `json:"abo" cn:"ABO 血型" form:"abo" gorm:"column:abo;comment:ABO 血型;type:varchar(200);"`
      Rh  string `json:"rh" cn:"RH 血型" form:"rh" gorm:"column:rh;comment:RH 血型;type:varchar(200);"`
      Cblbdm  string `json:"cblbdm" cn:"参保类别代码" form:"cblbdm" gorm:"column:cblbdm;comment:参保类别代码;type:varchar(200);"`
      Grdaid  string `json:"grdaid" cn:"个人档案ID" form:"grdaid" gorm:"column:grdaid;comment:个人档案ID;type:varchar(200);"`
      Yl1  string `json:"yl1" cn:"预留一" form:"yl1" gorm:"column:yl1;comment:预留一;type:varchar(200);"`
      Yl2  string `json:"yl2" cn:"预留二" form:"yl2" gorm:"column:yl2;comment:预留二;type:varchar(200);"`
      Sjscsj  string `json:"sjscsj" cn:"数据生成时间" form:"sjscsj" gorm:"column:sjscsj;comment:数据生成时间;type:varchar(200);"`
      Tbrqsj  string `json:"tbrqsj" cn:"填报日期" form:"tbrqsj" gorm:"column:tbrqsj;comment:填报日期;type:varchar(200);"`
      Mj  string `json:"mj" cn:"密级" form:"mj" gorm:"column:mj;comment:密级;type:varchar(200);"`
      Cxbz  string `json:"cxbz" cn:"撤销标志" form:"cxbz" gorm:"column:cxbz;comment:撤销标志;type:varchar(200);"`

}


// TableName BbPiPerson 表名
func (BbPiPerson) TableName() string {
  return "bb_pi_person"
}

