// 自动生成模板AppManage
package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AppManage 结构体
// 如果含有time.Time 请自行import time包
type AppManage struct {
	global.GVA_MODEL
	AppLabel   *int   `json:"appLabel" form:"appLabel" gorm:"column:app_label;comment:app标签（0核心应用、1增强应用、2已删除应用）;size:10;"`
	AppMore    string `json:"appMore" form:"appMore" gorm:"column:app_more;comment:md-textarea;"`
	AppName    string `json:"appName" form:"appName" gorm:"column:app_name;comment:app名称;size:255;"`
	AppRemark  string `json:"appRemark" form:"appRemark" gorm:"column:app_remark;comment:app备注;size:255;"`
	AppUrl     string `json:"appUrl" form:"appUrl" gorm:"column:app_url;comment:app的url;size:255;"`
	CategoryId *int   `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:分类id;size:10;"`
}

// TableName AppManage 表名
func (AppManage) TableName() string {
	return "app_manage"
}
