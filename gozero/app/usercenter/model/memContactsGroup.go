// 自动生成模板MemContactsGroup
package model

import (
	. "go-cms/common/baseModel"
)

// MemContactsGroup 结构体
type MemContactsGroup struct {
	BaseModel
	UserId int64  `db:"user_id" json:"userId" gorm:"column:user_id"` //用户id
	Name   string `db:"name" json:"name" gorm:"column:name"`         //名称
	Star   int    `db:"star" json:"star" gorm:"column:star"`         //标星
	Status int    `db:"status" json:"status" gorm:"column:status"`   //状态

}

// TableName MemContactsGroup 表名
func (MemContactsGroup) TableName() string {
	return "mem_contacts_group"
}

// MemContactsGroupSearch  查询
type MemContactsGroupSearch struct {
	BaseModel
	PageInfo
	UserId *int64 `json:"userId"  form:"userId" `
	Name   string `json:"name"   form:"name" `
	Star   *int   `json:"star"   form:"star" `
	Status *int   `json:"status"   form:"status" `
}
