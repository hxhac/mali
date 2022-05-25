package life

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LifeYearlyRouter struct{}

// InitLifeYearlyRouter 初始化 LifeYearly 路由信息
func (s *LifeYearlyRouter) InitLifeYearlyRouter(Router *gin.RouterGroup) {
	lifeYearlyRouter := Router.Group("lifeYearly").Use(middleware.OperationRecord())
	lifeYearlyRouterWithoutRecord := Router.Group("lifeYearly")
	lifeYearlyApi := v1.ApiGroupApp.DailyApiGroup.LifeYearlyApi
	{
		lifeYearlyRouter.POST("createLifeYearly", lifeYearlyApi.CreateLifeYearly)             // 新建LifeYearly
		lifeYearlyRouter.DELETE("deleteLifeYearly", lifeYearlyApi.DeleteLifeYearly)           // 删除LifeYearly
		lifeYearlyRouter.DELETE("deleteLifeYearlyByIds", lifeYearlyApi.DeleteLifeYearlyByIds) // 批量删除LifeYearly
		lifeYearlyRouter.PUT("updateLifeYearly", lifeYearlyApi.UpdateLifeYearly)              // 更新LifeYearly
	}
	{
		lifeYearlyRouterWithoutRecord.GET("findLifeYearly", lifeYearlyApi.FindLifeYearly)             // 根据ID获取LifeYearly
		lifeYearlyRouterWithoutRecord.GET("getLifeYearlyList", lifeYearlyApi.GetLifeYearlyList)       // 获取LifeYearly列表
		lifeYearlyRouterWithoutRecord.GET("getLifeYearlyOptions", lifeYearlyApi.GetLifeYearlyOptions) // 获取LifeYearly列表
	}
}
