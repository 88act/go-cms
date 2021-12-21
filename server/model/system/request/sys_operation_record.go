package request

import (
	"go-cms/model/common/request"
	"go-cms/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
