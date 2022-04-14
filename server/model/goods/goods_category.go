// 自动生成模板GoodsCategory
package goods

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// GoodsCategory 结构体
// 如果含有time.Time 请自行import time包
type GoodsCategory struct {
	global.GVA_MODEL
	CateName string `json:"cate_name" form:"cate_name" gorm:"column:cate_name;comment:;size:255;"`
}

// TableName GoodsCategory 表名
func (GoodsCategory) TableName() string {
	return "goods_category"
}
