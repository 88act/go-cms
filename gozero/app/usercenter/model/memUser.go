// 自动生成模板MemUser
package model

import (
	. "go-cms/common/baseModel"
	"time"
)

// 表格字段
const Field_MemUser_mini = "*" // "id,created_at,updated_at,unionid,openid,username,password,password_slat,email,mobile,nickname,realname,card_id,sex,birthday,avatar,job_id,branch_id,mobile_validated,email_validated,cardid_validated,remark,recommend_uid,recommend_code,money,money_jf,vip,vip_end,status,status_safe,reg_ip,login_ip,login_total,login_time,step,setting,city,province,country,

// MemUser 结构体
type MemUser struct {
	BaseModel
	CuId            int64      `json:"cuId" form:"cuId" cn:"客户id"  gorm:"column:cu_id;comment:客户id;type:bigint"`
	UserId          int64      `json:"userId" form:"userId" cn:"录入人"  gorm:"column:user_id;comment:录入人;type:bigint"`
	Guid            string     `json:"guid" form:"guid" cn:"唯一id"  gorm:"<-:create;column:guid;comment:唯一id;type:varchar(32);"`
	Username        string     `json:"username" form:"username" cn:"用户名"  gorm:"column:username;comment:用户名;type:varchar(40);"`
	Password        string     `json:"-" form:"password" cn:"密码"  gorm:"<-:create;column:password;comment:密码;type:varchar(32);"`
	PasswordSlat    string     `json:"-" form:"passwordSlat" cn:"密码盐"  gorm:"<-:create;column:password_slat;comment:密码盐;type:varchar(32);"`
	Nickname        string     `json:"nickname" form:"nickname" cn:"昵称"  gorm:"column:nickname;comment:昵称;type:varchar(200);"`
	Realname        string     `json:"realname" form:"realname" cn:"真实名"  gorm:"column:realname;comment:真实名;type:varchar(200);"`
	RoleList        string     `json:"roleList" form:"roleList" cn:"角色list"  gorm:"column:role_list;comment:角色list;type:varchar(255);"`
	Role            int64      `json:"role" form:"role" cn:"角色"  gorm:"column:role;comment:角色;type:bigint"`
	UserType        int        `json:"userType" form:"userType" cn:"类型"   gorm:"column:user_type;comment:类型:1员工 2管理员;type:smallint"`
	Email           string     `json:"email" form:"email" cn:"邮件"  gorm:"column:email;comment:邮件;type:varchar(60);"`
	Mobile          string     `json:"mobile" form:"mobile" cn:"手机"  gorm:"column:mobile;comment:手机;type:varchar(20);"`
	CardId          string     `json:"cardId" form:"cardId" cn:"身份证"  gorm:"column:card_id;comment:身份证;type:varchar(18);"`
	Sex             int        `json:"sex" form:"sex" cn:"性别"   gorm:"column:sex;comment:性别: 0保密 1 男 2 女;type:smallint"`
	Birthday        *time.Time `json:"birthday" form:"birthday" cn:"生日"  gorm:"column:birthday;comment:生日;type:datetime"`
	Avatar          string     `json:"avatar"  form:"avatar"  cn:"头像" gorm:"column:avatar;comment:头像;type:单图片"`
	JobId           int64      `json:"jobId" form:"jobId" cn:"岗位"  gorm:"column:job_id;comment:岗位;type:bigint"`
	DepartId        int64      `json:"departId" form:"departId" cn:"部门"  gorm:"column:depart_id;comment:部门;type:bigint"`
	MobileValidated bool       `json:"mobileValidated" form:"mobileValidated" cn:"验证手机"   gorm:"column:mobile_validated;comment:验证手机;type:tinyint"`
	EmailValidated  bool       `json:"emailValidated" form:"emailValidated" cn:"验证邮箱"   gorm:"column:email_validated;comment:验证邮箱;type:tinyint"`
	CardidValidated bool       `json:"cardidValidated" form:"cardidValidated" cn:"验证实名"   gorm:"column:cardid_validated;comment:验证实名;type:tinyint"`
	Remark          string     `json:"remark" form:"remark" cn:"备注"  gorm:"column:remark;comment:备注;type:varchar(255);"`
	StatusSafe      int        `json:"statusSafe" form:"statusSafe" cn:"安全状态"   gorm:"column:status_safe;comment:安全状态:0=正常，1=修改了密码 2=修改了手机号;type:smallint"`
	RegIp           int        `json:"regIp" form:"regIp" cn:"注册ip"   gorm:"column:reg_ip;comment:注册ip;type:int"`
	LoginIp         int        `json:"loginIp" form:"loginIp" cn:"登录ip"   gorm:"column:login_ip;comment:登录ip;type:int"`
	LoginTotal      int        `json:"loginTotal" form:"loginTotal" cn:"登录次数"   gorm:"column:login_total;comment:登录次数;type:int"`
	LoginTime       *time.Time `json:"loginTime" form:"loginTime" cn:"最后登录时间"  gorm:"column:login_time;comment:最后登录时间;type:datetime"`
	Scan            string     `json:"scan" form:"scan" cn:"扫码"  gorm:"column:scan;comment:扫码;type:varchar(32);"`
	ScanTime        *time.Time `json:"scanTime" form:"scanTime" cn:"扫码"  gorm:"column:scan_time;comment:扫码;type:datetime"`
	Setting         string     `json:"setting" form:"setting" cn:"设置值"  gorm:"column:setting;comment:设置值;type:varchar(1000);"`
	RtcModel        int        `json:"rtcModel" form:"rtcModel" cn:"远程协助模式"   gorm:"column:rtc_model;comment:远程协助模式;type:smallint"`
	Status          int        `json:"status" form:"status" cn:"状态"   gorm:"column:status;comment:状态:0=审核中,1=审核通过 2=审核未通过3=禁止登录4=非法信息5:已注销6:非法攻击;type:smallint"`
}

