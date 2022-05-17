package model

import (
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound

//  Order 交易状态 :  -1: 已取消 0:待支付 1:已支付 2:已使用 3:已经退款 4:已过期
var OrderStatus_Cancel int32 = -1 //-1: 已取消
var OrderStatus_WaitPay int32 = 0 //0:待支付
var OrderStatus_Payed int32 = 1   //1:已支付
var OrderStatus_Used int32 = 2    //2:已使用
var OrderStatus_Refund int32 = 3  //3:已经退款
var OrderStatus_Expire int32 = 4  //4:已过期
