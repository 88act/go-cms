package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type BasicAreaSearch struct {
	business.BasicArea
	request.PageInfo
}
