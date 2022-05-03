package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/goods"
)

type GoodsCategorySearch struct {
	goods.GoodsCategory
	request.PageInfo
}
