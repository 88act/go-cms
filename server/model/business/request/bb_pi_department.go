package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type BbPiDepartmentSearch struct{
    business.BbPiDepartment
    request.PageInfo
}