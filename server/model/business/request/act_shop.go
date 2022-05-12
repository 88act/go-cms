package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type ActShopSearch struct{
    business.ActShop
    request.PageInfo
}