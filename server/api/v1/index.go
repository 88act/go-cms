package v1

import (
	"go-cms/api/v1/business"
	"go-cms/api/v1/client"
	"go-cms/api/v1/common"
	"go-cms/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	BusinessApiGroup business.ApiGroup
	CommonApiGroup   common.ApiGroup
	ClientGroup      client.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
