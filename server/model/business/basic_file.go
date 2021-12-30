// 自动生成模板BasicFile
package business

import (
	"go-cms/global"
)

// BasicFile 结构体
// 如果含有time.Time 请自行import time包
type BasicFile struct {
	BasicFileMini
	Md5      string `json:"md5" cn:"md5值" form:"md5" gorm:"column:md5;comment:md5值;type:char;"`
	Sha1     string `json:"sha1" cn:"sha散列值" form:"sha1" gorm:"column:sha1;comment:sha散列值;type:char;"`
	Download *int   `json:"download" cn:"下载次数" form:"download" gorm:"column:download;comment:下载次数;type:int"`
	UsedTime *int   `json:"usedTime" cn:"使用次数" form:"usedTime" gorm:"column:used_time;comment:使用次数;type:int"`
}

// BasicFileMini 结构体
type BasicFileMini struct {
	global.GO_MODEL
	Path      string `json:"path" cn:"文件路径" form:"path" gorm:"column:path;comment:文件路径;type:varchar(255);"`
	Guid      string `json:"guid" cn:"唯一id" form:"guid" gorm:"column:guid;comment:唯一id;type:char;"`
	UserId    *int   `json:"userId" cn:"用户id" form:"userId" gorm:"column:user_id;comment:用户id;type:bigint"`
	Name      string `json:"name" cn:"文件名" form:"name" gorm:"column:name;comment:文件名;type:varchar(255);"`
	Module    *int   `json:"module" cn:"模块名" form:"module" gorm:"column:module;comment:模块名;type:smallint"`
	MediaType *int   `json:"mediaType" cn:"媒体类型" form:"mediaType" gorm:"column:media_type;comment:媒体类型;type:smallint"`
	Driver    *int   `json:"driver" cn:"存储位置" form:"driver" gorm:"column:driver;comment:存储位置:0本地1阿里云2腾讯云;type:smallint"`
	Ext       string `json:"ext" cn:"文件类型" form:"ext" gorm:"column:ext;comment:文件类型;type:char;"`
	Size      *int   `json:"size" cn:"文件大小[k]" form:"size" gorm:"column:size;comment:文件大小[k];type:int"`
	Sort      *int   `json:"sort" cn:"排序" form:"sort" gorm:"column:sort;comment:排序;type:int"`
}

// TableName BasicFile 表名
func (BasicFile) TableName() string {
	return "basic_file"
}
