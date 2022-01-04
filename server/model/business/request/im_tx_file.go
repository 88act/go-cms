package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type ImTxFileSearch struct {
	business.ImTxFile
	request.PageInfo
}
