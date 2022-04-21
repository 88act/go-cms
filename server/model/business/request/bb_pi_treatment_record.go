package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type BbPiTreatmentRecordSearch struct{
    business.BbPiTreatmentRecord
    request.PageInfo
}