package router

import (
	"go-cms/router/business"
	"go-cms/router/client"
	common "go-cms/router/common"
	"go-cms/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Business business.RouterGroup
	Common   common.RouterGroup
	Client   client.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
