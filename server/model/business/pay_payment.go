// 自动生成模板PayPayment
package business

import (
	"go-cms/global"
	"time"
)

// PayPayment 结构体
type PayPayment struct {
	PayPaymentMini
}

// PayPaymentMini 结构体
type PayPaymentMini struct {
	global.GO_MODEL
	Sn             string     `json:"sn" cn:"订单号" form:"sn" gorm:"column:sn;comment:订单号;type:char;"`
	UserId         *int       `json:"userId" cn:"用户id" form:"userId" gorm:"column:user_id;comment:用户id;type:bigint"`
	PayMode        string     `json:"payMode" cn:"支付方式" form:"payMode" gorm:"column:pay_mode;comment:支付方式:1微信2支付宝3网银;type:varchar(20);"`
	TradeType      string     `json:"tradeType" cn:"三方支付类型" form:"tradeType" gorm:"column:trade_type;comment:三方支付类型;type:varchar(20);"`
	TradeState     string     `json:"tradeState" cn:"三方交易状态" form:"tradeState" gorm:"column:trade_state;comment:三方交易状态;type:varchar(20);"`
	Price          *int       `json:"price" cn:"总金额" form:"price" gorm:"column:price;comment:总金额;type:int"`
	TransactionId  string     `json:"transactionId" cn:"三方支付单号" form:"transactionId" gorm:"column:transaction_id;comment:三方支付单号;type:char;"`
	TradeStateDesc string     `json:"tradeStateDesc" cn:"支付状态" form:"tradeStateDesc" gorm:"column:trade_state_desc;comment:支付状态;type:varchar(256);"`
	OrderSn        string     `json:"orderSn" cn:"业务单号" form:"orderSn" gorm:"column:order_sn;comment:业务单号;type:char;"`
	ServiceType    string     `json:"serviceType" cn:"业务类型" form:"serviceType" gorm:"column:service_type;comment:业务类型;type:varchar(32);"`
	StatusPay      *int       `json:"statusPay" cn:"支付状态" form:"statusPay" gorm:"column:status_pay;comment:支付状态:-1:支付失败 0:未支付 1:支付成功 2:已退款;type:smallint"`
	PayTime        *time.Time `json:"payTime" cn:"支付时间" form:"payTime" gorm:"column:pay_time;comment:支付时间;type:datetime"`
	Status         *int       `json:"status" cn:"状态" form:"status" gorm:"column:status;comment:状态:0未审核 1审核通过 2未通过审核 3黑名单;type:smallint"`
}

// TableName PayPayment 表名
func (PayPayment) TableName() string {
	return "pay_payment"
}
