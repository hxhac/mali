// 自动生成模板RssCategory
package rss

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/helper/time"
	"gorm.io/gorm"
)

// RssCategory 结构体
// 如果含有time.Time 请自行import time包
type RssCategory struct {
	global.GVA_MODEL
	Uuid            string `json:"uuid" form:"uuid" gorm:"column:uuid;comment:;size:10;"`
	CateName        string `json:"cateName" form:"cateName" gorm:"column:cate_name;comment:;size:255;"`
	Num             *int   `json:"num" form:"num" gorm:"column:num;comment:限制聚合feed源的item数量;size:10;default:99;"`
	Remark          string `json:"remark" form:"remark" gorm:"column:remark;comment:自定义该rss描述;size:255;"`
	Author          string `json:"author" form:"author" gorm:"column:author;comment:;size:255;"`
	IsMute          *bool  `json:"isMute" form:"isMute" gorm:"column:is_mute;comment:是否mute整个分类（默认0展示1不展示）;default:0;"`
	IsUpdate        *bool  `json:"isUpdate" form:"isUpdate" gorm:"column:is_update;comment:是否正常更新(默认0不正常更新1正常更新);default:0;"`
	UpdateTimeStub  string `json:"updateTimeStub" form:"updateTimeStub" gorm:"column:update_time_stub;comment:更新时间stub;size:255;"`
	UpdateTimestamp *int   `json:"updateTimestamp" form:"updateTimestamp" gorm:"column:update_timestamp;comment:该分类更新时间戳;size:20;default:0;not null;"`
	Cron            string `json:"cron" form:"cron" gorm:"column:cron;comment:;size:255;"`
}

// TableName RssCategory 表名
func (RssCategory) TableName() string {
	return "rss_category"
}

func (rc RssCategory) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(rc).UpdateColumn("update_timestamp", time.DiffDateTime(rc.UpdateTimeStub))
	return
}

// 在同一个事务中更新数据
func (rc RssCategory) AfterUpdate(tx *gorm.DB) (err error) {
	tx.Model(rc).UpdateColumn("update_timestamp", time.DiffDateTime(rc.UpdateTimeStub))
	return
}
