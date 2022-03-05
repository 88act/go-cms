package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type MemUserSafeSearch struct {
	business.MemUserSafe
	request.PageInfo
}
