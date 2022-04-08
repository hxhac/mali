// 自动生成模板LifeYearly
package daily

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// LifeYearly 结构体
// 如果含有time.Time 请自行import time包
type LifeYearly struct {
      global.GVA_MODEL
      Cron  string `json:"cron" form:"cron" gorm:"column:cron;comment:;size:255;"`
      Prefix  string `json:"prefix" form:"prefix" gorm:"column:prefix;comment:;size:255;"`
      Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:;size:255;"`
      Task  string `json:"task" form:"task" gorm:"column:task;comment:;size:255;"`
}


// TableName LifeYearly 表名
func (LifeYearly) TableName() string {
  return "life_yearly"
}

