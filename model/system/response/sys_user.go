package response

import "MEIS-server/model/system"

type UserLoginAfter struct {
	User      *system.SysUser `json:"user"`
	Token     string          `json:"token"`
	ExpiresAt int64           `json:"expires_at"`
}

type UserInfo struct {
	User  system.SysUser   `json:"user"`
	Menus []system.SysMenu `json:"menus"`
}
