package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type BbPiContactsSearch struct{
    business.BbPiContacts
    request.PageInfo
}