package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type ImTxMsgSearch struct {
	business.ImTxMsg
	request.PageInfo
}
