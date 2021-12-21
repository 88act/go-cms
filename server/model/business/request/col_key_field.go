package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type ColKeyFieldSearch struct {
	business.ColKeyField
	request.PageInfo
}
