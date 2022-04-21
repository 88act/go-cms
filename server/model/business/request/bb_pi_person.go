package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type BbPiPersonSearch struct{
    business.BbPiPerson
    request.PageInfo
}