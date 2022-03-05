package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type MemAddressSearch struct {
	business.MemAddress
	request.PageInfo
}
