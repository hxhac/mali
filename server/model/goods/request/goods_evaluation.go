package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/goods"
)

type GoodsEvaluationSearch struct {
	goods.GoodsEvaluation
	request.PageInfo
}
