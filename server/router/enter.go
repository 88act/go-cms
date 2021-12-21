package router

import (
	"go-cms/router/business"
	common "go-cms/router/common"
	"go-cms/router/example"
	"go-cms/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Business business.RouterGroup
	Common   common.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
