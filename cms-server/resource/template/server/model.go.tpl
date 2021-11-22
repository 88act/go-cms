// 自动生成模板{{.StructName}}
package business

import (
	"github.com/88act/go-cms/server/global"
)

// {{.StructName}} 结构体
// 如果含有time.Time 请自行import time包
type {{.StructName}} struct {
     {{.StructName}}Mini
{{- range .Fields}}
  {{- if eq .BeHide true }}
{{- if eq .FieldType "image" }}
      {{.FieldName}}  string `json:"{{.FieldJson}}" cn:"{{.FieldDesc}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- end }}"`
{{- else if ne .FieldType "string" }}
      {{.FieldName}}  *{{.FieldType}} `json:"{{.FieldJson}}" cn:"{{.FieldDesc}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- end }}"`
{{- else }}
      {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}" cn:"{{.FieldDesc}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- if eq .FieldType "string" -}}{{- if .DataTypeLong -}}({{.DataTypeLong}}){{- end -}}{{- end -}};{{- if ne .FieldType "string" -}}{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}{{- end -}}{{- end -}}"`
{{- end }} 
{{- end }} 
{{- end }}
}

// {{.StructName}}Mini 结构体
type {{.StructName}}Mini struct {
      global.GO_MODEL
     {{- range .Fields}}
    {{- if eq .BeHide false }}
{{- if eq .FieldType "image" }}
      {{.FieldName}}  string `json:"{{.FieldJson}}" cn:"{{.FieldDesc}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- end }}"`
{{- else if ne .FieldType "string" }}
      {{.FieldName}}  *{{.FieldType}} `json:"{{.FieldJson}}" cn:"{{.FieldDesc}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- end }}"`
{{- else }}
      {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}" cn:"{{.FieldDesc}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- if eq .FieldType "string" -}}{{- if .DataTypeLong -}}({{.DataTypeLong}}){{- end -}}{{- end -}};{{- if ne .FieldType "string" -}}{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}{{- end -}}{{- end -}}"`
{{- end }} 
  {{- end }}
   {{- end }}

}

{{ if .TableName }}
// TableName {{.StructName}} 表名
func ({{.StructName}}) TableName() string {
  return "{{.TableName}}"
}
{{ end }}
