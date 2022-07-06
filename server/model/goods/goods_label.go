// 自动生成模板GoodsLabel
package goods

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// GoodsLabel 结构体
// 如果含有time.Time 请自行import time包
type GoodsLabel struct {
	global.GVA_MODEL
	LabelName string `json:"labelName" form:"labelName" gorm:"column:label_name;comment:;size:255;"`
	Remark    string `json:"remark" form:"remark" gorm:"column:remark;comment:;size:255;"`
	Color     string `json:"color" form:"color" gorm:"column:color;comment:;size
:255;"`
}

// TableName GoodsLabel 表名
func (GoodsLabel) TableName() string {
	return "goods_label"
}
