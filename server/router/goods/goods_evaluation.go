package goods

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GoodsEvaluationRouter struct {
}

// InitGoodsEvaluationRouter 初始化 GoodsEvaluation 路由信息
func (s *GoodsEvaluationRouter) InitGoodsEvaluationRouter(Router *gin.RouterGroup) {
	goodsEvaluationRouter := Router.Group("goodsEvaluation").Use(middleware.OperationRecord())
	goodsEvaluationRouterWithoutRecord := Router.Group("goodsEvaluation")
	var goodsEvaluationApi = v1.ApiGroupApp.GoodsApiGroup.GoodsEvaluationApi
	{
		goodsEvaluationRouter.POST("createGoodsEvaluation", goodsEvaluationApi.CreateGoodsEvaluation)             // 新建GoodsEvaluation
		goodsEvaluationRouter.DELETE("deleteGoodsEvaluation", goodsEvaluationApi.DeleteGoodsEvaluation)           // 删除GoodsEvaluation
		goodsEvaluationRouter.DELETE("deleteGoodsEvaluationByIds", goodsEvaluationApi.DeleteGoodsEvaluationByIds) // 批量删除GoodsEvaluation
		goodsEvaluationRouter.PUT("updateGoodsEvaluation", goodsEvaluationApi.UpdateGoodsEvaluation)              // 更新GoodsEvaluation
	}
	{
		goodsEvaluationRouterWithoutRecord.GET("findGoodsEvaluation", goodsEvaluationApi.FindGoodsEvaluation)       // 根据ID获取GoodsEvaluation
		goodsEvaluationRouterWithoutRecord.GET("getGoodsEvaluationList", goodsEvaluationApi.GetGoodsEvaluationList) // 获取GoodsEvaluation列表
	}
}
