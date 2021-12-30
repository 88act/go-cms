// 自动生成模板MemUser
package business

import (
	"go-cms/global"
	"time"
)

// MemUser 结构体
// 如果含有time.Time 请自行import time包
type MemUser struct {
	MemUserMini
	Pws               string     `json:"pws" cn:"密码" form:"pws" gorm:"column:pws;comment:密码;type:varchar(32);"`
	PwsSlat           string     `json:"pwsSlat" cn:"密码盐" form:"pwsSlat" gorm:"column:pws_slat;comment:密码盐;type:varchar(32);"`
	CardId            string     `json:"cardId" cn:"身份证" form:"cardId" gorm:"column:card_id;comment:身份证;type:char;"`
	Birthday          *time.Time `json:"birthday" cn:"生日" form:"birthday" gorm:"column:birthday;comment:生日;type:datetime"`
	Avatar            string     `json:"avatar" cn:"头像" form:"avatar" gorm:"column:avatar;comment:头像;type:单图片"`
	MobileValidated   *bool      `json:"mobileValidated" cn:"验证手机" form:"mobileValidated" gorm:"column:mobile_validated;comment:验证手机;type:tinyint"`
	EmailValidated    *bool      `json:"emailValidated" cn:"验证邮箱" form:"emailValidated" gorm:"column:email_validated;comment:验证邮箱;type:tinyint"`
	RealnameValidated *bool      `json:"realnameValidated" cn:"验证实名" form:"realnameValidated" gorm:"column:realname_validated;comment:验证实名;type:tinyint"`
	LoginTimes        *int       `json:"loginTimes" cn:"登录次数" form:"loginTimes" gorm:"column:login_times;comment:登录次数;type:int"`
	RecommendId       *int       `json:"recommendId" cn:"推荐人ID" form:"recommendId" gorm:"column:recommend_id;comment:推荐人ID;type:bigint"`
	RecommendId2      *int       `json:"recommendId2" cn:"推荐人ID2" form:"recommendId2" gorm:"column:recommend_id2;comment:推荐人ID2;type:bigint"`
	RegIp             *int       `json:"regIp" cn:"注册ip" form:"regIp" gorm:"column:reg_ip;comment:注册ip;type:bigint"`
	LastLoginIp       *int       `json:"lastLoginIp" cn:"登录ip" form:"lastLoginIp" gorm:"column:last_login_ip;comment:登录ip;type:bigint"`
}

// MemUserMini 结构体
type MemUserMini struct {
	global.GO_MODEL
	Username      string     `json:"username" cn:"用户名" form:"username" gorm:"column:username;comment:用户名;type:varchar(40);"`
	Email         string     `json:"email" cn:"邮件" form:"email" gorm:"column:email;comment:邮件;type:varchar(60);"`
	Mobile        string     `json:"mobile" cn:"手机" form:"mobile" gorm:"column:mobile;comment:手机;type:varchar(20);"`
	Nickname      string     `json:"nickname" cn:"昵称" form:"nickname" gorm:"column:nickname;comment:昵称;type:varchar(50);"`
	Realname      string     `json:"realname" cn:"真实名" form:"realname" gorm:"column:realname;comment:真实名;type:varchar(20);"`
	Sex           *int       `json:"sex" cn:"0 保密 1 男 2 女" form:"sex" gorm:"column:sex;comment:0 保密 1 男 2 女;type:smallint"`
	RecommendCode string     `json:"recommendCode" cn:"推荐码16位（自己的）" form:"recommendCode" gorm:"column:recommend_code;comment:推荐码16位（自己的）;type:char;"`
	Status        *int       `json:"status" cn:"状态" form:"status" gorm:"column:status;comment:状态:0=审核中,1=审核通过 2=审核未通过3=禁止登录4=发色情非法信息5:已注销6:非法攻击;type:smallint"`
	StatusSafe    *int       `json:"statusSafe" cn:"安全状态" form:"statusSafe" gorm:"column:status_safe;comment:安全状态:0=正常，1=修改了密码 2=修改了手机号;type:smallint"`
	LastLoginTime *time.Time `json:"lastLoginTime" cn:"最后登录时间" form:"lastLoginTime" gorm:"column:last_login_time;comment:最后登录时间;type:datetime"`
}

// TableName MemUser 表名
func (MemUser) TableName() string {
	return "mem_user"
}
