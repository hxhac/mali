package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	xxxReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppManageService struct {
}

// CreateAppManage 创建AppManage记录
// Author [piexlmax](https://github.com/piexlmax)
func (appManageService *AppManageService) CreateAppManage(appManage app.AppManage) (err error) {
	err = global.GVA_DB.Create(&appManage).Error
	return err
}

// DeleteAppManage 删除AppManage记录
// Author [piexlmax](https://github.com/piexlmax)
func (appManageService *AppManageService) DeleteAppManage(appManage app.AppManage) (err error) {
	err = global.GVA_DB.Delete(&appManage).Error
	return err
}

// DeleteAppManageByIds 批量删除AppManage记录
// Author [piexlmax](https://github.com/piexlmax)
func (appManageService *AppManageService) DeleteAppManageByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]app.AppManage{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppManage 更新AppManage记录
// Author [piexlmax](https://github.com/piexlmax)
func (appManageService *AppManageService) UpdateAppManage(appManage app.AppManage) (err error) {
	err = global.GVA_DB.Save(&appManage).Error
	return err
}

// GetAppManage 根据id获取AppManage记录
// Author [piexlmax](https://github.com/piexlmax)
func (appManageService *AppManageService) GetAppManage(id uint) (err error, appManage app.AppManage) {
	err = global.GVA_DB.Where("id = ?", id).First(&appManage).Error
	return
}

// GetAppManageInfoList 分页获取AppManage记录
// Author [piexlmax](https://github.com/piexlmax)
func (appManageService *AppManageService) GetAppManageInfoList(info xxxReq.AppManageSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&app.AppManage{})
	var appManages []app.AppManage
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.CategoryId != nil {
		db = db.Where("`category_id` = ?", info.CategoryId)
	}
	if info.AppLabel != nil {
		db = db.Where("`app_label` = ?", info.AppLabel)
	}
	if info.IsUse != nil {
		db = db.Where("`is_use` = ?", info.IsUse)
	}
	if info.AppName != "" {
		db = db.Where("`app_name` LIKE ? Or `app_remark` LIKE ? Or `target` LIKE ?", "%"+info.AppName+"%", "%"+info.AppName+"%", "%"+info.AppName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("score desc").Limit(limit).Offset(offset).Find(&appManages).Error
	return err, appManages, total
}
