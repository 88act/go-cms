package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type BbPiInstitutionSearch struct{
    business.BbPiInstitution
    request.PageInfo
}