// 自动生成模板OrderOrder
package model

// OrderOrder 结构体
type OrderOrder struct {
	BaseModel
	Sn          string `db:"sn" json:"sn" gorm:"column:sn"`                              //订单号
	UserId      *int64 `db:"user_id" json:"userId" gorm:"column:user_id"`                //用户id
	OrderType   *int   `db:"order_type" json:"orderType" gorm:"column:order_type"`       //订单类型
	ObjId       *int64 `db:"obj_id" json:"objId" gorm:"column:obj_id"`                   //对象id
	Price       *int   `db:"price" json:"price" gorm:"column:price"`                     //实付价
	PriceSrc    *int   `db:"price_src" json:"priceSrc" gorm:"column:price_src"`          //原价
	CouponId    *int   `db:"coupon_id" json:"couponId" gorm:"column:coupon_id"`          //优惠券ID
	Remark      string `db:"remark" json:"remark" gorm:"column:remark"`                  //备注
	OrderCode   string `db:"order_code" json:"orderCode" gorm:"column:order_code"`       //核销码
	StatusPay   *int   `db:"status_pay" json:"statusPay" gorm:"column:status_pay"`       //支付状态
	StatusOrder *int   `db:"status_order" json:"statusOrder" gorm:"column:status_order"` //订单状态
	Status      *int   `db:"status" json:"status" gorm:"column:status"`
}

// TableName OrderOrder 表名
func (m *OrderOrder) TableName() string {
	return "order_order"
}

// OrderOrderSearch  查询
type OrderOrderSearch struct {
	OrderOrder
	PageInfo
}
