package request

import (
	"github.com/88act/go-cms/server/model/business"
	"github.com/88act/go-cms/server/model/common/request"
)

type CmsAdSeatSearch struct{
    business.CmsAdSeat
    request.PageInfo
}