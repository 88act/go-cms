package uuid

import (
	"fmt"
	"go-cms/common/tool"
	"time"
)

// 生成sn单号
type SnPrefix string

const (
	SN_PREFIX_HOMESTAY_ORDER SnPrefix = "ACT" // 订单前缀 go-cms_order/homestay_order
	SN_PREFIX_THIRD_PAYMENT  SnPrefix = "THR" //第三方支付流水记录前缀 go-cms_payment/third_payment
)

// 生成单号
func GenSn(snPrefix SnPrefix) string {
	return fmt.Sprintf("%s%s%s", snPrefix, time.Now().Format("20060102150405"), tool.Krand(8, tool.KC_RAND_KIND_NUM))
}
