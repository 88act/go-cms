// 自动生成模板MemContacts
package model

import (
	. "go-cms/common/baseModel"
	"time"
)

// MemContacts 结构体
type MemContacts struct {
	BaseModel
	UserId   int64      `db:"user_id" json:"userId" gorm:"column:user_id"`       //用户id
	Realname string     `db:"realname" json:"realname" gorm:"column:realname"`   //姓名
	Area     string     `db:"area" json:"area" gorm:"column:area"`               //国家
	Mobile   string     `db:"mobile" json:"mobile" gorm:"column:mobile"`         //手机
	Phone    string     `db:"phone" json:"phone" gorm:"column:phone"`            //固话
	Email    string     `db:"email" json:"email" gorm:"column:email"`            //邮箱地址
	Group    int        `db:"group" json:"group" gorm:"column:group"`            //分组
	Star     int        `db:"star" json:"star" gorm:"column:star"`               //标星
	LastTime *time.Time `db:"last_time" json:"lastTime" gorm:"column:last_time"` //最近联系
	Status   int        `db:"status" json:"status" gorm:"column:status"`         //状态

}

// TableName MemContacts 表名
func (MemContacts) TableName() string {
	return "mem_contacts"
}

// MemContactsSearch  查询
type MemContactsSearch struct {
	BaseModel
	PageInfo
	UserId   *int64     `json:"userId"  form:"userId" `
	Realname string     `json:"realname"   form:"realname" `
	Area     string     `json:"area"   form:"area" `
	Mobile   string     `json:"mobile"   form:"mobile" `
	Phone    string     `json:"phone"   form:"phone" `
	Email    string     `json:"email"   form:"email" `
	Group    *int       `json:"group"   form:"group" `
	Star     *int       `json:"star"   form:"star" `
	LastTime *time.Time `json:"lastTime" form:"lastTime" `
	Status   *int       `json:"status"   form:"status" `
}
