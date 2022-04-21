package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type BbPiStaffSearch struct{
    business.BbPiStaff
    request.PageInfo
}