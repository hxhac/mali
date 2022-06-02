// 自动生成模板RssCategory
package rss

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// RssCategory 结构体
// 如果含有time.Time 请自行import time包
type RssCategory struct {
      global.GVA_MODEL
      Uuid  *int `json:"uuid" form:"uuid" gorm:"column:uuid;comment:;size:10;"`
      CateName  string `json:"cateName" form:"cateName" gorm:"column:cate_name;comment:;size:255;"`
      Title  string `json:"title" form:"title" gorm:"column:title;comment:自定义标题;size:255;"`
      Num  *int `json:"num" form:"num" gorm:"column:num;comment:限制聚合feed源的item数量;size:10;"`
      Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:自定义该rss描述;size:255;"`
      Author  string `json:"author" form:"author" gorm:"column:author;comment:;size:255;"`
      IsMute  *bool `json:"isMute" form:"isMute" gorm:"column:is_mute;comment:是否mute整个分类（默认0展示1不展示）;"`
}


// TableName RssCategory 表名
func (RssCategory) TableName() string {
  return "rss_category"
}

