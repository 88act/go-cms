package global

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	Id        int64          `gorm:"primarykey" form:"id" json:"id"`                         // 主键ID
	CreatedAt time.Time      `gorm:"<-:create" form:"createdAt" json:"createdAt,omitempty" ` // 创建时间
	UpdatedAt time.Time      `gorm:"<-:update" form:"createdAt" json:"updatedAt,omitempty" ` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                                         // 删除时间
	//MapData     map[string]string `gorm:"-" sql:"-" json:"mapData,omitempty" `   // key-val 模式数据
	FileObjList []FileObj `json:"fileObjList" cn:"文件列表" gorm:"-" sql:" -"`
}

// 文件简单类型 add by ljd
type FileObj struct {
	Path string `json:"path"`
	Guid string `json:"guid"`
	//MediaType int    `json:"mediaType"`
	//Ext       string `json:"ext"`
}

// 报表数据类型
type ReportData struct {
	Day     string `json:"day"`
	Succeed int    `json:"succeed"`
	Fail    int    `json:"fail"`
	//MediaType int    `json:"mediaType"`
	//Ext       string `json:"ext"`
}

// var sizeOfMyStruct = int(unsafe.Sizeof(GO_MODEL{}))

// func (m *GO_MODEL) ToBytes() []byte {
// 	var x reflect.SliceHeader
// 	x.Len = sizeOfMyStruct
// 	x.Cap = sizeOfMyStruct
// 	x.Data = uintptr(unsafe.Pointer(m))
// 	return *(*[]byte)(unsafe.Pointer(&x))
// }
