package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/goods"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type GoodsEvaluationSearch struct{
    goods.GoodsEvaluation
    request.PageInfo
}
