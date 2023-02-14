package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// api分页条件查询及排序结构体
type SearchApiParams struct {
	OrderKey string `json:"orderKey"`
	system.SysApi
	request.PageInfo
	Desc bool `json:"desc"`
}
