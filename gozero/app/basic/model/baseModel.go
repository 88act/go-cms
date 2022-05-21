// 自动生成模板
package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	Id        int64             `gorm:"primarykey" json:"id"`                // 主键ID
	CreatedAt time.Time         `json:"CreatedAt,omitempty"`                 // 创建时间
	UpdatedAt time.Time         `json:"UpdatedAt,omitempty"`                 // 更新时间
	DeletedAt gorm.DeletedAt    `gorm:"index" json:"-"`                      // 删除时间
	MapData   map[string]string `gorm:"-" sql:"-" json:"mapData,omitempty" ` // key-val 模式数据
}

type IdsReq struct {
	Ids []int64 `json:"ids" form:"ids"`
}

type GetById struct {
	Id int64 `json:"id" form:"id"` // 主键ID
}

// PageInfo
type PageInfo struct {
	Page             int      `json:"page" form:"page"`            // 页码
	PageSize         int      `json:"pageSize" form:"pageSize"`    // 每页大小
	OrderKey         string   `json:"orderKey" form:"orderKey"`    // 排序
	OrderDesc        bool     `json:"orderDesc" form:"orderDesc"`  // 排序方式:升序false(默认)|降序true
	CreatedAtBetween []string `json:"-" form:"created-at-between"` // 查询 CreatedAt
}

//快速编辑
type QuickEdit struct {
	Id    int64  `json:"id" form:"id"`       //
	Table string `json:"table" form:"table"` //
	Field string `json:"field" form:"field"` //
	Value string `json:"value" form:"value"` //
}
