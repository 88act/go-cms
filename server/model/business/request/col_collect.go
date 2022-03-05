package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type ColCollectSearch struct {
	business.ColCollect
	request.PageInfo
}
