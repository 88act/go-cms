// 自动生成模板BasicFile
package business

import (
	"github.com/88act/go-cms/server/global"
)

// BasicFile 结构体
// 如果含有time.Time 请自行import time包
type BasicFile struct {
	global.GO_MODEL
	Guid      string `json:"guid" form:"guid" gorm:"column:guid;comment:唯一id;type:char;"`
	UserId    *int   `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;type:bigint"`
	Name      string `json:"name" form:"name" gorm:"column:name;comment:文件名;type:varchar(255);"`
	Module    *int   `json:"module" form:"module" gorm:"column:module;comment:模块名;type:smallint"`
	MediaType *int   `json:"mediaType" form:"mediaType" gorm:"column:media_type;comment:媒体类型;type:smallint"`
	Driver    *int   `json:"driver" form:"driver" gorm:"column:driver;comment:存储位置:0本地1阿里云2腾讯云;type:smallint"`
	Path      string `json:"path" form:"path" gorm:"column:path;comment:文件路径;type:varchar(255);"`
	Ext       string `json:"ext" form:"ext" gorm:"column:ext;comment:文件类型;type:char;"`
	Size      *int   `json:"size" form:"size" gorm:"column:size;comment:文件大小[k];type:int"`
	Md5       string `json:"md5" form:"md5" gorm:"column:md5;comment:md5值;type:char;"`
	Sha1      string `json:"sha1" form:"sha1" gorm:"column:sha1;comment:sha散列值;type:char;"`
	Sort      *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;type:int"`
	Download  *int   `json:"download" form:"download" gorm:"column:download;comment:下载次数;type:int"`
	UsedTime  *int   `json:"usedTime" form:"usedTime" gorm:"column:used_time;comment:使用次数;type:int"`
}

// TableName BasicFile 表名
func (BasicFile) TableName() string {
	return "basic_file"
}
