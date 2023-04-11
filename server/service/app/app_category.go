package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	xxxReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppCategoryService struct {
}

// CreateAppCategory 创建AppCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCategoryService *AppCategoryService) CreateAppCategory(appCategory app.AppCategory) (err error) {
	err = global.GVA_DB.Create(&appCategory).Error
	return err
}

// DeleteAppCategory 删除AppCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCategoryService *AppCategoryService) DeleteAppCategory(appCategory app.AppCategory) (err error) {
	err = global.GVA_DB.Delete(&appCategory).Error
	return err
}

// DeleteAppCategoryByIds 批量删除AppCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCategoryService *AppCategoryService) DeleteAppCategoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]app.AppCategory{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppCategory 更新AppCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCategoryService *AppCategoryService) UpdateAppCategory(appCategory app.AppCategory) (err error) {
	err = global.GVA_DB.Save(&appCategory).Error
	return err
}

// GetAppCategory 根据id获取AppCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCategoryService *AppCategoryService) GetAppCategory(id uint) (err error, appCategory app.AppCategory) {
	err = global.GVA_DB.Where("id = ?", id).First(&appCategory).Error
	return
}

// GetAppCategoryInfoList 分页获取AppCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCategoryService *AppCategoryService) GetAppCategoryInfoList(info xxxReq.AppCategorySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&app.AppCategory{})
	var appCategorys []app.AppCategory
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&appCategorys).Error
	return err, appCategorys, total
}
