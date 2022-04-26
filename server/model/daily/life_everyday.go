// 自动生成模板LifeEveryday
package daily

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/helper/time"
	"gorm.io/gorm"
)

// LifeEveryday 结构体
// 如果含有time.Time 请自行import time包
type LifeEveryday struct {
	global.GVA_MODEL
	Prefix            string `json:"prefix" form:"prefix" gorm:"column:prefix;comment:前缀;size:255;"`
	Task              string `json:"task" form:"task" gorm:"column:task;comment:;size:255;"`
	Remark            string `json:"remark" form:"remark" gorm:"column:remark;comment:任务备注;size:255;"`
	More              string `json:"more" form:"more" gorm:"column:more;"`
	Duration          string `json:"duration" form:"duration" gorm:"column:duration;comment:;size:255;"`
	TimeStub          string `json:"timeStub" form:"timeStub" gorm:"column:time_stub;comment:;size:255;"`
	TimeStubTimestamp int64  `json:"timeStubTimestamp" form:"timeStubTimestamp" gorm:"column:time_stub_timestamp;comment:;"`
}

// TableName LifeEveryday 表名
func (LifeEveryday) TableName() string {
	return "life_everyday"
}

func (le LifeEveryday) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(le).UpdateColumn("time_stub_timestamp", time.CheckDateTime(le.TimeStub).Timestamp())
	return
}

// 在同一个事务中更新数据
func (le LifeEveryday) AfterUpdate(tx *gorm.DB) (err error) {
	tx.Model(le).UpdateColumn("time_stub_timestamp", time.CheckDateTime(le.TimeStub).Timestamp())
	return
}
