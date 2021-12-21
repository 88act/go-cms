package system

import (
	"go-cms/global"
)

type JwtBlacklist struct {
	global.GO_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
