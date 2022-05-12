package system

import "go-cms/service"

type ApiGroup struct {
	SystemApiApi
	AuthorityApi
	BaseApi
	CasbinApi
	DictionaryApi
	DictionaryDetailApi
	SystemApi
	DBApi
	JwtApi
	OperationRecordApi
	AuthorityMenuApi
}

var authorityService = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
var apiService = service.ServiceGroupApp.SystemServiceGroup.ApiService
var menuService = service.ServiceGroupApp.SystemServiceGroup.MenuService
var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService
var dictionaryService = service.ServiceGroupApp.SystemServiceGroup.DictionaryService
var dictionaryDetailService = service.ServiceGroupApp.SystemServiceGroup.DictionaryDetailService
var initDBService = service.ServiceGroupApp.SystemServiceGroup.InitDBService
var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService
var baseMenuService = service.ServiceGroupApp.SystemServiceGroup.BaseMenuService
var operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
var userService = service.ServiceGroupApp.SystemServiceGroup.UserService
var systemConfigService = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
