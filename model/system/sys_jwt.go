package system

import "MEIS-server/global"

type JwtBlacklist struct {
	global.MEIS_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
