package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type BasicFileSearch struct {
	business.BasicFile
	request.PageInfo
}
