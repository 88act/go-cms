 
// 模板{{.StructName}}
// Author 10512203@qq.com
package business

import (
	"go-cms/global"
   	"go-cms/model/common/request"
    "time"
)

//表格字段
const Field_{{.StructName}}_mini = "" // "id,created_at,updated_at,{{- range .Fields}}{{- if eq .BeHide false }}{{.ColumnName}},{{- end }} {{- end }} 

// {{.StructName}} 结构体
type {{.StructName}} struct {
 global.GO_MODEL
    {{- range .Fields}}  
{{- if eq .FieldType "image" }}
   {{.FieldName}}  string `json:"{{.FieldJson}}"  form:"{{.FieldJson}}"  cn:"{{.FieldDesc}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- end }}"`
{{- else if eq .FieldType "json" }}
   {{.FieldName}}  string `json:"{{.FieldJson}}"  form:"{{.FieldJson}}" cn:"{{.FieldDesc}}"  gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- end }}"`
{{- else if eq .FieldType "bool" }}
    {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" cn:"{{.FieldDesc}}"   gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- end }}"`
{{- else if eq .FieldType "time.Time" }}
    {{.FieldName}}  *{{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" cn:"{{.FieldDesc}}"  gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- end }}"`
{{- else if eq .DataType "bigint" }} 
    {{.FieldName}}  int64 `json:"{{.FieldJson}}" form:"{{.FieldJson}}" cn:"{{.FieldDesc}}"  gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- end }}"`
{{- else if eq .FieldType "int" }} 
     {{.FieldName}} int `json:"{{.FieldJson}}" form:"{{.FieldJson}}" cn:"{{.FieldDesc}}"   gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- end }}"`
{{- else if eq .DataType "json" }} 
    {{.FieldName}}  string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" cn:"{{.FieldDesc}}"  gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- end }}"`
{{- else }} 
    {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" cn:"{{.FieldDesc}}"  gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- if eq .FieldType "string" -}}{{- if .DataTypeLong -}}({{.DataTypeLong}}){{- end -}}{{- end -}};{{- if ne .FieldType "string" -}}{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}{{- end -}}{{- end -}}"`
{{- end }}  
   {{- end }} 
}

{{ if .TableName }}
// TableName {{.StructName}} 表名
func ({{.StructName}}) TableName() string {
  return "{{.TableName}}"
}
{{ end }}



type {{.StructName}}Search struct{  
    request.PageInfo
    global.GO_MODEL
     {{- range .Fields}}  
{{- if eq .FieldType "image" }}
   {{.FieldName}}  string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
{{- else if eq .FieldType "bool" }}
    {{.FieldName}}  *{{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
{{- else if eq .FieldType "time.Time" }}
    {{.FieldName}}  *{{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
{{- else if eq .DataType "bigint" }} 
    {{.FieldName}}  *int64 `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
{{- else if eq .FieldType "int" }} 
     {{.FieldName}}  *int `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
{{- else if eq .DataType "json" }} 
    {{.FieldName}}  string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
{{- else }} 
    {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
{{- end }}  
   {{- end }} 
}