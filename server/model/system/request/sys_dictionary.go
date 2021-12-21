package request

import (
	"go-cms/model/common/request"
	"go-cms/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
