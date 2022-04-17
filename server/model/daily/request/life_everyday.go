package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/daily"
)

type LifeEverydaySearch struct {
	daily.LifeEveryday
	request.PageInfo
}
