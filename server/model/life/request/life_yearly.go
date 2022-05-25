package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/life"
)

type LifeYearlySearch struct {
	life.LifeYearly
	request.PageInfo
}
