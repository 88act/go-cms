package service

import (
	"go-cms/service/business"
	"go-cms/service/common"
	"go-cms/service/example"
	"go-cms/service/system"
)

type ServiceGroup struct {
	ExampleServiceGroup  example.ServiceGroup
	SystemServiceGroup   system.ServiceGroup
	BusinessServiceGroup business.ServiceGroup
	CommonServiceGroup   common.ServiceGroup // add by ljd 新增通用 service
}

var ServiceGroupApp = new(ServiceGroup)
