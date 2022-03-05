package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type MemUserSearch struct {
	business.MemUser
	request.PageInfo
}
