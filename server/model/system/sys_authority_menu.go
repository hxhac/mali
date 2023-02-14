package system

type SysMenu struct {
	Btns        map[string]string      `json:"btns" gorm:"-"`
	MenuId      string                 `json:"menuId" gorm:"comment:菜单ID"`
	AuthorityId string                 `json:"-" gorm:"comment:角色ID"`
	Children    []SysMenu              `json:"children" gorm:"-"`
	Parameters  []SysBaseMenuParameter `json:"parameters" gorm:"foreignKey:SysBaseMenuID;references:MenuId"`
	SysBaseMenu
}

func (s SysMenu) TableName() string {
	return "authority_menu"
}
