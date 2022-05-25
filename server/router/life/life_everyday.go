package life

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LifeEverydayRouter struct{}

// InitLifeEverydayRouter 初始化 LifeEveryday 路由信息
func (s *LifeEverydayRouter) InitLifeEverydayRouter(Router *gin.RouterGroup) {
	lifeEverydayRouter := Router.Group("lifeEveryday").Use(middleware.OperationRecord())
	lifeEverydayRouterWithoutRecord := Router.Group("lifeEveryday")
	lifeEverydayApi := v1.ApiGroupApp.DailyApiGroup.LifeEverydayApi
	{
		lifeEverydayRouter.POST("createLifeEveryday", lifeEverydayApi.CreateLifeEveryday)             // 新建LifeEveryday
		lifeEverydayRouter.DELETE("deleteLifeEveryday", lifeEverydayApi.DeleteLifeEveryday)           // 删除LifeEveryday
		lifeEverydayRouter.DELETE("deleteLifeEverydayByIds", lifeEverydayApi.DeleteLifeEverydayByIds) // 批量删除LifeEveryday
		lifeEverydayRouter.PUT("updateLifeEveryday", lifeEverydayApi.UpdateLifeEveryday)              // 更新LifeEveryday
	}
	{
		lifeEverydayRouterWithoutRecord.GET("findLifeEveryday", lifeEverydayApi.FindLifeEveryday)       // 根据ID获取LifeEveryday
		lifeEverydayRouterWithoutRecord.GET("getLifeEverydayList", lifeEverydayApi.GetLifeEverydayList) // 获取LifeEveryday列表
	}
}
