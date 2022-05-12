package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type ActFavSearch struct{
    business.ActFav
    request.PageInfo
}