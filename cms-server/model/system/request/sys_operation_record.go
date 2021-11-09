package request

import (
	"github.com/88act/go-cms/server/model/common/request"
	"github.com/88act/go-cms/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
