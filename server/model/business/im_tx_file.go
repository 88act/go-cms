// 自动生成模板ImTxFile
package business

import (
	"go-cms/global"
)

// ImTxFile 结构体
// 如果含有time.Time 请自行import time包
type ImTxFile struct {
	ImTxFileMini
}

// ImTxFileMini 结构体
type ImTxFileMini struct {
	global.GO_MODEL
	MsgId     *int   `json:"msgId" cn:"消息id" form:"msgId" gorm:"column:msg_id;comment:消息id;type:bigint"`
	MediaType *int   `json:"mediaType" cn:"文件类型" form:"mediaType" gorm:"column:media_type;comment:文件类型;type:smallint"`
	Url       string `json:"url" cn:"远程地址" form:"url" gorm:"column:url;comment:远程地址;type:varchar(255);"`
	Status    *int   `json:"status" cn:"下载状态" form:"status" gorm:"column:status;comment:下载状态;type:smallint"`
	Local     string `json:"local" cn:"本地路径" form:"local" gorm:"column:local;comment:本地路径;type:varchar(255);"`
}

// TableName ImTxFile 表名
func (ImTxFile) TableName() string {
	return "im_tx_file"
}
