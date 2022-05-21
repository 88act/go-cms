// 自动生成模板BasicFile
package model

// BasicFile 结构体
type BasicFile struct {
	BaseModel
	Guid      string `db:"guid" json:"guid" gorm:"column:guid"`                  //唯一id
	UserId    *int64 `db:"user_id" json:"userId" gorm:"column:user_id"`          //用户id
	Name      string `db:"name" json:"name" gorm:"column:name"`                  //文件名
	Module    *int   `db:"module" json:"module" gorm:"column:module"`            //模块名
	MediaType *int   `db:"media_type" json:"mediaType" gorm:"column:media_type"` //媒体类型
	Driver    *int   `db:"driver" json:"driver" gorm:"column:driver"`            //存储位置
	Path      string `db:"path" json:"path" gorm:"column:path"`                  //文件路径
	Ext       string `db:"ext" json:"ext" gorm:"column:ext"`                     //文件类型
	Size      *int   `db:"size" json:"size" gorm:"column:size"`                  //文件大小[k]
	Md5       string `db:"md5" json:"md5" gorm:"column:md5"`                     //md5值
	Sha1      string `db:"sha1" json:"sha1" gorm:"column:sha1"`                  //sha散列值
	Sort      *int   `db:"sort" json:"sort" gorm:"column:sort"`                  //排序
	Download  *int   `db:"download" json:"download" gorm:"column:download"`      //下载次数
	UsedTime  *int   `db:"used_time" json:"usedTime" gorm:"column:used_time"`    //使用次数

}

// TableName BasicFile 表名
func (BasicFile) TableName() string {
	return "basic_file"
}

// BasicFileSearch  查询
type BasicFileSearch struct {
	BasicFile
	PageInfo
}
