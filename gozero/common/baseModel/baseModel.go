// 自动生成模板MemUser
package baseModel

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	Id          int64             `gorm:"primarykey" json:"id"`                // 主键ID
	CreatedAt   time.Time         `json:"createdAt,omitempty"`                 // 创建时间
	UpdatedAt   *time.Time        `json:"updatedAt,omitempty"`                 // 更新时间
	DeletedAt   gorm.DeletedAt    `gorm:"index" json:"-"`                      // 删除时间
	MapData     map[string]string `gorm:"-" sql:"-" json:"mapData,omitempty" ` // key-val 模式数据
	FileObjList []FileObj         `json:"fileObjList" cn:"文件列表" gorm:"-" sql:"-"`
}

// 文件简单类型 add by ljd
type FileObj struct {
	Path string `json:"path"`
	Guid string `json:"guid"`
	//MediaType int    `json:"mediaType"`
	//Ext       string `json:"ext"`
}

type IdsReq struct {
	Ids []int64 `json:"ids" form:"ids"`
}

type GetById struct {
	Id int64 `json:"id" form:"id"` // 主键ID
}

// PageInfo
type PageInfo struct {
	Page           int        `json:"page" form:"page"`           // 页码
	PageSize       int        `json:"pageSize" form:"pageSize"`   // 每页大小
	OrderKey       string     `json:"orderKey" form:"orderKey"`   // 排序
	OrderDesc      bool       `json:"orderDesc" form:"orderDesc"` // 排序方式:升序false(默认)|降序true
	CreatedAtBegin *time.Time `json:"createdAtBegin" form:"createdAtBegin"`
	CreatedAtEnd   *time.Time `json:"createdAtEnd" form:"createdAtEnd"`
	//CreatedAtBetween []string   `json:"createdAtBetween" form:"createdAtBetween"` // 准备删除字段
}

// 快速编辑
type QuickEdit struct {
	Id    int64  `json:"id" form:"id"`       //
	Table string `json:"table" form:"table"` //
	Field string `json:"field" form:"field"` //
	Value string `json:"value" form:"value"` //
}

var ErrNotFound = gorm.ErrRecordNotFound // sqlx.ErrNotFound
