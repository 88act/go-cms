package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type OrderOrderSearch struct{
    business.OrderOrder
    request.PageInfo
}