package goods

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/goods"
	goodsReq "github.com/flipped-aurora/gin-vue-admin/server/model/goods/request"
)

type GoodsLabelService struct{}

// CreateGoodsLabel 创建GoodsLabel记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsLabelService *GoodsLabelService) CreateGoodsLabel(goodsLabel goods.GoodsLabel) (err error) {
	err = global.GVA_DB.Create(&goodsLabel).Error
	return err
}

// DeleteGoodsLabel 删除GoodsLabel记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsLabelService *GoodsLabelService) DeleteGoodsLabel(goodsLabel goods.GoodsLabel) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&goodsLabel).Error
	return err
}

// DeleteGoodsLabelByIds 批量删除GoodsLabel记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsLabelService *GoodsLabelService) DeleteGoodsLabelByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&[]goods.GoodsLabel{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateGoodsLabel 更新GoodsLabel记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsLabelService *GoodsLabelService) UpdateGoodsLabel(goodsLabel goods.GoodsLabel) (err error) {
	err = global.GVA_DB.Save(&goodsLabel).Error
	return err
}

// GetGoodsLabel 根据id获取GoodsLabel记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsLabelService *GoodsLabelService) GetGoodsLabel(id uint) (err error, goodsLabel goods.GoodsLabel) {
	err = global.GVA_DB.Where("id = ?", id).First(&goodsLabel).Error
	return
}

// GetGoodsLabelInfoList 分页获取GoodsLabel记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsLabelService *GoodsLabelService) GetGoodsLabelInfoList(info goodsReq.GoodsLabelSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&goods.GoodsLabel{})
	var goodsLabels []goods.GoodsLabel
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("score DESC").Find(&goodsLabels).
		Error
	return err, goodsLabels, total
}
