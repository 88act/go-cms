// 自动生成模板BasicSms
package model
import (
 
	. "go-cms/common/baseModel"
)

// BasicSms 结构体
type BasicSms struct {
	BaseModel
	UserId     *int64 `db:"user_id" json:"userId" gorm:"column:user_id"`               //用户id
	UserIdSend *int64 `db:"user_id_send" json:"userIdSend" gorm:"column:user_id_send"` //发送人id
	Phone      string `db:"phone" json:"phone" gorm:"column:phone"`                    //手机
	Title      string `db:"title" json:"title" gorm:"column:title"`                    //标题
	Content    string `db:"content" json:"content" gorm:"column:content"`              //内容
	TempletId  string `db:"templet_id" json:"templetId" gorm:"column:templet_id"`      //模板id
	VerilyCode string `db:"verily_code" json:"verilyCode" gorm:"column:verily_code"`   //验证码
	Attachment string `db:"attachment" json:"attachment" gorm:"column:attachment"`     //附件
	Status     *int   `db:"status" json:"status" gorm:"column:status"`                 //记录状态

}

// TableName BasicSms 表名
func (BasicSms) TableName() string {
	return "basic_sms"
}

// BasicSmsSearch  查询
type BasicSmsSearch struct {
	BasicSms
	PageInfo
}
