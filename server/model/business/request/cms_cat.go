package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type CmsCatSearch struct{
    business.CmsCat
    request.PageInfo
}