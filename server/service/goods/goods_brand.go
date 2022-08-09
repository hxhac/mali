package goods

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/goods"
	goodsReq "github.com/flipped-aurora/gin-vue-admin/server/model/goods/request"
)

type GoodsBrandService struct{}

// CreateGoodsBrand 创建GoodsBrand记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsBrandService *GoodsBrandService) CreateGoodsBrand(goodsBrand goods.GoodsBrand) (err error) {
	err = global.GVA_DB.Create(&goodsBrand).Error
	return err
}

// DeleteGoodsBrand 删除GoodsBrand记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsBrandService *GoodsBrandService) DeleteGoodsBrand(goodsBrand goods.GoodsBrand) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&goodsBrand).Error
	return err
}

// DeleteGoodsBrandByIds 批量删除GoodsBrand记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsBrandService *GoodsBrandService) DeleteGoodsBrandByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&[]goods.GoodsBrand{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateGoodsBrand 更新GoodsBrand记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsBrandService *GoodsBrandService) UpdateGoodsBrand(goodsBrand goods.GoodsBrand) (err error) {
	err = global.GVA_DB.Save(&goodsBrand).Error
	return err
}

// GetGoodsBrand 根据id获取GoodsBrand记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsBrandService *GoodsBrandService) GetGoodsBrand(id uint) (err error, goodsBrand goods.GoodsBrand) {
	err = global.GVA_DB.Where("id = ?", id).First(&goodsBrand).Error
	return
}

// GetGoodsBrandInfoList 分页获取GoodsBrand记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsBrandService *GoodsBrandService) GetGoodsBrandInfoList(info goodsReq.GoodsBrandSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&goods.GoodsBrand{})
	var goodsBrands []goods.GoodsBrand
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	// err = db.Table("goods_evaluation as ge").
	// 	Select("gb.*, count(ge.brand) as num").
	// 	Joins("left join goods_brand as gb on gb.id = ge.brand").
	// 	Group("ge.brand").
	// 	Order("count(ge.brand) DESC").
	// 	Limit(limit).Offset(offset).Find(&goodsBrands).Error

	// err = db.Table("goods_evaluation as ge").
	// 	Select("gb.*, count(ge.brand) as num").
	// 	Joins("left join goods_brand as gb on gb.id = ge.brand").
	// 	Group("ge.brand").
	// 	Order("count(ge.brand) DESC").
	// 	Limit(limit).Offset(offset).Find(&goodsBrands).Error

	err = db.Raw(`select gb.*, count(ge.brand) as num
from goods_brand as gb
         left join goods_evaluation as ge on gb.id = ge.brand
group by gb.id
order by num desc;`).Limit(limit).Offset(offset).Find(&goodsBrands).Error
	return err, goodsBrands, total
}
