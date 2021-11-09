package system

import (
	"github.com/88act/go-cms/server/global"
)

type JwtBlacklist struct {
	global.GO_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
