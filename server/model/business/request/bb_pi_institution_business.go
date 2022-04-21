package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type BbPiInstitutionBusinessSearch struct{
    business.BbPiInstitutionBusiness
    request.PageInfo
}