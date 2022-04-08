// 自动生成模板GoodsEvaluation
package goods

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// GoodsEvaluation 结构体
// 如果含有time.Time 请自行import time包
type GoodsEvaluation struct {
	global.GVA_MODEL
	BrandId *int     `json:"brandId" form:"brandId" gorm:"column:brand_id;comment:;size:10;"`
	CateId  *int     `json:"cateId" form:"cateId" gorm:"column:cate_id;comment:;size:10;"`
	Detail  string   `json:"detail" form:"detail" gorm:"column:detail;comment:;size:255;"`
	Name    string   `json:"name" form:"name" gorm:"column:name;comment:;size:255;"`
	Price   *int     `json:"price" form:"price" gorm:"column:price;comment:;size:10;"`
	Rank    *float64 `json:"rank" form:"rank" gorm:"column:rank;comment:;size:22;"`
	Remark  string   `json:"remark" form:"remark" gorm:"column:remark;comment:;"`
	UserId  *int     `json:"userId" form:"userId" gorm:"column:user_id;comment:;size:10;"`
}

// TableName GoodsEvaluation 表名
func (GoodsEvaluation) TableName() string {
	return "goods_evaluation"
}
