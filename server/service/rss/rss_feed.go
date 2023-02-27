package rss

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/rss"
	rssReq "github.com/flipped-aurora/gin-vue-admin/server/model/rss/request"
)

type RssFeedService struct{}

// CreateRssFeed 创建RssFeed记录
// Author [piexlmax](https://github.com/piexlmax)
func (rssFeedService *RssFeedService) CreateRssFeed(rssFeed rss.RssFeed) (err error) {
	err = global.GVA_DB.Create(&rssFeed).Error
	return err
}

// DeleteRssFeed 删除RssFeed记录
// Author [piexlmax](https://github.com/piexlmax)
func (rssFeedService *RssFeedService) DeleteRssFeed(rssFeed rss.RssFeed) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&rssFeed).Error
	return err
}

// DeleteRssFeedByIds 批量删除RssFeed记录
// Author [piexlmax](https://github.com/piexlmax)
func (rssFeedService *RssFeedService) DeleteRssFeedByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&[]rss.RssFeed{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateRssFeed 更新RssFeed记录
// Author [piexlmax](https://github.com/piexlmax)
func (rssFeedService *RssFeedService) UpdateRssFeed(rssFeed rss.RssFeed) (err error) {
	err = global.GVA_DB.Save(&rssFeed).Error
	return err
}

// GetRssFeed 根据id获取RssFeed记录
// Author [piexlmax](https://github.com/piexlmax)
func (rssFeedService *RssFeedService) GetRssFeed(id uint) (err error, rssFeed rss.RssFeed) {
	err = global.GVA_DB.Where("id = ?", id).First(&rssFeed).Error
	return
}

// GetRssFeedInfoList 分页获取RssFeed记录
// Author [piexlmax](https://github.com/piexlmax)
func (rssFeedService *RssFeedService) GetRssFeedInfoList(info rssReq.RssFeedSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&rss.RssFeed{})
	var rssFeeds []rss.RssFeed
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.CateId != nil {
		db = db.Where("cate_id = ?", info.CateId)
	}
	if info.RssName != "" {
		db = db.Where("rss_name LIKE ?", "%"+info.RssName+"%")
	}
	if info.IsStarred != nil {
		db = db.Where("is_starred = ?", info.IsStarred)
	}
	if info.IsPause != nil {
		db = db.Where("is_pause = ?", info.IsPause)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Preload("RssCategory").
		Order("is_starred DESC").
		Order("is_pause ASC").
		Order("score DESC").
		Order("updated_at DESC").
		Order("id DESC").
		Find(&rssFeeds).Error
	return err, rssFeeds, total
}

func (rssFeedService *RssFeedService) UpdateUpdatedTime(rssFeed rss.RssFeed) (err error) {
	err = global.GVA_DB.Select("updated_at").Where("source_url", rssFeed.SourceUrl).Updates(rssFeed).Error
	return err
}
