package daily

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/daily"
	dailyReq "github.com/flipped-aurora/gin-vue-admin/server/model/daily/request"
)

type LifeYearlyService struct {
}

// CreateLifeYearly 创建LifeYearly记录
// Author [piexlmax](https://github.com/piexlmax)
func (lifeYearlyService *LifeYearlyService) CreateLifeYearly(lifeYearly daily.LifeYearly) (err error) {
	err = global.GVA_DB.Create(&lifeYearly).Error
	return err
}

// DeleteLifeYearly 删除LifeYearly记录
// Author [piexlmax](https://github.com/piexlmax)
func (lifeYearlyService *LifeYearlyService) DeleteLifeYearly(lifeYearly daily.LifeYearly) (err error) {
	err = global.GVA_DB.Delete(&lifeYearly).Error
	return err
}

// DeleteLifeYearlyByIds 批量删除LifeYearly记录
// Author [piexlmax](https://github.com/piexlmax)
func (lifeYearlyService *LifeYearlyService) DeleteLifeYearlyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]daily.LifeYearly{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateLifeYearly 更新LifeYearly记录
// Author [piexlmax](https://github.com/piexlmax)
func (lifeYearlyService *LifeYearlyService) UpdateLifeYearly(lifeYearly daily.LifeYearly) (err error) {
	err = global.GVA_DB.Save(&lifeYearly).Error
	return err
}

// GetLifeYearly 根据id获取LifeYearly记录
// Author [piexlmax](https://github.com/piexlmax)
func (lifeYearlyService *LifeYearlyService) GetLifeYearly(id uint) (err error, lifeYearly daily.LifeYearly) {
	err = global.GVA_DB.Where("id = ?", id).First(&lifeYearly).Error
	return
}

// GetLifeYearlyInfoList 分页获取LifeYearly记录
// Author [piexlmax](https://github.com/piexlmax)
func (lifeYearlyService *LifeYearlyService) GetLifeYearlyInfoList(info dailyReq.LifeYearlySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&daily.LifeYearly{})
	var lifeYearlys []daily.LifeYearly
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Prefix != "" {
		db = db.Where("prefix = ?", info.Prefix)
	}
	if info.Cron != "" {
		db = db.Where("cron = ?", info.Cron)
	}
	if info.Task != "" {
		db = db.Where("task LIKE ?", "%"+info.Task+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&lifeYearlys).Error
	return err, lifeYearlys, total
}

func (lifeYearlyService *LifeYearlyService) FindAll() ([]*daily.LifeYearly, error) {
	var yearly []*daily.LifeYearly

	if err := global.GVA_DB.Find(&yearly).Error; err != nil {
		return nil, err
	}
	return yearly, nil
}

func (lifeYearlyService *LifeYearlyService) GetLifeYearlyColumn(column string) (err error, columns []string) {

	err = global.GVA_DB.Model(&daily.LifeYearly{}).Distinct().Pluck(column, &columns).Error
	return err, columns
}
