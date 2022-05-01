package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/goods"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type GoodsBrandSearch struct{
    goods.GoodsBrand
    request.PageInfo
}
