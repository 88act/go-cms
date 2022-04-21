package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type BbPiDeviceSearch struct{
    business.BbPiDevice
    request.PageInfo
}