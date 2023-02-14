// 自动生成模板GoodsLabel
package goods

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// GoodsLabel 结构体
// 如果含有time.Time 请自行import time包
type GoodsLabel struct {
	Score     *int   `json:"score" form:"score" gorm:"column:score;comment:评分;size:10;default:1;not null;"`
	LabelName string `json:"labelName" form:"labelName" gorm:"column:label_name;comment:;size:255;"`
	Remark    string `json:"remark" form:"remark" gorm:"column:remark;comment:;size:255;"`
	Color     string `json:"color" form:"color" gorm:"column:color;comment:;size:255;"`
	global.GVA_MODEL
}

// TableName GoodsLabel 表名
func (GoodsLabel) TableName() string {
	return "goods_label"
}
