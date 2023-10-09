package api

import v1 "MEIS-server/api/v1"

type ApiGroup struct {
	v1.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
