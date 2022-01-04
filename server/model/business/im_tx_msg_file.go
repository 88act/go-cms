// 自动生成模板ImTxMsgFile
package business

import (
	"go-cms/global"
	"time"
)

// ImTxMsgFile 结构体
// 如果含有time.Time 请自行import time包
type ImTxMsgFile struct {
	ImTxMsgFileMini
	Url       string `json:"url" cn:"下载路径" form:"url" gorm:"column:url;comment:下载路径;type:varchar(256);"`
	ErrorInfo string `json:"errorInfo" cn:"请求信息" form:"errorInfo" gorm:"column:error_info;comment:请求信息;type:varchar(255);"`
	LocalFile string `json:"localFile" cn:"本地路径" form:"localFile" gorm:"column:local_file;comment:本地路径;type:varchar(255);"`
}

// ImTxMsgFileMini 结构体
type ImTxMsgFileMini struct {
	global.GO_MODEL
	ChatType     string     `json:"chatType" cn:"消息类型" form:"chatType" gorm:"column:chat_type;comment:消息类型;type:varchar(100);"`
	MsgTime      string     `json:"msgTime" cn:"消息时间" form:"msgTime" gorm:"column:msg_time;comment:消息时间;type:varchar(100);"`
	ExpireTime   *time.Time `json:"expireTime" cn:"过期时间" form:"expireTime" gorm:"column:expire_time;comment:过期时间;type:datetime"`
	FileSize     *int       `json:"fileSize" cn:"文件大小" form:"fileSize" gorm:"column:file_size;comment:文件大小;type:int"`
	FileMd5      string     `json:"fileMd5" cn:"文件md5" form:"fileMd5" gorm:"column:file_md5;comment:文件md5"`
	GzipSize     *int       `json:"gzipSize" cn:"压缩大小" form:"gzipSize" gorm:"column:gzip_size;comment:压缩大小;type:int"`
	GzipMd5      string     `json:"gzipMd5" cn:"压缩md5" form:"gzipMd5" gorm:"column:gzip_md5;comment:压缩md5"`
	ErrorCode    string     `json:"errorCode" cn:"请求code" form:"errorCode" gorm:"column:error_code;comment:请求code;type:varchar(255);"`
	ActionStatus string     `json:"actionStatus" cn:"请求状态" form:"actionStatus" gorm:"column:action_status;comment:请求状态;type:varchar(100);"`
	Status       *int       `json:"status" cn:"状态" form:"status" gorm:"column:status;comment:状态:0未下载 1已下载  2已处理;type:smallint"`
}

// TableName ImTxMsgFile 表名
func (ImTxMsgFile) TableName() string {
	return "im_tx_msg_file"
}
