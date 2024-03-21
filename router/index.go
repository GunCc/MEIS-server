package router

import (
	"MEIS-server/router/oa"
	"MEIS-server/router/system"
)

type RouterGroup struct {
	System system.SystemRouterGroup
	OA     oa.OARouterGroup
}

var RouterGroupApp = new(RouterGroup)
