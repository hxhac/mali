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
	geRouter := Router.Group("ge").Use(middleware.OperationRecord())
	geRouterWithoutRecord := Router.Group("ge")
	var geApi = v1.ApiGroupApp.GoodsApiGroup.GoodsEvaluationApi
	{
		geRouter.POST("createGoodsEvaluation", geApi.CreateGoodsEvaluation)   // 新建GoodsEvaluation
		geRouter.DELETE("deleteGoodsEvaluation", geApi.DeleteGoodsEvaluation) // 删除GoodsEvaluation
		geRouter.DELETE("deleteGoodsEvaluationByIds", geApi.DeleteGoodsEvaluationByIds) // 批量删除GoodsEvaluation
		geRouter.PUT("updateGoodsEvaluation", geApi.UpdateGoodsEvaluation)    // 更新GoodsEvaluation
	}
	{
		geRouterWithoutRecord.GET("findGoodsEvaluation", geApi.FindGoodsEvaluation)        // 根据ID获取GoodsEvaluation
		geRouterWithoutRecord.GET("getGoodsEvaluationList", geApi.GetGoodsEvaluationList)  // 获取GoodsEvaluation列表
	}
}
