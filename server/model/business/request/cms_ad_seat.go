package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type CmsAdSeatSearch struct{
    business.CmsAdSeat
    request.PageInfo
}