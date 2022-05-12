// 自动生成模板PayPayment
package model

import (
	"time"
)

// PayPayment 结构体
type PayPayment struct {
	BaseModel
	Sn             string     `db:"sn" json:"sn" gorm:"column:sn"`                                         //流水单号
	Version        *int64     `db:"version" json:"version" gorm:"column:version"`                          //版本号
	UserId         *int64     `db:"user_id" json:"userId" gorm:"column:user_id"`                           //用户id
	PayMode        string     `db:"pay_mode" json:"payMode" gorm:"column:pay_mode"`                        //支付方式
	TradeType      string     `db:"trade_type" json:"tradeType" gorm:"column:trade_type"`                  //支付类型
	TradeState     string     `db:"trade_state" json:"tradeState" gorm:"column:trade_state"`               //交易状态
	PayTotal       *int64     `db:"pay_total" json:"payTotal" gorm:"column:pay_total"`                     //支付金额(分)
	TransactionId  string     `db:"transaction_id" json:"transactionId" gorm:"column:transaction_id"`      //第三方支付单号
	TradeStateDesc string     `db:"trade_state_desc" json:"tradeStateDesc" gorm:"column:trade_state_desc"` //支付状态描述
	OrderSn        string     `db:"order_sn" json:"orderSn" gorm:"column:order_sn"`                        //业务单号
	ServiceType    string     `db:"service_type" json:"serviceType" gorm:"column:service_type"`            //业务类型
	PayStatus      *int       `db:"pay_status" json:"payStatus" gorm:"column:pay_status"`                  //平台交易状态
	PayTime        *time.Time `db:"pay_time" json:"payTime" gorm:"column:pay_time"`                        //支付时间
	Status         *int       `db:"status" json:"status" gorm:"column:status"`                             //状态

}

// TableName PayPayment 表名
func (PayPayment) TableName() string {
	return "pay_payment"
}

// PayPaymentSearch  查询
type PayPaymentSearch struct {
	PayPayment
	PageInfo
}
