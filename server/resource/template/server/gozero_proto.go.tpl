 syntax = "proto3";

option go_package = "./pb";

package pb;
// proto 文件 
message {{.StructName}}  { 
  int64 id  = 1;  
  int64 createdAt  = 1; 
  int64 updatedAt  = 1; 
 {{- range $index, $obj := .Fields}} 
{{- if eq $obj.FieldType "image" }}
  string {{$obj.FieldJson}}   = {{$index}};    //{{.FieldDesc}}
{{- else if eq .FieldType "time.Time" }} 
  int64 {{$obj.FieldJson}}   = {{$index}};   //{{.FieldDesc}}
 {{- else if eq .FieldType "bool" }} 
  bool {{$obj.FieldJson}}   = {{$index}};       //{{.FieldDesc}}
{{- else if eq .DataType "bigint" }} 
  int64 {{$obj.FieldJson}}  = {{$index}};     //{{.FieldDesc}}
{{- else if eq .FieldType "int" }} 
  int32 {{$obj.FieldJson}}  = {{$index}};    //{{.FieldDesc}}
{{- else }}
  {{$obj.FieldType}} {{$obj.FieldJson}}  = {{$index}};    //{{.FieldDesc}} 
{{- end }}  
{{- end }}   
}
 
 // api 文件 
 // {{.StructName}} 结构体 
type (
{{.StructName}}  { 
    Id   int64   `json:"id"`      //  id
    CreatedAt   int64   `json:"createdAt"`      //  创建时间 
    UpdatedAt   int64   `json:"updatedAt"`      //  更新时间
 {{- range .Fields}} 
{{- if eq .FieldType "image" }}
    {{.FieldName}}  string `json:"{{.FieldJson}}"`     //{{.FieldDesc}}
{{- else if eq .FieldType "time.Time" }} 
    {{.FieldName}}  int64 `json:"{{.FieldJson}}"`     //{{.FieldDesc}}
 {{- else if eq .DataType "bigint" }} 
    {{.FieldName}}  int64   `json:"{{.FieldJson}}"`     //{{.FieldDesc}}
{{- else if ne .FieldType "string" }}
    {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}"`   //{{.FieldDesc}}
{{- else }}
    {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}"`    //{{.FieldDesc}}
{{- end }}  
   {{- end }}
   }
)

//service
service  {  
}

  