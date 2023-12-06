// 自动生成模板BasicEmail
package model

import (
	. "go-cms/common/baseModel"
)

// BasicEmail 结构体
type BasicEmail struct {
	BaseModel
	UserId     *int64 `db:"user_id" json:"userId" gorm:"column:user_id"`               //用户id
	UserIdSend *int64 `db:"user_id_send" json:"userIdSend" gorm:"column:user_id_send"` //发送人id
	Email      string `db:"email" json:"email" gorm:"column:email"`                    //收件邮箱
	Title      string `db:"title" json:"title" gorm:"column:title"`                    //标题
	Content    string `db:"content" json:"content" gorm:"column:content"`              //内容
	VerilyCode string `db:"verily_code" json:"verilyCode" gorm:"column:verily_code"`   //验证码
	Attachment string `db:"attachment" json:"attachment" gorm:"column:attachment"`     //附件
	Status     *int   `db:"status" json:"status" gorm:"column:status"`                 //记录状态

}

// TableName BasicEmail 表名
func (BasicEmail) TableName() string {
	return "basic_email"
}

// BasicEmailSearch  查询
type BasicEmailSearch struct {
	BasicEmail
	PageInfo
}
