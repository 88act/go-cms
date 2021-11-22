package v1

import (
	"github.com/88act/go-cms/server/api/v1/business"
	"github.com/88act/go-cms/server/api/v1/common"
	 
	"github.com/88act/go-cms/server/api/v1/system"
)

type ApiGroup struct {
 
	SystemApiGroup   system.ApiGroup
	BusinessApiGroup business.ApiGroup
	CommonApiGroup   common.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
