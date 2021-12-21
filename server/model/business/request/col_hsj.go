package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type ColHsjSearch struct {
	business.ColHsj
	request.PageInfo
}
