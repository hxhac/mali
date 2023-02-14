package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type ExaCustomer struct {
	CustomerName       string `json:"customerName" form:"customerName" gorm:"comment:客户名"`
	CustomerPhoneData  string `json:"customerPhoneData" form:"customerPhoneData" gorm:"comment:客户手机号"`
	SysUserAuthorityID string `json:"sysUserAuthorityID" form:"sysUserAuthorityID" gorm:"comment:管理角色ID"`
	global.GVA_MODEL
	SysUser   system.SysUser `json:"sysUser" form:"sysUser" gorm:"comment:管理详情"`
	SysUserID uint           `json:"sysUserId" form:"sysUserId" gorm:"comment:管理ID"`
}
