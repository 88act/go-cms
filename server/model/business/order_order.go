// 自动生成模板OrderOrder
package business

import (
	"go-cms/global"
)

// OrderOrder 结构体 
type OrderOrder struct {
     OrderOrderMini
}

// OrderOrderMini 结构体
type OrderOrderMini struct {
      global.GO_MODEL
      Sn  string `json:"sn" cn:"订单号" form:"sn" gorm:"column:sn;comment:订单号;type:char;"`
      UserId  *int `json:"userId" cn:"用户id" form:"userId" gorm:"column:user_id;comment:用户id;type:bigint"`
      OrderType  *int `json:"orderType" cn:"订单类型" form:"orderType" gorm:"column:order_type;comment:订单类型:1活动,2实物商品,3企业vip3,4个人vip,5抽奖红包;type:smallint"`
      ObjId  *int `json:"objId" cn:"对象id" form:"objId" gorm:"column:obj_id;comment:对象id;type:bigint"`
      Price  *int `json:"price" cn:"实付价" form:"price" gorm:"column:price;comment:实付价;type:int"`
      PriceSrc  *int `json:"priceSrc" cn:"原价" form:"priceSrc" gorm:"column:price_src;comment:原价;type:int"`
      CouponId  *int `json:"couponId" cn:"优惠券ID" form:"couponId" gorm:"column:coupon_id;comment:优惠券ID;type:int"`
      Remark  string `json:"remark" cn:"备注" form:"remark" gorm:"column:remark;comment:备注;type:varchar(500);"`
      OrderCode  string `json:"orderCode" cn:"核销码" form:"orderCode" gorm:"column:order_code;comment:核销码;type:char;"`
      StatusPay  *int `json:"statusPay" cn:"支付状态" form:"statusPay" gorm:"column:status_pay;comment:支付状态:0未支付 1已支付 2 退款审核中 3退款已审核  -1已退款;type:smallint"`
      StatusOrder  *int `json:"statusOrder" cn:"订单状态" form:"statusOrder" gorm:"column:status_order;comment:订单状态:-1已取消 0未支付 1已支付  2已使用(已收货) 3已退款  4已过期 5已评论 6已完成;type:smallint"`
      Status  *int `json:"status" cn:"记录状态" form:"status" gorm:"column:status;comment:记录状态:0未审核 1审核通过 2未通过审核 3黑名单;type:smallint"`

}


// TableName OrderOrder 表名
func (OrderOrder) TableName() string {
  return "order_order"
}

