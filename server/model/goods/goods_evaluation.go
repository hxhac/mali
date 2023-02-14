// 自动生成模板GoodsEvaluation
package goods

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// GoodsEvaluation 结构体
// 如果含有time.Time 请自行import time包
type GoodsEvaluation struct {
	GoodsLabel GoodsLabel `json:"goods_label" gorm:"foreignKey:Label"`
	GoodsBrand GoodsBrand `json:"goods_brand" gorm:"foreignKey:Brand"`
	Brand      *int       `json:"brand" form:"brand" gorm:"column:brand;comment:品牌;not null;"`
	Score      *int       `json:"score" form:"score" gorm:"column:score;comment:评分;size:10;not null;"`
	UseTimes   *int       `json:"useTimes" form:"useTimes" gorm:"column:use_times;comment:复购使用次数;not null;default:1"`
	Price      *int       `json:"price" form:"price" gorm:"column:price;comment:商品价格;not null;"`
	Label      *int       `json:"label" form:"label" gorm:"column:label;comment:商品标签;not null;"`
	global.GVA_MODEL
	More      string        `json:"more" form:"more" gorm:"column:more;comment:;"`
	CleanCron string        `json:"cleanCron" form:"cleanCron" gorm:"column:clean_cron;comment:消耗品的更换周期/耐用品的清洁周期;"`
	BuyCron   string        `json:"buyCron" form:"buyCron" gorm:"column:buy_cron;comment:复购周期;"`
	GoodsName string        `json:"goodsName" form:"goodsName" gorm:"column:goods_name;comment:商品名称;size:255;not null;"`
	Category  pq.Int64Array `json:"category" form:"category[]" gorm:"type:varchar(255)[];column:category;comment:分类;size:255;not null;"`
}

// TableName GoodsEvaluation 表名
func (GoodsEvaluation) TableName() string {
	return "goods_evaluation"
}

func (ge GoodsEvaluation) AfterCreate(tx *gorm.DB) (err error) {
	// 拼接分类名称
	res := []string{}
	var goodsCategory GoodsCategory
	for _, v := range ge.Category {
		if err = global.GVA_DB.Where("id = ?", v).Select("cate_name").First(&goodsCategory).Error; err != nil {
			return
		} else {
			res = append(res, goodsCategory.CateName)
		}
	}
	// tx.Model(ge).UpdateColumn("category_name", strings.Join(res, " / "))

	// 给该商品在goods_brand表中的num字段加1
	var gb GoodsBrand
	tx.Model(gb).Where("id = ?", ge.Brand).UpdateColumn("num", gorm.Expr("num + ?", 1))
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
	// tx.Model(ge).UpdateColumn("category_name", strings.Join(res, " / "))
	return
}

// 因为是硬删除，所以要在BeforeDelete里操作
func (ge GoodsEvaluation) BeforeDelete(tx *gorm.DB) (err error) {
	// 把该商品在goods_brand表中的num字段减1
	var gb GoodsBrand
	// 子查询处理
	sub := tx.Model(ge).Select("brand").Where("id", ge.GVA_MODEL.ID)
	tx.Model(gb).Where("id = (?)", sub).UpdateColumn("num", gorm.Expr("num - ?", 1))
	return
}
