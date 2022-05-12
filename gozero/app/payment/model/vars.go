package model

import (
	"database/sql"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

//统一model 执行接口
type Executable interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

var ErrNotFound = sqlx.ErrNotFound

//支付业务类型
var ThirdPaymentServiceTypeHomestayOrder string = "homestayOrder" //民宿支付

//支付方式
var ThirdPaymentPayModelWechatPay = "WECHAT_PAY" //微信支付

//支付状态 -1支付失败 0待支付 1支付成功 2退款审核中 3退款已审核 4已退款
var PayStatus_Fail int32 = -1     //支付失败
var PayStatus_WaitPay int32 = 0   //待支付
var PayStatus_PayOk int32 = 1     //支付成功
var PayStatus_Refunding int32 = 2 // 退款审核中
var PayStatus_Verify int32 = 3    // 退款已审核
var PayStatus_refunded int32 = 2  //已退款
