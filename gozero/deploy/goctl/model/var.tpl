
var (
	{{.lowerStartCamelObject}}FieldNames          = builder.RawFieldNames(&{{.upperStartCamelObject}}{}{{if .postgreSql}},true{{end}})
	{{.lowerStartCamelObject}}Rows                = strings.Join({{.lowerStartCamelObject}}FieldNames, ",")
	{{.lowerStartCamelObject}}RowsExpectAutoSet   = {{if .postgreSql}}strings.Join(stringx.Remove({{.lowerStartCamelObject}}FieldNames, {{if .autoIncrement}}"{{.originalPrimaryKey}}",{{end}} "created_at", "updated_at"), ","){{else}}strings.Join(stringx.Remove({{.lowerStartCamelObject}}FieldNames, {{if .autoIncrement}}"{{.originalPrimaryKey}}",{{end}} "`created_at`", "`updated_at`","`deleted_at`"), ","){{end}}
	{{.lowerStartCamelObject}}RowsWithPlaceHolder = {{if .postgreSql}}builder.PostgreSqlJoin(stringx.Remove({{.lowerStartCamelObject}}FieldNames, "{{.originalPrimaryKey}}", "created_at", "updated_at")){{else}}strings.Join(stringx.Remove({{.lowerStartCamelObject}}FieldNames, "{{.originalPrimaryKey}}", "`created_at`", "`updated_at`","`deleted_at`"), "=?,") + "=?"{{end}}

	{{if .withCache}}{{.cacheKeys}}{{end}}
)
