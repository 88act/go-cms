package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type MemUserLogSearch struct {
	business.MemUserLog
	request.PageInfo
}
