package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/daily"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type LifeYearlySearch struct{
    daily.LifeYearly
    request.PageInfo
}
