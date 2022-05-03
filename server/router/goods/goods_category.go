package goods

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GoodsCategoryRouter struct{}

// InitGoodsCategoryRouter 初始化 GoodsCategory 路由信息
func (s *GoodsCategoryRouter) InitGoodsCategoryRouter(Router *gin.RouterGroup) {
	goodsCategoryRouter := Router.Group("goodsCategory").Use(middleware.OperationRecord())
	goodsCategoryRouterWithoutRecord := Router.Group("goodsCategory")
	goodsCategoryApi := v1.ApiGroupApp.GoodsApiGroup.GoodsCategoryApi
	{
		goodsCategoryRouter.POST("createGoodsCategory", goodsCategoryApi.CreateGoodsCategory)             // 新建GoodsCategory
		goodsCategoryRouter.DELETE("deleteGoodsCategory", goodsCategoryApi.DeleteGoodsCategory)           // 删除GoodsCategory
		goodsCategoryRouter.DELETE("deleteGoodsCategoryByIds", goodsCategoryApi.DeleteGoodsCategoryByIds) // 批量删除GoodsCategory
		goodsCategoryRouter.PUT("updateGoodsCategory", goodsCategoryApi.UpdateGoodsCategory)              // 更新GoodsCategory
	}
	{
		// 根据ID获取GoodsCategory
		goodsCategoryRouterWithoutRecord.GET("findGoodsCategory", goodsCategoryApi.FindGoodsCategory)
		// 获取GoodsCategory列表
		goodsCategoryRouterWithoutRecord.GET("getGoodsCategoryList", goodsCategoryApi.GetGoodsCategoryList)
	}
}
