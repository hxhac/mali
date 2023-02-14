package system

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AutoCodeStruct 初始版本自动化代码工具
type AutoCodeStruct struct {
	StructName         string   `json:"structName"`
	TableName          string   `json:"tableName"`
	PackageName        string   `json:"packageName"`
	HumpPackageName    string   `json:"humpPackageName"`
	Abbreviation       string   `json:"abbreviation"`
	Description        string   `json:"description"`
	Package            string   `json:"package"`
	PackageT           string   `json:"-"`
	Fields             []*Field `json:"fields"`
	DictTypes          []string `json:"-"`
	AutoCreateApiToSql bool     `json:"autoCreateApiToSql"`
	AutoMoveFile       bool     `json:"autoMoveFile"`
}

type Field struct {
	FieldName       string `json:"fieldName"`       // Field名
	FieldDesc       string `json:"fieldDesc"`       // 中文名
	FieldType       string `json:"fieldType"`       // Field数据类型
	FieldJson       string `json:"fieldJson"`       // FieldJson
	DataTypeLong    string `json:"dataTypeLong"`    // 数据库字段长度
	Comment         string `json:"comment"`         // 数据库字段描述
	ColumnName      string `json:"columnName"`      // 数据库字段
	FieldSearchType string `json:"fieldSearchType"` // 搜索条件
	DictType        string `json:"dictType"`        // 字典
}

var AutoMoveErr error = errors.New("创建代码成功并移动文件成功")

type SysAutoCode struct {
	PackageName string `json:"packageName" gorm:"comment:包名"`
	Label       string `json:"label" gorm:"comment:展示名"`
	Desc        string `json:"desc" gorm:"comment:描述"`
	global.GVA_MODEL
}
