// 自动生成模板BbPiTreatmentOrder
package business

import (
	"go-cms/global"
)

// BbPiTreatmentOrder 结构体
// 如果含有time.Time 请自行import time包
type BbPiTreatmentOrder struct {
     BbPiTreatmentOrderMini
}

// BbPiTreatmentOrderMini 结构体
type BbPiTreatmentOrderMini struct {
      global.GO_MODEL
      Status  *int `json:"status" cn:"上传" form:"status" gorm:"column:status;comment:上传:0未上传 1已上传 2上传失败;type:smallint"`
      Jgdm  string `json:"jgdm" cn:"机构标识" form:"jgdm" gorm:"column:jgdm;comment:机构标识;type:varchar(200);"`
      Fwwddm  string `json:"fwwddm" cn:"服务网点代码" form:"fwwddm" gorm:"column:fwwddm;comment:服务网点代码;type:varchar(200);"`
      Cfbh  string `json:"cfbh" cn:"处方编号" form:"cfbh" gorm:"column:cfbh;comment:处方编号;type:varchar(200);"`
      Cfmxid  string `json:"cfmxid" cn:"处方明细ID" form:"cfmxid" gorm:"column:cfmxid;comment:处方明细ID;type:varchar(200);"`
      Kh  string `json:"kh" cn:"卡号" form:"kh" gorm:"column:kh;comment:卡号;type:varchar(200);"`
      Klx  string `json:"klx" cn:"卡类型" form:"klx" gorm:"column:klx;comment:卡类型;type:varchar(200);"`
      Cfklsj  string `json:"cfklsj" cn:"处方开立日期" form:"cfklsj" gorm:"column:cfklsj;comment:处方开立日期;type:varchar(200);"`
      Cfyxts  string `json:"cfyxts" cn:"处方有效天数" form:"cfyxts" gorm:"column:cfyxts;comment:处方有效天数;type:varchar(200);"`
      Cfklksbm  string `json:"cfklksbm" cn:"处方科室编码" form:"cfklksbm" gorm:"column:cfklksbm;comment:处方科室编码;type:varchar(200);"`
      Cfklksmc  string `json:"cfklksmc" cn:"处方科室名称" form:"cfklksmc" gorm:"column:cfklksmc;comment:处方科室名称;type:varchar(200);"`
      Mzh  string `json:"mzh" cn:"门诊号" form:"mzh" gorm:"column:mzh;comment:门诊号;type:varchar(200);"`
      Xm  string `json:"xm" cn:"姓名" form:"xm" gorm:"column:xm;comment:姓名;type:varchar(200);"`
      Xbdm  string `json:"xbdm" cn:"性别代码" form:"xbdm" gorm:"column:xbdm;comment:性别代码;type:varchar(200);"`
      Csrq  string `json:"csrq" cn:"出生日期" form:"csrq" gorm:"column:csrq;comment:出生日期;type:varchar(200);"`
      Nls  string `json:"nls" cn:"年龄岁" form:"nls" gorm:"column:nls;comment:年龄岁;type:varchar(200);"`
      Nly  string `json:"nly" cn:"年龄月" form:"nly" gorm:"column:nly;comment:年龄月;type:varchar(200);"`
      Jzrqsj  string `json:"jzrqsj" cn:"就诊日期时间" form:"jzrqsj" gorm:"column:jzrqsj;comment:就诊日期时间;type:varchar(200);"`
      Yzsm  string `json:"yzsm" cn:"医嘱说明" form:"yzsm" gorm:"column:yzsm;comment:医嘱说明;type:varchar(200);"`
      Pxh  string `json:"pxh" cn:"排序号" form:"pxh" gorm:"column:pxh;comment:排序号;type:varchar(200);"`
      Yzxmlxdm  string `json:"yzxmlxdm" cn:"医嘱项目类型代码" form:"yzxmlxdm" gorm:"column:yzxmlxdm;comment:医嘱项目类型代码;type:varchar(200);"`
      Ypcfsx  string `json:"ypcfsx" cn:"药品处方属性" form:"ypcfsx" gorm:"column:ypcfsx;comment:药品处方属性;type:varchar(200);"`
      Zylbdm  string `json:"zylbdm" cn:"中药类别代码" form:"zylbdm" gorm:"column:zylbdm;comment:中药类别代码;type:varchar(200);"`
      Cfmxxmbm  string `json:"cfmxxmbm" cn:"处方明细医保编码" form:"cfmxxmbm" gorm:"column:cfmxxmbm;comment:处方明细医保编码;type:varchar(200);"`
      Ypid  string `json:"ypid" cn:"药监药品ID" form:"ypid" gorm:"column:ypid;comment:药监药品ID;type:varchar(200);"`
      Cfmxmc  string `json:"cfmxmc" cn:"处方明细名称" form:"cfmxmc" gorm:"column:cfmxmc;comment:处方明细名称;type:varchar(200);"`
      Ywmc  string `json:"ywmc" cn:"药物名称" form:"ywmc" gorm:"column:ywmc;comment:药物名称;type:varchar(200);"`
      Ypgg  string `json:"ypgg" cn:"药品规格" form:"ypgg" gorm:"column:ypgg;comment:药品规格;type:varchar(200);"`
      Dddz  string `json:"dddz" cn:"DDD值" form:"dddz" gorm:"column:dddz;comment:DDD值;type:varchar(200);"`
      Ywjxdm  string `json:"ywjxdm" cn:"药物剂型代码" form:"ywjxdm" gorm:"column:ywjxdm;comment:药物剂型代码;type:varchar(200);"`
      Ywsycjl  string `json:"ywsycjl" cn:"药物使用次剂量" form:"ywsycjl" gorm:"column:ywsycjl;comment:药物使用次剂量;type:varchar(200);"`
      Ywsyjldw  string `json:"ywsyjldw" cn:"药物使用剂量单位" form:"ywsyjldw" gorm:"column:ywsyjldw;comment:药物使用剂量单位;type:varchar(200);"`
      Ywsypcdm  string `json:"ywsypcdm" cn:"药物使用频次代码" form:"ywsypcdm" gorm:"column:ywsypcdm;comment:药物使用频次代码;type:varchar(200);"`
      Ywsypcmc  string `json:"ywsypcmc" cn:"药物使用频次名称" form:"ywsypcmc" gorm:"column:ywsypcmc;comment:药物使用频次名称;type:varchar(200);"`
      Yytjdm  string `json:"yytjdm" cn:"用药途径代码" form:"yytjdm" gorm:"column:yytjdm;comment:用药途径代码;type:varchar(200);"`
      Yytjmc  string `json:"yytjmc" cn:"用药途径名称" form:"yytjmc" gorm:"column:yytjmc;comment:用药途径名称;type:varchar(200);"`
      Ywsyzjl  string `json:"ywsyzjl" cn:"药物使用总剂量" form:"ywsyzjl" gorm:"column:ywsyzjl;comment:药物使用总剂量;type:varchar(200);"`
      Cfypzh  string `json:"cfypzh" cn:"处方药品组号" form:"cfypzh" gorm:"column:cfypzh;comment:处方药品组号;type:varchar(200);"`
      Zyypcf  string `json:"zyypcf" cn:"中药饮片处方" form:"zyypcf" gorm:"column:zyypcf;comment:中药饮片处方;type:varchar(200);"`
      Zyypjs  string `json:"zyypjs" cn:"中药饮片剂数" form:"zyypjs" gorm:"column:zyypjs;comment:中药饮片剂数;type:varchar(200);"`
      Zyypjzf  string `json:"zyypjzf" cn:"中药饮片煎煮法" form:"zyypjzf" gorm:"column:zyypjzf;comment:中药饮片煎煮法;type:varchar(200);"`
      Zyyyff  string `json:"zyyyff" cn:"中药用药方法" form:"zyyyff" gorm:"column:zyyyff;comment:中药用药方法;type:varchar(200);"`
      Fyjl  string `json:"fyjl" cn:"发药剂量" form:"fyjl" gorm:"column:fyjl;comment:发药剂量;type:varchar(200);"`
      Fyjldw  string `json:"fyjldw" cn:"发药剂量单位" form:"fyjldw" gorm:"column:fyjldw;comment:发药剂量单位;type:varchar(200);"`
      Dj  string `json:"dj" cn:"单价" form:"dj" gorm:"column:dj;comment:单价;type:varchar(200);"`
      Zje  string `json:"zje" cn:"总金额" form:"zje" gorm:"column:zje;comment:总金额;type:varchar(200);"`
      Cfklysgh  string `json:"cfklysgh" cn:"处方开立医师工号" form:"cfklysgh" gorm:"column:cfklysgh;comment:处方开立医师工号;type:varchar(200);"`
      Cfklysqm  string `json:"cfklysqm" cn:"处方开立医师签名" form:"cfklysqm" gorm:"column:cfklysqm;comment:处方开立医师签名;type:varchar(200);"`
      Cfshyjsgh  string `json:"cfshyjsgh" cn:"处方审核工号" form:"cfshyjsgh" gorm:"column:cfshyjsgh;comment:处方审核工号;type:varchar(200);"`
      Cfshyjsqm  string `json:"cfshyjsqm" cn:"处方审核签名" form:"cfshyjsqm" gorm:"column:cfshyjsqm;comment:处方审核签名;type:varchar(200);"`
      Cftpyjsgh  string `json:"cftpyjsgh" cn:"处方调配工号" form:"cftpyjsgh" gorm:"column:cftpyjsgh;comment:处方调配工号;type:varchar(200);"`
      Cftpyjsqm  string `json:"cftpyjsqm" cn:"处方调配签名" form:"cftpyjsqm" gorm:"column:cftpyjsqm;comment:处方调配签名;type:varchar(200);"`
      Cfhdyjsgh  string `json:"cfhdyjsgh" cn:"处方核对工号" form:"cfhdyjsgh" gorm:"column:cfhdyjsgh;comment:处方核对工号;type:varchar(200);"`
      Cfhdyjsqm  string `json:"cfhdyjsqm" cn:"处方核对签名" form:"cfhdyjsqm" gorm:"column:cfhdyjsqm;comment:处方核对签名;type:varchar(200);"`
      Cffyyjsgh  string `json:"cffyyjsgh" cn:"处方发药工号" form:"cffyyjsgh" gorm:"column:cffyyjsgh;comment:处方发药工号;type:varchar(200);"`
      Cffyyjsqm  string `json:"cffyyjsqm" cn:"处方发药签名" form:"cffyyjsqm" gorm:"column:cffyyjsqm;comment:处方发药签名;type:varchar(200);"`
      Zxjg  string `json:"zxjg" cn:"执行结果" form:"zxjg" gorm:"column:zxjg;comment:执行结果;type:varchar(200);"`
      Bz  string `json:"bz" cn:"备注" form:"bz" gorm:"column:bz;comment:备注;type:varchar(200);"`
      Qyjgdm  string `json:"qyjgdm" cn:"取药机构代码" form:"qyjgdm" gorm:"column:qyjgdm;comment:取药机构代码;type:varchar(200);"`
      Qyjgmc  string `json:"qyjgmc" cn:"取药机构名称" form:"qyjgmc" gorm:"column:qyjgmc;comment:取药机构名称;type:varchar(200);"`
      Sjscsj  string `json:"sjscsj" cn:"数据生成时间" form:"sjscsj" gorm:"column:sjscsj;comment:数据生成时间;type:varchar(200);"`
      Tbrqsj  string `json:"tbrqsj" cn:"填报日期" form:"tbrqsj" gorm:"column:tbrqsj;comment:填报日期;type:varchar(200);"`
      Cxbz  string `json:"cxbz" cn:"撤销标志" form:"cxbz" gorm:"column:cxbz;comment:撤销标志;type:varchar(200);"`
      Mj  string `json:"mj" cn:"密级" form:"mj" gorm:"column:mj;comment:密级;type:varchar(200);"`

}


// TableName BbPiTreatmentOrder 表名
func (BbPiTreatmentOrder) TableName() string {
  return "bb_pi_treatment_order"
}

