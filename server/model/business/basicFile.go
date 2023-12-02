// 模板BasicFile
// Author 10512203@qq.com
package business

import (
	"go-cms/global"
	"go-cms/model/common/request"
)

// 表格字段
const Field_BasicFile_mini = "" // "id,created_at,updated_at,guid,user_id,user_id_sys,name,cat_id,module,media_type,driver,path,file_type,ext,size,md5,sha1,sort,download,used_time,status,

// BasicFile 结构体
type BasicFile struct {
	global.BaseModel
	Guid      string `json:"guid" form:"guid" cn:"唯一id"  gorm:"column:guid;comment:唯一id;type:varchar(32);"`
	UserId    int64  `json:"userId" form:"userId" cn:"用户id"  gorm:"column:user_id;comment:用户id;type:bigint"`
	Name      string `json:"name" form:"name" cn:"文件名"  gorm:"column:name;comment:文件名;type:varchar(255);"`
	CatId     int64  `json:"catId" form:"catId" cn:"栏目id"  gorm:"column:cat_id;comment:栏目id;type:bigint"`
	Module    int    `json:"module" form:"module" cn:"模块id"   gorm:"column:module;comment:模块id;type:smallint"`
	MediaType int    `json:"mediaType" form:"mediaType" cn:"媒体类型"   gorm:"column:media_type;comment:媒体类型: 1 图片 2 视频 3音频 4文档;type:smallint"`
	Driver    int    `json:"driver" form:"driver" cn:"存储位置"   gorm:"column:driver;comment:存储位置:0本地1阿里云2腾讯云;type:smallint"`
	Path      string `json:"path" form:"path" cn:"文件路径"  gorm:"column:path;comment:文件路径;type:varchar(255);"`
	FileType  string `json:"fileType" form:"fileType" cn:"文件类型"  gorm:"column:file_type;comment:文件类型;type:varchar(50);"`
	Ext       string `json:"ext" form:"ext" cn:"文件类型"  gorm:"column:ext;comment:文件类型;type:varchar(8);"`
	Size      int    `json:"size" form:"size" cn:"文件大小[k]"   gorm:"column:size;comment:文件大小[k];type:int"`
	Md5       string `json:"md5" form:"md5" cn:"md5值"  gorm:"column:md5;comment:md5值;type:varchar(50);"`
	Sha1      string `json:"sha1" form:"sha1" cn:"sha散列值"  gorm:"column:sha1;comment:sha散列值;type:varchar(50);"`
	Sort      int    `json:"sort" form:"sort" cn:"排序"   gorm:"column:sort;comment:排序;type:int"`
	Download  int    `json:"download" form:"download" cn:"下载次数"   gorm:"column:download;comment:下载次数;type:int"`
	UsedTime  int    `json:"usedTime" form:"usedTime" cn:"使用次数"   gorm:"column:used_time;comment:使用次数;type:int"`
	Status    int    `json:"status" form:"status" cn:"status字段"   gorm:"column:status;comment:;type:smallint"`
}

// TableName BasicFile 表名
func (BasicFile) TableName() string {
	return "basic_file"
}

type BasicFileSearch struct {
	request.PageInfo
	global.BaseModel
	Guid      string `json:"guid"   form:"guid" `
	UserId    *int64 `json:"userId"  form:"userId" `
	Name      string `json:"name"   form:"name" `
	CatId     *int64 `json:"catId"  form:"catId" `
	Module    *int   `json:"module"   form:"module" `
	MediaType *int   `json:"mediaType"   form:"mediaType" `
	Driver    *int   `json:"driver"   form:"driver" `
	Path      string `json:"path"   form:"path" `
	FileType  string `json:"fileType"   form:"fileType" `
	Ext       string `json:"ext"   form:"ext" `
	Size      *int   `json:"size"   form:"size" `
	Md5       string `json:"md5"   form:"md5" `
	Sha1      string `json:"sha1"   form:"sha1" `
	Sort      *int   `json:"sort"   form:"sort" `
	Download  *int   `json:"download"   form:"download" `
	UsedTime  *int   `json:"usedTime"   form:"usedTime" `
	Status    *int   `json:"status"   form:"status" `
}

type BasicFileMini struct {
	Guid string `json:"guid" cn:"唯一id" form:"guid" gorm:"column:guid;comment:唯一id;type:char;"`
	Name string `json:"name" cn:"文件名" form:"name" gorm:"column:name;comment:文件名;type:varchar(255);"`
	Path string `json:"path" cn:"文件路径" form:"path" gorm:"column:path;comment:文件路径;type:varchar(255);"`
}
