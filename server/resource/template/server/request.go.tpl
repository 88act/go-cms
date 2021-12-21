package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type {{.StructName}}Search struct{
    business.{{.StructName}}
    request.PageInfo
}