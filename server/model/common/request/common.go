package request

import (
	"go-cms/global"
)

// Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
	//Add By ljd 20210708, 统一在这里增加 OrderKey OrderDesc CreatedAtBegin  CreatedAtEnd
	OrderKey       string  `json:"orderKey" form:"orderKey"`   // 排序
	OrderDesc      bool    `json:"orderDesc" form:"orderDesc"` // 排序方式:升序false(默认)|降序true
	Where          *string `json:"where" form:"where"`         // 过滤条件
	CreatedAtBegin *string `json:"createdAtBegin" form:"createdAtBegin"`
	CreatedAtEnd   *string `json:"createdAtEnd" form:"createdAtEnd"`
	//CreatedAtBetween [2]string `json:"createdAtBetween" form:"createdAtBetween"` //查询创建时间 by ljd 20210801
	//CreatedAtBetween make( string,2) `json:"createdAtBetween" form:"createdAtBetween"` //查询创建时间 by ljd 20210801
}

// 快速编辑  新增 BY ljd 20210811------------------------------------
type QuickEdit struct {
	Id    int64  `json:"id" form:"id"`       //
	Table string `json:"table" form:"table"` //
	Field string `json:"field" form:"field"` //
	Value string `json:"value" form:"value"` //
}

// 获取 pid 对应的data   新增 BY ljd 20220615
type PidDataReq struct {
	//CuId      *int64 `json:"cuId" form:"cuId"`           // 客户Id
	Table     string `json:"table" form:"table"`         // 表
	PidField  string `json:"pidField" form:"pidField"`   // value 字段名 (一般是id )
	PidValue  *int64 `json:"pidValue" form:"pidValue"`   // PidValue 字段的值
	NameField string `json:"nameField" form:"nameField"` //label 字段名
	Pid       string `json:"pid" form:"pid"`             // pid 字段名 (一般是pid )
	Where     string `json:"where" form:"where"`         // where 字段名
	Level     int    `json:"level" form:"level"`         // level ,树状数据的最多层级，为0表无限级
}

type PidData struct {
	Value int64  `json:"value" form:"value"`
	Label string `json:"label" form:"label"`
	//Type  int    `json:"type" form:"type"`
	//Children *[]PidData `gorm:"-" sql:"-" json:"children,omitempty" `
}
type PidData2 struct {
	PidData
	Type   int    `json:"type" form:"type"`
	Module string `json:"module" form:"module"`
	//Children *[]PidData `gorm:"-" sql:"-" json:"children,omitempty" `
}

type PidTreeData struct {
	Value    int64          `json:"value" form:"value" gorm:"column:value"`
	Pid      int64          `json:"-" form:"pid" gorm:"column:pid"`
	Label    string         `json:"label" form:"label" gorm:"column:label"`
	Type     int            `json:"type,omitempty" form:"type"`
	Children []*PidTreeData `gorm:"-" sql:"-" json:"children,omitempty" `
}

// 导出 cityArea 时 使用
//
//	type PidTreeData struct {
//		Value    string         `json:"value" form:"value" gorm:"column:value"`
//		Id       int64          `json:"id" form:"id" gorm:"column:id"`
//		Pid      int64          `json:"-" form:"pid" gorm:"column:pid"`
//		Label    string         `json:"label" form:"label" gorm:"column:label"`
//		Text     string         `json:"text" form:"label" gorm:"column:label"`
//		Type     int            `json:"-" form:"type"`
//		Children []*PidTreeData `gorm:"-" sql:"-" json:"children,omitempty" `
//	}
type PidDataResp struct {
	List []PidData
}
type PidDataResp2 struct {
	List []PidData2
}
type DictDataResp struct {
	List []DictData
}

type PidDataTreeResp struct {
	List []*PidTreeData
}

// 获取 pid 对应的data   新增 BY ljd 20220615
type GetDictReq struct {
	Pid        *int64 `json:"pid" form:"pid"`               //pid
	Type       *int   `json:"type" form:"type"`             //type
	Module     string `json:"module" form:"module"`         //模块
	ValueField string `json:"valueField" form:"valueField"` // 选择作为 value 的字段 ,通常是 id 或 value 字段 ,不同情况选择不同的值
}
type DictData struct {
	Value *string `json:"value" form:"value"`
	Label *string `json:"label" form:"label"`
	//Children *[]PidData `gorm:"-" sql:"-" json:"children,omitempty" `
}

// 更新树状结构的pid和sort
type UpdatePidSortReq struct {
	Table    string `json:"table" form:"table"` //表
	Id       int64  `json:"id" form:"id"`       //主动的 node
	Pid      int64  `json:"pid" form:"pid"`
	Sort     int    `json:"sort" form:"sort"`
	Id2      int64  `json:"id2" form:"id2"` //被动的 node
	Pid2     int64  `json:"pid2" form:"pid2"`
	Sort2    int    `json:"sort2" form:"sort2"`
	DropType string `json:"dropType" form:"dropType"`
}

// -------------------------------------------------
type FileUpload struct {
	global.BaseModel
	Name string `json:"name" gorm:"comment:文件名"` // 文件名
	Url  string `json:"url" gorm:"comment:文件地址"` // 文件地址
	Tag  string `json:"tag" gorm:"comment:文件标签"` // 文件标签
	Key  string `json:"key" gorm:"comment:编号"`   // 编号
}

// Find by id structure
type IdReq struct {
	Id int64 `json:"id" form:"id"` // 主键ID
}

// Find by id structure
type IdValueReq struct {
	Id    int64 `json:"id" form:"id"`       // 主键ID
	Value int64 `json:"value" form:"value"` // Value
}

type IdsReq struct {
	Ids []int64 `json:"ids" form:"ids"`
}

// Get role by id structure
type GetAuthorityId struct {
	AuthorityId string `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct{}
