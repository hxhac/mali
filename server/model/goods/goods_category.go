// 自动生成模板GoodsCategory
package goods

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// GoodsCategory 结构体
// 如果含有time.Time 请自行import time包
type GoodsCategory struct {
      global.GVA_MODEL
      Pid  *int `json:"pid" form:"pid" gorm:"column:pid;comment:;size:10;"`
      CateName  string `json:"cateName" form:"cateName" gorm:"column:cate_name;comment:分类名称;size:255;"`
      More  string `json:"more" form:"more" gorm:"column:more;comment:备注;"`
}


// TableName GoodsCategory 表名
func (GoodsCategory) TableName() string {
  return "goods_category"
}