// TableName MemUser 表名
func (MemUser) TableName() string {
	return "mem_user"
}

type MemUserSearch struct {
	BaseModel
	PageInfo
	CuId            *int64     `json:"cuId"  form:"cuId" `
	UserId          *int64     `json:"userId"  form:"userId" `
	Guid            string     `json:"guid"   form:"guid" `
	Username        string     `json:"username"   form:"username" `
	Password        string     `json:"password"   form:"password" `
	PasswordSlat    string     `json:"passwordSlat"   form:"passwordSlat" `
	Nickname        string     `json:"nickname"   form:"nickname" `
	Realname        string     `json:"realname"   form:"realname" `
	RoleList        string     `json:"roleList"   form:"roleList" `
	Role            *int64     `json:"role"  form:"role" `
	UserType        *int       `json:"userType"   form:"userType" `
	Email           string     `json:"email"   form:"email" `
	Mobile          string     `json:"mobile"   form:"mobile" `
	CardId          string     `json:"cardId"   form:"cardId" `
	Sex             *int       `json:"sex"   form:"sex" `
	Birthday        *time.Time `json:"birthday" form:"birthday" `
	Avatar          string     `json:"avatar" form:"avatar" `
	JobId           *int64     `json:"jobId"  form:"jobId" `
	DepartId        *int64     `json:"departId"  form:"departId" `
	MobileValidated *bool      `json:"mobileValidated" form:"mobileValidated" `
	EmailValidated  *bool      `json:"emailValidated" form:"emailValidated" `
	CardidValidated *bool      `json:"cardidValidated" form:"cardidValidated" `
	Remark          string     `json:"remark"   form:"remark" `
	StatusSafe      *int       `json:"statusSafe"   form:"statusSafe" `
	RegIp           *int       `json:"regIp"   form:"regIp" `
	LoginIp         *int       `json:"loginIp"   form:"loginIp" `
	LoginTotal      *int       `json:"loginTotal"   form:"loginTotal" `
	LoginTime       *time.Time `json:"loginTime" form:"loginTime" `
	Scan            string     `json:"scan"   form:"scan" `
	ScanTime        *time.Time `json:"scanTime" form:"scanTime" `
	Setting         string     `json:"setting"   form:"setting" `
	RtcModel        *int       `json:"rtcModel"   form:"rtcModel" `
	Status          *int       `json:"status"   form:"status" `
}
