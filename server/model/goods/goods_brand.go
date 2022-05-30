// 自动生成模板GoodsBrand
package goods

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// GoodsBrand 结构体
// 如果含有time.Time 请自行import time包
type GoodsBrand struct {
	global.GVA_MODEL
	BrandName string `json:"brandName" form:"brandName" gorm:"column:brand_name;comment:;size:255;unique;"`
	More      string `json:"more" form:"more" gorm:"column:more;comment:;"`
	Num       *int   `json:"num" form:"num" gorm:"column:num;default:0;comment:该品牌商品数"`
}

// TableName GoodsBrand 表名
func (GoodsBrand) TableName() string {
	return "goods_brand"
}
