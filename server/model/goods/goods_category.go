// 自动生成模板GoodsCategory
package goods

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// GoodsCategory 结构体
// 如果含有time.Time 请自行import time包
type GoodsCategory struct {
	global.GVA_MODEL
	Pid        *int            `json:"pid" form:"pid" gorm:"column:pid;comment:;size:10;"`
	CateName   string          `json:"cateName" form:"cateName" gorm:"column:cate_name;comment:分类名称;size:255;unique;"`
	More       string          `json:"more" form:"more" gorm:"column:more;comment:备注;"`
	IsWishList *bool           `json:"isWishList" form:"isWishList" gorm:"is_wish_list;comment:是否在待购清单;default:0;not null;"`
	IsDisable  *bool           `json:"isDisable" form:"isDisable" gorm:"is_disable;comment:是否disable某些分类，默认不disable(0);default:0;not null;"`
	Children   []GoodsCategory `json:"children" gorm:"-"`
}

// TableName GoodsCategory 表名
func (GoodsCategory) TableName() string {
	return "goods_category"
}
