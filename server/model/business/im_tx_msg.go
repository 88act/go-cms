// 自动生成模板ImTxMsg
package business

import (
	"go-cms/global"
	"time"
)

// ImTxMsg 结构体
// 如果含有time.Time 请自行import time包
type ImTxMsg struct {
	ImTxMsgMini
	MsgSeq          *int   `json:"msgSeq" cn:"seq" form:"msgSeq" gorm:"column:msg_seq;comment:seq;type:int"`
	MsgRandom       *int   `json:"msgRandom" cn:"random" form:"msgRandom" gorm:"column:msg_random;comment:random"`
	MsgContent      string `json:"msgContent" cn:"内容" form:"msgContent" gorm:"column:msg_content;comment:内容;type:text;"`
	MediaList       string `json:"mediaList" cn:"媒体文件" form:"mediaList" gorm:"column:media_list;comment:媒体文件;type:varchar(500);"`
	MediaListTx     string `json:"mediaListTx" cn:"媒体远程文件" form:"mediaListTx" gorm:"column:media_list_tx;comment:媒体远程文件;type:varchar(500);"`
	ClientIp        string `json:"clientIp" cn:"IP地址" form:"clientIp" gorm:"column:client_ip;comment:IP地址;type:varchar(255);"`
	MsgFromPlatform string `json:"msgFromPlatform" cn:"平台" form:"msgFromPlatform" gorm:"column:msg_from_platform;comment:平台;type:varchar(255);"`
}

// ImTxMsgMini 结构体
type ImTxMsgMini struct {
	global.GO_MODEL
	OrderId      string     `json:"orderId" cn:"订单id" form:"orderId" gorm:"column:order_id;comment:订单id;type:varchar(50);"`
	ChatType     string     `json:"chatType" cn:"消息类型" form:"chatType" gorm:"column:chat_type;comment:消息类型;type:varchar(100);"`
	MsgTime      string     `json:"msgTime" cn:"消息时间" form:"msgTime" gorm:"column:msg_time;comment:消息时间;type:varchar(100);"`
	FromAccount  string     `json:"fromAccount" cn:"发送人" form:"fromAccount" gorm:"column:from_account;comment:发送人;type:varchar(100);"`
	ToAccount    string     `json:"toAccount" cn:"接收人" form:"toAccount" gorm:"column:to_account;comment:接收人;type:varchar(100);"`
	MsgTimestamp *time.Time `json:"msgTimestamp" cn:"时间撮" form:"msgTimestamp" gorm:"column:msg_timestamp;comment:时间撮;type:datetime"`
	MsgType      string     `json:"msgType" cn:"消息类型" form:"msgType" gorm:"column:msg_type;comment:消息类型;type:varchar(100);"`
	MsgText      string     `json:"msgText" cn:"文本" form:"msgText" gorm:"column:msg_text;comment:文本;type:varchar(4000);"`
	StatusMedia  *int       `json:"statusMedia" cn:"媒体下载" form:"statusMedia" gorm:"column:status_media;comment:媒体下载:0未下载,1已下载,2 失败;type:smallint"`
	Status       *int       `json:"status" cn:"状态" form:"status" gorm:"column:status;comment:状态:0未上传,1已上传,2 失败;type:smallint"`
	OrderStatus  *int       `json:"order_status" cn:"订单状态" form:"order_status" gorm:"column:order_status;comment:订单状态;type:smallint"`
}

// TableName ImTxMsg 表名
func (ImTxMsg) TableName() string {
	return "im_tx_msg"
}
