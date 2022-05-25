// 自动生成模板LifeDoc
package life

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// LifeDoc 结构体
// 如果含有time.Time 请自行import time包
type LifeDoc struct {
	global.GVA_MODEL
	Title string `json:"title" form:"title" gorm:"column:title;comment:标题;size:255;"`
	More  string `json:"more" form:"more" gorm:"column:more;comment:文章正文;"`
}

// TableName LifeDoc 表名
func (LifeDoc) TableName() string {
	return "life_doc"
}
