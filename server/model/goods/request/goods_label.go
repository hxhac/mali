package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/goods"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type GoodsLabelSearch struct{
    goods.GoodsLabel
    request.PageInfo
}
