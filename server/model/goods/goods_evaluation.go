// 自动生成模板GoodsEvaluation
package goods

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// GoodsEvaluation 结构体
// 如果含有time.Time 请自行import time包
type GoodsEvaluation struct {
	global.GVA_MODEL
	GoodsName string `json:"goodsName" form:"goodsName" gorm:"column:goods_name;comment:商品名称;size:255;not null;"`
	Price     *int   `json:"price" form:"price" gorm:"column:price;comment:商品价格;not null;default:0;"`
	Brand     string `json:"brand" form:"brand" gorm:"column:brand;comment:品牌;size:255;not null;default:其他;"`
	Category  string `json:"category" form:"category" gorm:"column:category;comment:分类;size:255;not null;default:其他;"`
	Score     *int   `json:"score" form:"score" gorm:"column:score;comment:评分;size:10;not null;default:0;"`
	Remark    string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:191;"`
	More      string `json:"more" form:"more" gorm:"column:more;comment:;"`
}

// TableName GoodsEvaluation 表名
func (GoodsEvaluation) TableName() string {
	return "goods_evaluation"
}
