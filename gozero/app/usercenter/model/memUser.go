// 自动生成模板MemUser
package model

import (
	"time"
)

// MemUser 结构体
type MemUser struct {
	BaseModel
	Username        string     `json:"username" cn:"用户名" form:"username" gorm:"column:username;comment:用户名;type:varchar(40);"`
	Password        string     `json:"password" cn:"密码" form:"password" gorm:"column:password;comment:密码;type:varchar(32);"`
	PasswordSlat    string     `json:"passwordSlat" cn:"密码盐" form:"passwordSlat" gorm:"column:password_slat;comment:密码盐;type:varchar(32);"`
	Email           string     `json:"email" cn:"邮件" form:"email" gorm:"column:email;comment:邮件;type:varchar(60);"`
	Mobile          string     `json:"mobile" cn:"手机" form:"mobile" gorm:"column:mobile;comment:手机;type:varchar(20);"`
	Nickname        string     `json:"nickname" cn:"昵称" form:"nickname" gorm:"column:nickname;comment:昵称;type:varchar(50);"`
	Realname        string     `json:"realname" cn:"真实名" form:"realname" gorm:"column:realname;comment:真实名;type:varchar(30);"`
	CardId          string     `json:"cardId" cn:"身份证" form:"cardId" gorm:"column:card_id;comment:身份证;type:char;"`
	Sex             *int       `json:"sex" cn:"性别" form:"sex" gorm:"column:sex;comment:性别:0 保密 1 男 2 女;type:smallint"`
	Birthday        *time.Time `json:"birthday" cn:"生日" form:"birthday" gorm:"column:birthday;comment:生日;type:datetime"`
	Avatar          string     `json:"avatar" cn:"头像" form:"avatar" gorm:"column:avatar;comment:头像;type:varchar(256);"`
	MobileValidated *bool      `json:"mobileValidated" cn:"验证手机" form:"mobileValidated" gorm:"column:mobile_validated;comment:验证手机;type:tinyint"`
	EmailValidated  *bool      `json:"emailValidated" cn:"验证邮箱" form:"emailValidated" gorm:"column:email_validated;comment:验证邮箱;type:tinyint"`
	CardidValidated *bool      `json:"cardidValidated" cn:"验证实名" form:"cardidValidated" gorm:"column:cardid_validated;comment:验证实名;type:tinyint"`
	Info            string     `json:"info" cn:"备注" form:"info" gorm:"column:info;comment:备注;type:varchar(255);"`
	RecommendCode   string     `json:"recommendCode" cn:"推荐码16位（自己的）" form:"recommendCode" gorm:"column:recommend_code;comment:推荐码16位（自己的）;type:varchar(32);"`
	Status          *int       `json:"status" cn:"状态" form:"status" gorm:"column:status;comment:状态:0=审核中,1=审核通过 2=审核未通过3=禁止登录4=发色情非法信息5:已注销6:非法攻击;type:smallint"`
	StatusSafe      *int       `json:"statusSafe" cn:"安全状态" form:"statusSafe" gorm:"column:status_safe;comment:安全状态:0=正常，1=修改了密码 2=修改了手机号;type:smallint"`
	RegIp           *int       `json:"regIp" cn:"注册ip" form:"regIp" gorm:"column:reg_ip;comment:注册ip;type:int"`
	LoginIp         *int       `json:"loginIp" cn:"登录ip" form:"loginIp" gorm:"column:login_ip;comment:登录ip;type:int"`
	LoginTotal      *int       `json:"loginTotal" cn:"登录次数" form:"loginTotal" gorm:"column:login_total;comment:登录次数;type:int"`
	LoginTime       *time.Time `json:"loginTime" cn:"最后登录时间" form:"loginTime" gorm:"column:login_time;comment:最后登录时间;type:datetime"`
}

// TableName MemUser 表名
func (MemUser) TableName() string {
	return "mem_user"
}

// MemUserSearch  查询
type MemUserSearch struct {
	MemUser
	PageInfo
}
