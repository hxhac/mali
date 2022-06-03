package rss

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RssFeedRouter struct{}

// InitRssFeedRouter 初始化 RssFeed 路由信息
func (s *RssFeedRouter) InitRssFeedRouter(Router *gin.RouterGroup) {
	rssFeedRouter := Router.Group("rssFeed").Use(middleware.OperationRecord())
	rssFeedRouterWithoutRecord := Router.Group("rssFeed")
	rssFeedApi := v1.ApiGroupApp.RssApiGroup.RssFeedApi
	{
		rssFeedRouter.POST("createRssFeed", rssFeedApi.CreateRssFeed)             // 新建RssFeed
		rssFeedRouter.DELETE("deleteRssFeed", rssFeedApi.DeleteRssFeed)           // 删除RssFeed
		rssFeedRouter.DELETE("deleteRssFeedByIds", rssFeedApi.DeleteRssFeedByIds) // 批量删除RssFeed
		rssFeedRouter.PUT("updateRssFeed", rssFeedApi.UpdateRssFeed)              // 更新RssFeed
	}
	{
		rssFeedRouterWithoutRecord.GET("findRssFeed", rssFeedApi.FindRssFeed)       // 根据ID获取RssFeed
		rssFeedRouterWithoutRecord.GET("getRssFeedList", rssFeedApi.GetRssFeedList) // 获取RssFeed列表
	}
}
