package goods

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/goods"
	goodsReq "github.com/flipped-aurora/gin-vue-admin/server/model/goods/request"
)

type GoodsCategoryService struct{}

// CreateGoodsCategory 创建GoodsCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsCategoryService *GoodsCategoryService) CreateGoodsCategory(goodsCategory goods.GoodsCategory) (err error) {
	err = global.GVA_DB.Create(&goodsCategory).Error
	return err
}

// DeleteGoodsCategory 删除GoodsCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsCategoryService *GoodsCategoryService) DeleteGoodsCategory(goodsCategory goods.GoodsCategory) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&goodsCategory).Error
	return err
}

// DeleteGoodsCategoryByIds 批量删除GoodsCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsCategoryService *GoodsCategoryService) DeleteGoodsCategoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&[]goods.GoodsCategory{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateGoodsCategory 更新GoodsCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsCategoryService *GoodsCategoryService) UpdateGoodsCategory(goodsCategory goods.GoodsCategory) (err error) {
	err = global.GVA_DB.Save(&goodsCategory).Error
	return err
}

// GetGoodsCategory 根据id获取GoodsCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsCategoryService *GoodsCategoryService) GetGoodsCategory(id uint) (err error, goodsCategory goods.GoodsCategory) {
	err = global.GVA_DB.Where("id = ?", id).First(&goodsCategory).Error
	return
}

// GetGoodsCategoryInfoList 分页获取GoodsCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsCategoryService *GoodsCategoryService) GetGoodsCategoryInfoList(info goodsReq.GoodsCategorySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&goods.GoodsCategory{})
	var goodsCategorys []goods.GoodsCategory
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Where("pid = ?", "0").Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Where("pid = ?", "0").Find(&goodsCategorys).Error
	if len(goodsCategorys) > 0 {
		for k := range goodsCategorys {
			err = goodsCategoryService.findChildrenCategory(&goodsCategorys[k])
		}
	}
	return err, goodsCategorys, total
}

func (goodsCategoryService *GoodsCategoryService) findChildrenCategory(category *goods.GoodsCategory) (err error) {
	err = global.GVA_DB.Where("pid = ?", category.ID).Find(&category.Children).Error
	if len(category.Children) > 0 {
		for k := range category.Children {
			err = goodsCategoryService.findChildrenCategory(&category.Children[k])
		}
	}
	return err
}
