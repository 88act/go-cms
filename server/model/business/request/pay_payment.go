package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type PayPaymentSearch struct{
    business.PayPayment
    request.PageInfo
}