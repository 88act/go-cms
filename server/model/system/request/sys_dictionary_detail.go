package request

import (
	"go-cms/model/common/request"
	"go-cms/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
