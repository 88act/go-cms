// 自动生成模板{{.StructName}}
package model

import (
	"time"
     . "go-cms/common/baseModel"
)
  // {{.StructName}} 结构体 
type {{.StructName}} struct {
     BaseModel
 {{- range .Fields}} 
{{- if eq .FieldType "image" }}
     {{.FieldName}}  string `db:"{{.ColumnName}}" json:"{{.FieldJson}}" gorm:"column:{{.ColumnName}}"`   //{{.FieldDesc}}
{{- else if eq .DataType "json" }}
   {{.FieldName}}  string `db:"{{.ColumnName}}" json:"{{.FieldJson}}" gorm:"column:{{.ColumnName}}"`   //{{.FieldDesc}}
{{- else if eq .FieldType "string" }}
     {{.FieldName}}  {{.FieldType}} `db:"{{.ColumnName}}" json:"{{.FieldJson}}" gorm:"column:{{.ColumnName}}"`   //{{.FieldDesc}}
{{- else if eq .FieldType "time.Time" }}
     {{.FieldName}}  *{{.FieldType}} `db:"{{.ColumnName}}" json:"{{.FieldJson}}" gorm:"column:{{.ColumnName}}"`  //{{.FieldDesc}}      
{{- else if eq .DataType "bigint" }} 
     {{.FieldName}} int64 `db:"{{.ColumnName}}" json:"{{.FieldJson}}" gorm:"column:{{.ColumnName}}"`  //{{.FieldDesc}}
{{- else if eq .FieldType "int" }} 
     {{.FieldName}} int `db:"{{.ColumnName}}" json:"{{.FieldJson}}" gorm:"column:{{.ColumnName}}"`  //{{.FieldDesc}}
{{- else }}
     {{.FieldName}} {{.FieldType}} `db:"{{.ColumnName}}" json:"{{.FieldJson}}" gorm:"column:{{.ColumnName}}"`  //{{.FieldDesc}}
{{- end }}  
{{- end }}
 
}
 
{{ if .TableName }}
// TableName {{.StructName}} 表名
func ({{.StructName}}) TableName() string {
  return "{{.TableName}}"
}
{{ end }}


// {{.StructName}}Search  查询
type {{.StructName}}Search struct { 
     BaseModel
	PageInfo
    {{- range .Fields}}  
{{- if eq .FieldType "image" }}
   {{.FieldName}}  string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
{{- else if eq .FieldType "bool" }}
    {{.FieldName}}  *{{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
{{- else if eq .FieldType "time.Time" }}
    {{.FieldName}}  *{{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
{{- else if eq .DataType "bigint" }} 
    {{.FieldName}}  *int64 `json:"{{.FieldJson}}"  form:"{{.FieldJson}}" `
{{- else if eq .FieldType "int" }} 
     {{.FieldName}}  *int `json:"{{.FieldJson}}"   form:"{{.FieldJson}}" `
{{- else if eq .DataType "json" }} 
    {{.FieldName}}  string `json:"{{.FieldJson}}"  form:"{{.FieldJson}}" `
{{- else }} 
    {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}"   form:"{{.FieldJson}}" `
{{- end }}  
   {{- end }} 
}
 