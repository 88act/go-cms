package request

import (
	"github.com/88act/go-cms/server/model/business"
	"github.com/88act/go-cms/server/model/common/request"
)

type CmsCatSearch struct{
    business.CmsCat
    request.PageInfo
}