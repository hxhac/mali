package goods

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GoodsLabelRouter struct {
}

// InitGoodsLabelRouter 初始化 GoodsLabel 路由信息
func (s *GoodsLabelRouter) InitGoodsLabelRouter(Router *gin.RouterGroup) {
	goodsLabelRouter := Router.Group("goodsLabel").Use(middleware.OperationRecord())
	goodsLabelRouterWithoutRecord := Router.Group("goodsLabel")
	var goodsLabelApi = v1.ApiGroupApp.GoodsApiGroup.GoodsLabelApi
	{
		goodsLabelRouter.POST("createGoodsLabel", goodsLabelApi.CreateGoodsLabel)   // 新建GoodsLabel
		goodsLabelRouter.DELETE("deleteGoodsLabel", goodsLabelApi.DeleteGoodsLabel) // 删除GoodsLabel
		goodsLabelRouter.DELETE("deleteGoodsLabelByIds", goodsLabelApi.DeleteGoodsLabelByIds) // 批量删除GoodsLabel
		goodsLabelRouter.PUT("updateGoodsLabel", goodsLabelApi.UpdateGoodsLabel)    // 更新GoodsLabel
	}
	{
		goodsLabelRouterWithoutRecord.GET("findGoodsLabel", goodsLabelApi.FindGoodsLabel)        // 根据ID获取GoodsLabel
		goodsLabelRouterWithoutRecord.GET("getGoodsLabelList", goodsLabelApi.GetGoodsLabelList)  // 获取GoodsLabel列表
	}
}
