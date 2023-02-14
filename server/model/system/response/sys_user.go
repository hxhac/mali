package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type SysUserResponse struct {
	User system.SysUser `json:"user"`
}

type LoginResponse struct {
	Token     string         `json:"token"`
	User      system.SysUser `json:"user"`
	ExpiresAt int64          `json:"expiresAt"`
}
