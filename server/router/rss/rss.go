package rss

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type RssRouter struct{}

func (s *RssRouter) InitRssRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	rssRouter := Router.Group("rss")
	baseApi := v1.ApiGroupApp.SystemApiGroup.RssApi
	{
		rssRouter.GET("yearly", baseApi.HabitYearlyRss)
		rssRouter.GET("everyday", baseApi.HabitEverydayRss)
	}
	return rssRouter
}
