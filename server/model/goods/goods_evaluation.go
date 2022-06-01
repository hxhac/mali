// 自动生成模板GoodsEvaluation
package goods

import (
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// GoodsEvaluation 结构体
// 如果含有time.Time 请自行import time包
type GoodsEvaluation struct {
	global.GVA_MODEL
	GoodsName string `json:"goodsName" form:"goodsName" gorm:"column:goods_name;comment:商品名称;size:255;not null;"`
	Price     *int   `json:"price" form:"price" gorm:"column:price;comment:商品价格;not null;"`
	Brand     *int   `json:"brand" form:"brand" gorm:"column:brand;comment:品牌;not null;"`
	Score     *int   `json:"score" form:"score" gorm:"column:score;comment:评分;size:10;not null;"`
	IsStarred *bool  `json:"isStarred" form:"isStarred" gorm:"column:is_starred;comment:是否加星;"`
	Remark    string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:191;"`
	More      string `json:"more" form:"more" gorm:"column:more;comment:;"`
	UseTimes  *int   `json:"useTimes" form:"useTimes" gorm:"column:use_times;comment:复购使用次数;not null;default:1"`
	// GoodsPic   string        `json:"goodsPic" form:"goodsPic" gorm:"column:goods_pic;type:longtext;comment:商品图片;"`
	Category   pq.Int64Array `json:"category" form:"category[]" gorm:"type:varchar(255)[];column:category;comment:分类;size:255;not null;"`
	GoodsBrand GoodsBrand    `json:"goods_brand" gorm:"foreignKey:Brand"`
}

// TableName GoodsEvaluation 表名
func (GoodsEvaluation) TableName() string {
	return "goods_evaluation"
}

func (ge GoodsEvaluation) AfterCreate(tx *gorm.DB) (err error) {
	res := []string{}
	var goodsCategory GoodsCategory
	for _, v := range ge.Category {
		if err = global.GVA_DB.Where("id = ?", v).Select("cate_name").First(&goodsCategory).Error; err != nil {
			return
		} else {
			res = append(res, goodsCategory.CateName)
		}
	}
	tx.Model(ge).UpdateColumn("category_name", strings.Join(res, " / "))
	return
}

func (ge GoodsEvaluation) AfterUpdate(tx *gorm.DB) (err error) {
	res := []string{}
	var goodsCategory GoodsCategory
	for _, v := range ge.Category {
		if err = global.GVA_DB.Where("id = ?", v).Select("cate_name").First(&goodsCategory).Error; err != nil {
			return
		} else {
			res = append(res, goodsCategory.CateName)
		}
	}
	tx.Model(ge).UpdateColumn("category_name", strings.Join(res, " / "))
	return
}
