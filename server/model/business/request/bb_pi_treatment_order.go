package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type BbPiTreatmentOrderSearch struct{
    business.BbPiTreatmentOrder
    request.PageInfo
}