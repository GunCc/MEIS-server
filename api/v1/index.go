package v1

import (
	"MEIS-server/api/v1/oa"
	"MEIS-server/api/v1/system"
)

type ApiGroup struct {
	SystemApi system.SystemApi
	OAApi     oa.OAApi
}
