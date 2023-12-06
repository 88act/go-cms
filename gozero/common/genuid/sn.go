package genuid

import (
	"fmt"
	"go-cms/common/tool"
	"time"
)

// 生成sn单号
type SnPrefix string

const (
	SN_PREFIX_ACT  SnPrefix = "ACT" // ACT 订单前缀
	SN_PREFIX_SELL SnPrefix = "SEL" //Sell 推广
	SN_PREFIX_PAY  SnPrefix = "PAY" //第三方支付流水记录前缀
)

// 生成单号 3 + 14 +8 = 25
func GenSn(snPrefix SnPrefix) string {
	return fmt.Sprintf("%s%s%s%s", snPrefix, time.Now().Format("20060102150405"), tool.Krand(2, tool.KC_RAND_KIND_UPPER), tool.Krand(6, tool.KC_RAND_KIND_ALL))
	//return fmt.Sprintf("%s%s%s", snPrefix, time.Now().Format("20060102150405"), tool.GetRandomStr(8))
}
