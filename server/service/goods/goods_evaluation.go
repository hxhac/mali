package goods

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/goods"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    goodsReq "github.com/flipped-aurora/gin-vue-admin/server/model/goods/request"
)

type GoodsEvaluationService struct {
}

// CreateGoodsEvaluation 创建GoodsEvaluation记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsEvaluationService *GoodsEvaluationService) CreateGoodsEvaluation(goodsEvaluation goods.GoodsEvaluation) (err error) {
	err = global.GVA_DB.Create(&goodsEvaluation).Error
	return err
}

// DeleteGoodsEvaluation 删除GoodsEvaluation记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsEvaluationService *GoodsEvaluationService)DeleteGoodsEvaluation(goodsEvaluation goods.GoodsEvaluation) (err error) {
	err = global.GVA_DB.Delete(&goodsEvaluation).Error
	return err
}

// DeleteGoodsEvaluationByIds 批量删除GoodsEvaluation记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsEvaluationService *GoodsEvaluationService)DeleteGoodsEvaluationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]goods.GoodsEvaluation{},"id in ?",ids.Ids).Error
	return err
}

// UpdateGoodsEvaluation 更新GoodsEvaluation记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsEvaluationService *GoodsEvaluationService)UpdateGoodsEvaluation(goodsEvaluation goods.GoodsEvaluation) (err error) {
	err = global.GVA_DB.Save(&goodsEvaluation).Error
	return err
}

// GetGoodsEvaluation 根据id获取GoodsEvaluation记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsEvaluationService *GoodsEvaluationService)GetGoodsEvaluation(id uint) (err error, goodsEvaluation goods.GoodsEvaluation) {
	err = global.GVA_DB.Where("id = ?", id).First(&goodsEvaluation).Error
	return
}

// GetGoodsEvaluationInfoList 分页获取GoodsEvaluation记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsEvaluationService *GoodsEvaluationService)GetGoodsEvaluationInfoList(info goodsReq.GoodsEvaluationSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&goods.GoodsEvaluation{})
    var goodsEvaluations []goods.GoodsEvaluation
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&goodsEvaluations).Error
	return err, goodsEvaluations, total
}
