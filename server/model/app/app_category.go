// 自动生成模板AppCategory
package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AppCategory 结构体
// 如果含有time.Time 请自行import time包
type AppCategory struct {
	global.GVA_MODEL
	CategoryName string `json:"categoryName" form:"categoryName" gorm:"column:category_name;comment:分类名称;size:255;"`
	More         string `json:"more" form:"more" gorm:"column:more;comment:备注;"`
}

// TableName AppCategory 表名
func (AppCategory) TableName() string {
	return "app_category"
}
