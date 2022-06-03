package life

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/life"
	lifeReq "github.com/flipped-aurora/gin-vue-admin/server/model/life/request"
)

type LifeDocService struct{}

// CreateLifeDoc 创建LifeDoc记录
// Author [piexlmax](https://github.com/piexlmax)
func (lifeDocService *LifeDocService) CreateLifeDoc(lifeDoc life.LifeDoc) (err error) {
	err = global.GVA_DB.Create(&lifeDoc).Error
	return err
}

// DeleteLifeDoc 删除LifeDoc记录
// Author [piexlmax](https://github.com/piexlmax)
func (lifeDocService *LifeDocService) DeleteLifeDoc(lifeDoc life.LifeDoc) (err error) {
	err = global.GVA_DB.Delete(&lifeDoc).Error
	return err
}

// DeleteLifeDocByIds 批量删除LifeDoc记录
// Author [piexlmax](https://github.com/piexlmax)
func (lifeDocService *LifeDocService) DeleteLifeDocByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]life.LifeDoc{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateLifeDoc 更新LifeDoc记录
// Author [piexlmax](https://github.com/piexlmax)
func (lifeDocService *LifeDocService) UpdateLifeDoc(lifeDoc life.LifeDoc) (err error) {
	err = global.GVA_DB.Save(&lifeDoc).Error
	return err
}

// GetLifeDoc 根据id获取LifeDoc记录
// Author [piexlmax](https://github.com/piexlmax)
func (lifeDocService *LifeDocService) GetLifeDoc(id uint) (err error, lifeDoc life.LifeDoc) {
	err = global.GVA_DB.Where("id = ?", id).First(&lifeDoc).Error
	return
}

// GetLifeDocInfoList 分页获取LifeDoc记录
// Author [piexlmax](https://github.com/piexlmax)
func (lifeDocService *LifeDocService) GetLifeDocInfoList(info lifeReq.LifeDocSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&life.LifeDoc{})
	var lifeDocs []life.LifeDoc
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&lifeDocs).Error
	return err, lifeDocs, total
}
