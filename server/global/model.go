package global

import (
	"time"

	"gorm.io/gorm"
)

type GO_MODEL struct {
	ID        uint              `gorm:"primarykey" json:"ID"`                // 主键ID
	CreatedAt time.Time         `json:"CreatedAt,omitempty"`                 // 创建时间
	UpdatedAt time.Time         `json:"UpdatedAt,omitempty"`                 // 更新时间
	DeletedAt gorm.DeletedAt    `gorm:"index" json:"-"`                      // 删除时间
	MapData   map[string]string `gorm:"-" sql:"-" json:"mapData,omitempty" ` // key-val 模式数据
}
