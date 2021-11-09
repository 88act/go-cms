package request

import (
	"github.com/88act/go-cms/server/model/business"
	"github.com/88act/go-cms/server/model/common/request"
)

type {{.StructName}}Search struct{
    business.{{.StructName}}
    request.PageInfo
}