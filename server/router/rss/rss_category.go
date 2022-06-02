package rss

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RssCategoryRouter struct {
}

// InitRssCategoryRouter 初始化 RssCategory 路由信息
func (s *RssCategoryRouter) InitRssCategoryRouter(Router *gin.RouterGroup) {
	rssCategoryRouter := Router.Group("rssCategory").Use(middleware.OperationRecord())
	rssCategoryRouterWithoutRecord := Router.Group("rssCategory")
	var rssCategoryApi = v1.ApiGroupApp.RssApiGroup.RssCategoryApi
	{
		rssCategoryRouter.POST("createRssCategory", rssCategoryApi.CreateRssCategory)   // 新建RssCategory
		rssCategoryRouter.DELETE("deleteRssCategory", rssCategoryApi.DeleteRssCategory) // 删除RssCategory
		rssCategoryRouter.DELETE("deleteRssCategoryByIds", rssCategoryApi.DeleteRssCategoryByIds) // 批量删除RssCategory
		rssCategoryRouter.PUT("updateRssCategory", rssCategoryApi.UpdateRssCategory)    // 更新RssCategory
	}
	{
		rssCategoryRouterWithoutRecord.GET("findRssCategory", rssCategoryApi.FindRssCategory)        // 根据ID获取RssCategory
		rssCategoryRouterWithoutRecord.GET("getRssCategoryList", rssCategoryApi.GetRssCategoryList)  // 获取RssCategory列表
	}
}
