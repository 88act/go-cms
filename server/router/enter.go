package router

import (
	"go-cms/router/business"
	common "go-cms/router/common"

	"go-cms/router/system"
)

type RouterGroup struct {
	System system.RouterGroup

	Business business.RouterGroup
	Common   common.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
