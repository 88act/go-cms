package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type BbPiUpFieldSearch struct{
    business.BbPiUpField
    request.PageInfo
}