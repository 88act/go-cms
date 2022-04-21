// 自动生成模板BbPiInstitution
package business

import (
	"go-cms/global"
)

// BbPiInstitution 结构体
// 如果含有time.Time 请自行import time包
type BbPiInstitution struct {
     BbPiInstitutionMini
}

// BbPiInstitutionMini 结构体
type BbPiInstitutionMini struct {
      global.GO_MODEL
      Status  *int `json:"status" cn:"上传" form:"status" gorm:"column:status;comment:上传:0未上传 1已上传 2上传失败;type:smallint"`
      Jgdm  string `json:"jgdm" cn:"机构标识" form:"jgdm" gorm:"column:jgdm;comment:机构标识;type:varchar(200);"`
      Zzjgdm  string `json:"zzjgdm" cn:"统一信用代码" form:"zzjgdm" gorm:"column:zzjgdm;comment:统一信用代码;type:varchar(200);"`
      Jgmc  string `json:"jgmc" cn:"机构名称" form:"jgmc" gorm:"column:jgmc;comment:机构名称;type:varchar(200);"`
      Xzqhdm  string `json:"xzqhdm" cn:"行政区划代码" form:"xzqhdm" gorm:"column:xzqhdm;comment:行政区划代码;type:varchar(200);"`
      Jglx  string `json:"jglx" cn:"机构类型" form:"jglx" gorm:"column:jglx;comment:机构类型;type:varchar(200);"`
      Jgclrq  string `json:"jgclrq" cn:"机构成立日期" form:"jgclrq" gorm:"column:jgclrq;comment:机构成立日期;type:varchar(200);"`
      Dwlsgxdm  string `json:"dwlsgxdm" cn:"单位隶属关系代码" form:"dwlsgxdm" gorm:"column:dwlsgxdm;comment:单位隶属关系代码;type:varchar(200);"`
      Jgflgllbdm  string `json:"jgflgllbdm" cn:"机构分类代码" form:"jgflgllbdm" gorm:"column:jgflgllbdm;comment:机构分类代码;type:varchar(200);"`
      Jgfldm  string `json:"jgfldm" cn:"机构分类代码" form:"jgfldm" gorm:"column:jgfldm;comment:机构分类代码;type:varchar(200);"`
      Jjlxdm  string `json:"jjlxdm" cn:"经济类型代码" form:"jjlxdm" gorm:"column:jjlxdm;comment:经济类型代码;type:varchar(200);"`
      Dz  string `json:"dz" cn:"地址" form:"dz" gorm:"column:dz;comment:地址;type:varchar(200);"`
      Styyzzjgdm  string `json:"styyzzjgdm" cn:"组织机构代码" form:"styyzzjgdm" gorm:"column:styyzzjgdm;comment:组织机构代码;type:varchar(200);"`
      Styymc  string `json:"styymc" cn:"医院名称" form:"styymc" gorm:"column:styymc;comment:医院名称;type:varchar(200);"`
      Styljgjb  string `json:"styljgjb" cn:"医疗机构级别" form:"styljgjb" gorm:"column:styljgjb;comment:医疗机构级别;type:varchar(200);"`
      Styljgdj  string `json:"styljgdj" cn:"医院机构等级" form:"styljgdj" gorm:"column:styljgdj;comment:医院机构等级;type:varchar(200);"`
      Hlwyywz  string `json:"hlwyywz" cn:"医院网址" form:"hlwyywz" gorm:"column:hlwyywz;comment:医院网址;type:varchar(200);"`
      Xkzhm  string `json:"xkzhm" cn:"许可证号码" form:"xkzhm" gorm:"column:xkzhm;comment:许可证号码;type:varchar(200);"`
      Xkxmmc  string `json:"xkxmmc" cn:"许可项目名称" form:"xkxmmc" gorm:"column:xkxmmc;comment:许可项目名称;type:varchar(200);"`
      Xkzyxq  string `json:"xkzyxq" cn:"许可证有效期" form:"xkzyxq" gorm:"column:xkzyxq;comment:许可证有效期;type:varchar(200);"`
      Xxaqdjbh  string `json:"xxaqdjbh" cn:"信息安等级保护" form:"xxaqdjbh" gorm:"column:xxaqdjbh;comment:信息安等级保护;type:varchar(200);"`
      Xxaqdjbhbh  string `json:"xxaqdjbhbh" cn:"信息安等保编号" form:"xxaqdjbhbh" gorm:"column:xxaqdjbhbh;comment:信息安等保编号;type:varchar(200);"`
      Kbzjjes  string `json:"kbzjjes" cn:"开办资金金额数" form:"kbzjjes" gorm:"column:kbzjjes;comment:开办资金金额数;type:varchar(200);"`
      Frdb  string `json:"frdb" cn:"法人代表" form:"frdb" gorm:"column:frdb;comment:法人代表;type:varchar(200);"`
      Jgszdmzzzdfbz  string `json:"jgszdmzzzdfbz" cn:"机构民族自治标志" form:"jgszdmzzzdfbz" gorm:"column:jgszdmzzzdfbz;comment:机构民族自治标志;type:varchar(200);"`
      Sffzjg  string `json:"sffzjg" cn:"是否分支机构" form:"sffzjg" gorm:"column:sffzjg;comment:是否分支机构;type:varchar(200);"`
      Jgdemc  string `json:"jgdemc" cn:"机构第二名称" form:"jgdemc" gorm:"column:jgdemc;comment:机构第二名称;type:varchar(200);"`
      Jgms  string `json:"jgms" cn:"机构描述" form:"jgms" gorm:"column:jgms;comment:机构描述;type:varchar(200);"`
      Yzbm  string `json:"yzbm" cn:"邮政编码" form:"yzbm" gorm:"column:yzbm;comment:邮政编码;type:varchar(200);"`
      Dhhm  string `json:"dhhm" cn:"电话号码" form:"dhhm" gorm:"column:dhhm;comment:电话号码;type:varchar(200);"`
      Dwdzyx  string `json:"dwdzyx" cn:"电子信箱" form:"dwdzyx" gorm:"column:dwdzyx;comment:电子信箱;type:varchar(200);"`
      Dsfjgdm  string `json:"dsfjgdm" cn:"第三方信用代码" form:"dsfjgdm" gorm:"column:dsfjgdm;comment:第三方信用代码;type:varchar(200);"`
      Dsfjgmc  string `json:"dsfjgmc" cn:"第三方机构名称" form:"dsfjgmc" gorm:"column:dsfjgmc;comment:第三方机构名称;type:varchar(200);"`
      Xyyxq  string `json:"xyyxq" cn:"协议有效期" form:"xyyxq" gorm:"column:xyyxq;comment:协议有效期;type:varchar(200);"`
      Bspt  string `json:"bspt" cn:"部署平台" form:"bspt" gorm:"column:bspt;comment:部署平台;type:varchar(200);"`
      Jgmsxx  string `json:"jgmsxx" cn:"架构描述" form:"jgmsxx" gorm:"column:jgmsxx;comment:架构描述;type:varchar(200);"`
      Jfmj  string `json:"jfmj" cn:"机房面积" form:"jfmj" gorm:"column:jfmj;comment:机房面积;type:varchar(200);"`
      Sfslgd  string `json:"sfslgd" cn:"是否双路发电" form:"sfslgd" gorm:"column:sfslgd;comment:是否双路发电;type:varchar(200);"`
      Bz  string `json:"bz" cn:"备注" form:"bz" gorm:"column:bz;comment:备注;type:varchar(200);"`
      Sjscsj  string `json:"sjscsj" cn:"数据生成日期时间" form:"sjscsj" gorm:"column:sjscsj;comment:数据生成日期时间;type:varchar(200);"`
      Tbrqsj  string `json:"tbrqsj" cn:"填报日期" form:"tbrqsj" gorm:"column:tbrqsj;comment:填报日期;type:varchar(200);"`
      Cxbz  string `json:"cxbz" cn:"撤销标志" form:"cxbz" gorm:"column:cxbz;comment:撤销标志;type:varchar(200);"`

}


// TableName BbPiInstitution 表名
func (BbPiInstitution) TableName() string {
  return "bb_pi_institution"
}

