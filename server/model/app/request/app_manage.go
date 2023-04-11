package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppManageSearch struct {
	app.AppManage
	request.PageInfo
}
