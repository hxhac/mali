package rss

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/rss"
	rssReq "github.com/flipped-aurora/gin-vue-admin/server/model/rss/request"
)

type RssCategoryService struct{}

// CreateRssCategory 创建RssCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (rssCategoryService *RssCategoryService) CreateRssCategory(rssCategory rss.RssCategory) (err error) {
	err = global.GVA_DB.Create(&rssCategory).Error
	return err
}

// DeleteRssCategory 删除RssCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (rssCategoryService *RssCategoryService) DeleteRssCategory(rssCategory rss.RssCategory) (err error) {
	err = global.GVA_DB.Delete(&rssCategory).Error
	return err
}

// DeleteRssCategoryByIds 批量删除RssCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (rssCategoryService *RssCategoryService) DeleteRssCategoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]rss.RssCategory{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateRssCategory 更新RssCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (rssCategoryService *RssCategoryService) UpdateRssCategory(rssCategory rss.RssCategory) (err error) {
	err = global.GVA_DB.Save(&rssCategory).Error
	return err
}

// GetRssCategory 根据id获取RssCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (rssCategoryService *RssCategoryService) GetRssCategory(id uint) (err error, rssCategory rss.RssCategory) {
	err = global.GVA_DB.Where("id = ?", id).First(&rssCategory).Error
	return
}

// GetRssCategory 根据id获取RssCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (rssCategoryService *RssCategoryService) GetRssCategoryByUUID(uuid string) (err error, rssCategory rss.RssCategory) {
	err = global.GVA_DB.Where("uuid = ?", uuid).First(&rssCategory).Error
	return
}

// GetRssURLs 获取所有该分类下的feed源
func (rssCategoryService *RssCategoryService) GetRssURLs(uuid string) (err error, result []string) {
	// type Result struct {
	// 	url string
	// }
	// var result []Result
	db := global.GVA_DB.Model(&rss.RssCategory{})
	err = db.Raw(`select url
from rss_feed
where cate_id = (select id from rss_category where uuid = ?)
  and is_pause = false
  and is_deleted != null
order by is_starred desc;`, uuid).Scan(&result).Error

	return err, result
}

// GetRssCategoryInfoList 分页获取RssCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (rssCategoryService *RssCategoryService) GetRssCategoryInfoList(info rssReq.RssCategorySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&rss.RssCategory{})
	var rssCategorys []rss.RssCategory
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	// err = db.Limit(limit).Offset(offset).Order("").Find(&rssCategorys).Error
	err = db.Raw(`select rc.*, count(rf.id) as num
from rss_category as rc
         left join rss_feed as rf on rc.id = rf.cate_id
group by rc.id
order by num desc;`).Limit(limit).Offset(offset).Find(&rssCategorys).Error
	return err, rssCategorys, total
}
