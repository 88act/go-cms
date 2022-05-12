package model

import (
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound

//  Order 交易状态 :  -1: 已取消 0:待支付 1:未使用 2:已使用  3:已过期
var OrderStatus_Cancel int32 = -1
var OrderStatus_WaitPay int32 = 0
var OrderStatus_WaitUse int32 = 1
var OrderStatus_Used int32 = 2
var OrderStatus_Refund int32 = 3
var OrderStatus_Expire int32 = 4
