// 自动生成模板RssFeed
package rss

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// RssFeed 结构体
// 如果含有time.Time 请自行import time包
type RssFeed struct {
	global.GVA_MODEL
	RssName     string      `json:"rssName" form:"rssName" gorm:"column:rss_name;comment:item名称;size:255;not null; unique;"`
	CateId      *int        `json:"cateId" form:"cateId" gorm:"column:cate_id;comment:rss分类;size:10;"`
	Url         string      `json:"url" form:"url" gorm:"column:url;comment:rss的url;size:255;"`
	Keywords    string      `json:"keywords" form:"keywords" gorm:"column:keywords;comment:关键字过滤，逗号连接;size:255;"`
	Remark      string      `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`
	IsPause     *bool       `json:"isPause" form:"isPause" gorm:"column:is_pause;comment:是否暂停（0不暂停1暂停）;default:0"`
	IsStarred   *bool       `json:"isStarred" form:"isStarred" gorm:"column:is_starred;comment:是否收藏;default:0;not null;"`
	RssCategory RssCategory `json:"rss_category" gorm:"foreignKey:CateId"`
}

// TableName RssFeed 表名
func (RssFeed) TableName() string {
	return "rss_feed"
}
