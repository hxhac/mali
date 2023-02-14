package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type JwtBlacklist struct {
	Jwt string `gorm:"type:text;comment:jwt"`
	global.GVA_MODEL
}
