package goods

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/goods"
	goodsReq "github.com/flipped-aurora/gin-vue-admin/server/model/goods/request"
)

type GoodsEvaluationService struct{}

// CreateGoodsEvaluation 创建GoodsEvaluation记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsEvaluationService *GoodsEvaluationService) CreateGoodsEvaluation(goodsEvaluation goods.GoodsEvaluation) (err error) {
	err = global.GVA_DB.Create(&goodsEvaluation).Error
	return err
}

// DeleteGoodsEvaluation 删除GoodsEvaluation记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsEvaluationService *GoodsEvaluationService) DeleteGoodsEvaluation(goodsEvaluation goods.GoodsEvaluation) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&goodsEvaluation).Error
	return err
}

// DeleteGoodsEvaluationByIds 批量删除GoodsEvaluation记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsEvaluationService *GoodsEvaluationService) DeleteGoodsEvaluationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&[]goods.GoodsEvaluation{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateGoodsEvaluation 更新GoodsEvaluation记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsEvaluationService *GoodsEvaluationService) UpdateGoodsEvaluation(goodsEvaluation goods.GoodsEvaluation) (err error) {
	err = global.GVA_DB.Save(&goodsEvaluation).Error
	return err
}

// GetGoodsEvaluation 根据id获取GoodsEvaluation记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsEvaluationService *GoodsEvaluationService) GetGoodsEvaluation(id uint) (err error, goodsEvaluation goods.GoodsEvaluation) {
	err = global.GVA_DB.Where("id = ?", id).First(&goodsEvaluation).Error
	return
}

// GetGoodsEvaluationInfoList 分页获取GoodsEvaluation记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsEvaluationService *GoodsEvaluationService) GetGoodsEvaluationInfoList(info goodsReq.GoodsEvaluationSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&goods.GoodsEvaluation{})
	var goodsEvaluations []goods.GoodsEvaluation
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Brand != nil {
		db = db.Where("brand = ?", info.Brand)
	}
	if info.Category != nil {
		res := 0
		for _, v := range info.Category {
			res = int(v)
		}
		db = db.Where("find_in_set(?, replace(replace(category, '{', ''), '}', ''))", res)
	}

	if info.GoodsName != "" {
		db = db.Where("goods_name LIKE ?", "%"+info.GoodsName+"%")
	}
	if info.More != "" {
		db = db.Where("more LIKE ?", "%"+info.More+"%")
	}
	if info.Label != nil {
		db = db.Where("label = ?", info.Label)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Preload("GoodsBrand").Preload("GoodsLabel").
		Order("score DESC").
		Order("use_times DESC").
		Order("price DESC").
		Order("id DESC").
		Limit(limit).Offset(offset).
		Find(&goodsEvaluations).Error

	// 	scw := fmt.Sprintf(`select ge.*, group_concat(gc.cate_name SEPARATOR '/') as category_name, gb.brand_name
	// from goods_evaluation as ge
	//          join goods_category as gc on FIND_IN_SET(gc.id, replace(replace(ge.category, '{', ''), '}', ''))
	//          left join goods_brand as gb on ge.brand = gb.id
	// %s
	// group by ge.id, ge.is_starred, ge.score
	// order by ge.is_starred desc, ge.score desc, ge.use_times desc
	// limit ? offset ?;`, wh)

	return err, goodsEvaluations, total
}

func (goodsEvaluationService *GoodsEvaluationService) GetGoodsEvaluationColumn(column string) (err error, columns []string) {
	err = global.GVA_DB.Model(&goods.GoodsEvaluation{}).Distinct().Pluck(column, &columns).Error
	return err, columns
}

func (goodsEvaluationService *GoodsEvaluationService) GetGoodsEvaluationByLabel(label uint) (err error, goodsEvaluations []goods.GoodsEvaluation, total int64) {

	// 创建db
	db := global.GVA_DB.Model(&goods.GoodsEvaluation{})
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("label = ?", label).Where("buy_cron is not null or clean_cron is not null")

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Preload("GoodsBrand").Preload("GoodsLabel").
		Order("score DESC").
		Order("use_times DESC").
		Order("price DESC").
		Order("id DESC").
		Find(&goodsEvaluations).Error

	return err, goodsEvaluations, total
}
