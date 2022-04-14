// 自动生成模板GoodsEvaluation
package goods

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// GoodsEvaluation 结构体
// 如果含有time.Time 请自行import time包
type GoodsEvaluation struct {
	global.GVA_MODEL
	Name          string        `json:"name" form:"name" gorm:"column:name;comment:;size:255;"`
	Price         *int          `json:"price" form:"price" gorm:"column:price;comment:;size:10;"`
	CateId        *int          `json:"cateId" form:"cateId" gorm:"column:cate_id;comment:;size:10;"`
	Score         *float64      `json:"score" form:"score" gorm:"column:score;comment:;size:22;"`
	Remark        string        `json:"remark" form:"remark" gorm:"column:remark;comment:;"`
	GoodsCategory GoodsCategory `json:"goods_category" gorm:"foreignKey:CateId;"`
}

// TableName GoodsEvaluation 表名
func (GoodsEvaluation) TableName() string {
	return "goods_evaluation"
}
