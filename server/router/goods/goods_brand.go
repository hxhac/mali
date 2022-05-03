package goods

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GoodsBrandRouter struct{}

// InitGoodsBrandRouter 初始化 GoodsBrand 路由信息
func (s *GoodsBrandRouter) InitGoodsBrandRouter(Router *gin.RouterGroup) {
	goodsBrandRouter := Router.Group("goodsBrand").Use(middleware.OperationRecord())
	goodsBrandRouterWithoutRecord := Router.Group("goodsBrand")
	goodsBrandApi := v1.ApiGroupApp.GoodsApiGroup.GoodsBrandApi
	{
		goodsBrandRouter.POST("createGoodsBrand", goodsBrandApi.CreateGoodsBrand)             // 新建GoodsBrand
		goodsBrandRouter.DELETE("deleteGoodsBrand", goodsBrandApi.DeleteGoodsBrand)           // 删除GoodsBrand
		goodsBrandRouter.DELETE("deleteGoodsBrandByIds", goodsBrandApi.DeleteGoodsBrandByIds) // 批量删除GoodsBrand
		goodsBrandRouter.PUT("updateGoodsBrand", goodsBrandApi.UpdateGoodsBrand)              // 更新GoodsBrand
	}
	{
		goodsBrandRouterWithoutRecord.GET("findGoodsBrand", goodsBrandApi.FindGoodsBrand)       // 根据ID获取GoodsBrand
		goodsBrandRouterWithoutRecord.GET("getGoodsBrandList", goodsBrandApi.GetGoodsBrandList) // 获取GoodsBrand列表
	}
}
