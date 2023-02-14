package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/satori/go.uuid"
)

type SysUser struct {
	Authority SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	global.GVA_MODEL
	HeaderImg   string         `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"`
	Password    string         `json:"-"  gorm:"comment:用户登录密码"`
	NickName    string         `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`
	SideMode    string         `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`
	Username    string         `json:"userName" gorm:"comment:用户登录名"`
	BaseColor   string         `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`
	ActiveColor string         `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"`
	AuthorityId string         `json:"authorityId" gorm:"default:888;comment:用户角色ID"`
	Phone       string         `json:"phone"  gorm:"comment:用户手机号"`
	Email       string         `json:"email"  gorm:"comment:用户邮箱"`
	Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
	UUID        uuid.UUID      `json:"uuid" gorm:"comment:用户UUID"`
}
