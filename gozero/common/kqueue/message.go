//KqMessage
package kqueue

//第三方支付回调更改支付状态通知
type ThirdPaymentUpdatePayStatusNotifyMessage struct {
	PayStatus int32  `json:"payStatus"`
	OrderSn   string `json:"orderSn"`
}

//短信kf消息
type SmsMessage struct {
	Mobile   string `json:"mobile"`   // 电话
	TmpId    string `json:"tmpId"`    // 模板id
	Title    string `json:"title"`    // 标题
	BodyJson string `json:"bodyJson"` // 详细json
}

//邮件kf消息
type EmailMessage struct {
	Email    string `json:"mobile"`   // 邮件
	TmpId    string `json:"tmpId"`    // 模板id
	Title    string `json:"title"`    // 标题
	Content  string `json:"content"`  // 内容
	BodyJson string `json:"bodyJson"` // 详细json
}

// 消息
type ImageZipMessage struct {
	Path string `json:"path"` // 路径
	Size string `json:"size"` // 尺寸
}
