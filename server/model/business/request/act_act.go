package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type ActActSearch struct{
    business.ActAct
    request.PageInfo
}