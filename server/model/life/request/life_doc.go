package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/life"
)

type LifeDocSearch struct {
	life.LifeDoc
	request.PageInfo
}
