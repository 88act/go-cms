package service

import (
	"github.com/88act/go-cms/server/service/business"
	"github.com/88act/go-cms/server/service/common"
	 
	"github.com/88act/go-cms/server/service/system"
)

type ServiceGroup struct {
 
	SystemServiceGroup   system.ServiceGroup
	BusinessServiceGroup business.ServiceGroup
	CommonServiceGroup   common.ServiceGroup // add by ljd 新增通用 service
}

var ServiceGroupApp = new(ServiceGroup)
