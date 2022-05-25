package life

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/life"
	dailyReq "github.com/flipped-aurora/gin-vue-admin/server/model/life/request"
)

type LifeEverydayService struct{}

// CreateLifeEveryday 创建LifeEveryday记录
// Author [piexlmax](https://github.com/piexlmax)
func (lifeEverydayService *LifeEverydayService) CreateLifeEveryday(lifeEveryday life.LifeEveryday) (err error) {
	err = global.GVA_DB.Create(&lifeEveryday).Error
	return err
}

// DeleteLifeEveryday 删除LifeEveryday记录
// Author [piexlmax](https://github.com/piexlmax)
func (lifeEverydayService *LifeEverydayService) DeleteLifeEveryday(lifeEveryday life.LifeEveryday) (err error) {
	err = global.GVA_DB.Delete(&lifeEveryday).Error
	return err
}

// DeleteLifeEverydayByIds 批量删除LifeEveryday记录
// Author [piexlmax](https://github.com/piexlmax)
func (lifeEverydayService *LifeEverydayService) DeleteLifeEverydayByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]life.LifeEveryday{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateLifeEveryday 更新LifeEveryday记录
// Author [piexlmax](https://github.com/piexlmax)
func (lifeEverydayService *LifeEverydayService) UpdateLifeEveryday(lifeEveryday life.LifeEveryday) (err error) {
	err = global.GVA_DB.Save(&lifeEveryday).Error
	return err
}

// GetLifeEveryday 根据id获取LifeEveryday记录
// Author [piexlmax](https://github.com/piexlmax)
func (lifeEverydayService *LifeEverydayService) GetLifeEveryday(id uint) (err error, lifeEveryday life.LifeEveryday) {
	err = global.GVA_DB.Where("id = ?", id).First(&lifeEveryday).Error
	return
}

// GetLifeEverydayInfoList 分页获取LifeEveryday记录
// Author [piexlmax](https://github.com/piexlmax)
func (lifeEverydayService *LifeEverydayService) GetLifeEverydayInfoList(info dailyReq.LifeEverydaySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&life.LifeEveryday{})
	var lifeEverydays []life.LifeEveryday
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Prefix != "" {
		db = db.Where("prefix = ?", info.Prefix)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("time_stub_timestamp ASC").Limit(limit).Offset(offset).Find(&lifeEverydays).Error
	return err, lifeEverydays, total
}

func (lifeEverydayService *LifeEverydayService) FindAll() ([]*life.LifeEveryday, error) {
	var everyday []*life.LifeEveryday

	if err := global.GVA_DB.Find(&everyday).Error; err != nil {
		return nil, err
	}
	return everyday, nil
}
