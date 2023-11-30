package service

import (
	"go-cms/service/business"
	"go-cms/service/common"

	"go-cms/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	BusinessServiceGroup business.ServiceGroup
	CommonServiceGroup   common.ServiceGroup // add by ljd 新增通用 service
}

var ServiceGroupApp = new(ServiceGroup)
