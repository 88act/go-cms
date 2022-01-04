// 自动生成模板SysSuperBuilderHistories
package business

import (
	"go-cms/global"
)

// SysSuperBuilderHistories 结构体
// 如果含有time.Time 请自行import time包
type SysSuperBuilderHistories struct {
	SysSuperBuilderHistoriesMini
	RequestMeta string `json:"requestMeta" cn:"表单" form:"requestMeta" gorm:"column:request_meta;comment:表单;type:text;"`
}

// SysSuperBuilderHistoriesMini 结构体
type SysSuperBuilderHistoriesMini struct {
	global.GO_MODEL
	TableName        string `json:"tableName" cn:"表名" form:"tableName" gorm:"column:table_name;comment:表名;type:varchar(191);"`
	SuperBuilderPath string `json:"superBuilderPath" cn:"superBuilderPath字段" form:"superBuilderPath" gorm:"column:super_builder_path;comment:;type:text;"`
	InjectionMeta    string `json:"injectionMeta" cn:"injectionMeta字段" form:"injectionMeta" gorm:"column:injection_meta;comment:;type:text;"`
	StructName       string `json:"structName" cn:"struct名称" form:"structName" gorm:"column:struct_name;comment:struct名称;type:varchar(191);"`
	StructCnName     string `json:"structCnName" cn:"struct名称" form:"structCnName" gorm:"column:struct_cn_name;comment:struct名称;type:varchar(191);"`
	ApiIds           string `json:"apiIds" cn:"apiIds字段" form:"apiIds" gorm:"column:api_ids;comment:;type:varchar(191);"`
	Flag             *int   `json:"flag" cn:"flag字段" form:"flag" gorm:"column:flag;comment:;type:bigint"`
}

// TableName SysSuperBuilderHistories 表名
// func (SysSuperBuilderHistories) TableName() string {
// 	return "sys_super_builder_histories"
// }
