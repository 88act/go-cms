package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type CmsAdSearch struct {
	business.CmsAd
	request.PageInfo
}
