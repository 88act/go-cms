package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type BbPiServicePointSearch struct{
    business.BbPiServicePoint
    request.PageInfo
}