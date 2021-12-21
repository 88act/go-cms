package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type ImTximSearch struct{
    business.ImTxim
    request.PageInfo
}