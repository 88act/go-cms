package system

import "go-cms/service"

type ApiGroup struct {
	BaseUserApi
	BaseMenuApi
	SystemApi
}

var userService = service.ServiceGroupApp.SystemServiceGroup.UserService
var systemConfigService = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
