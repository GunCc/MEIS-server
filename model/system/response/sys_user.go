package system

import "MEIS-server/model/system"

type UserLoginAfter struct {
	User      *system.SysUser `json:"user"`
	Token     string          `json:"token"`
	ExpiresAt int64           `json:"expires_at"`
}
