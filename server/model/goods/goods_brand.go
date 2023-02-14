// 自动生成模板GoodsBrand
package goods

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// GoodsBrand 结构体
// 如果含有time.Time 请自行import time包
type GoodsBrand struct {
	Num       *int   `json:"num" form:"num" gorm:"column:num;default:0;comment:该品牌商品数"`
	IsDisable *bool  `json:"isDisable" form:"isDisable" gorm:"is_disable;comment:是否disable某个品牌，默认不disable(0);default:0;not null;"`
	BrandName string `json:"brandName" form:"brandName" gorm:"column:brand_name;comment:;size:255;unique;"`
	More      string `json:"more" form:"more" gorm:"column:more;comment:;"`
	global.GVA_MODEL
}

// TableName GoodsBrand 表名
func (GoodsBrand) TableName() string {
	return "goods_brand"
}
