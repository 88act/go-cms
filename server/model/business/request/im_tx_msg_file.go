package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type ImTxMsgFileSearch struct {
	business.ImTxMsgFile
	request.PageInfo
}
