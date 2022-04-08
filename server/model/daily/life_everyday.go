// 自动生成模板LifeEveryday
package daily

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// LifeEveryday 结构体
// 如果含有time.Time 请自行import time包
type LifeEveryday struct {
      global.GVA_MODEL
      Prefix  string `json:"prefix" form:"prefix" gorm:"column:prefix;comment:前缀;size:255;"`
      Task  string `json:"task" form:"task" gorm:"column:task;comment:;size:255;"`
      Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:任务备注;size:255;"`
      Duration  string `json:"duration" form:"duration" gorm:"column:duration;comment:;size:255;"`
      TimeStub  string `json:"timeStub" form:"timeStub" gorm:"column:time_stub;comment:;size:255;"`
}


// TableName LifeEveryday 表名
func (LifeEveryday) TableName() string {
  return "life_everyday"
}

