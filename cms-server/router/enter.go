package router

import (
	"github.com/88act/go-cms/server/router/business"
	common "github.com/88act/go-cms/server/router/common"
 
	"github.com/88act/go-cms/server/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
 
	Business business.RouterGroup
	Common   common.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
