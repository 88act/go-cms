package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type BbPiTreatmentReferralSearch struct{
    business.BbPiTreatmentReferral
    request.PageInfo
}