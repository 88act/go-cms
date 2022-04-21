package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type BbPiTreatmentDiagnoseSearch struct{
    business.BbPiTreatmentDiagnose
    request.PageInfo
}