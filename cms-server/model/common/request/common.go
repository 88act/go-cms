package request

import "github.com/88act/go-cms/server/global"

// Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
	//Add By ljd 20210708, 统一在这里增加 OrderKey OrderDesc CreatedAtBegin  CreatedAtEnd
	OrderKey  string `json:"orderKey" form:"orderKey"`   // 排序
	OrderDesc bool   `json:"orderDesc" form:"orderDesc"` // 排序方式:升序false(默认)|降序true
	//CreatedAtBetween [2]string `json:"createdAtBetween" form:"createdAtBetween"` //查询创建时间 by ljd 20210801
	//CreatedAtBetween make( string,2) `json:"createdAtBetween" form:"createdAtBetween"` //查询创建时间 by ljd 20210801
}

//快速编辑  新增 BY ljd 20210811
type QuickEdit struct {
	Id    uint   `json:"id" form:"id"`       //
	Table string `json:"table" form:"table"` //
	Field string `json:"field" form:"field"` //
	Value string `json:"value" form:"value"` //
}

type FileUpload struct {
	global.GO_MODEL
	Name string `json:"name" gorm:"comment:文件名"` // 文件名
	Url  string `json:"url" gorm:"comment:文件地址"` // 文件地址
	Tag  string `json:"tag" gorm:"comment:文件标签"` // 文件标签
	Key  string `json:"key" gorm:"comment:编号"`   // 编号
}

// Find by id structure
type GetById struct {
	ID float64 `json:"id" form:"id"` // 主键ID
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// Get role by id structure
type GetAuthorityId struct {
	AuthorityId string `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct{}
