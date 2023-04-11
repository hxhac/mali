package app

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppManageRouter struct {
}

// InitAppManageRouter 初始化 AppManage 路由信息
func (s *AppManageRouter) InitAppManageRouter(Router *gin.RouterGroup) {
	appManageRouter := Router.Group("appManage").Use(middleware.OperationRecord())
	appManageRouterWithoutRecord := Router.Group("appManage")
	var appManageApi = v1.ApiGroupApp.AppApiGroup.AppManageApi
	{
		appManageRouter.POST("createAppManage", appManageApi.CreateAppManage)             // 新建AppManage
		appManageRouter.DELETE("deleteAppManage", appManageApi.DeleteAppManage)           // 删除AppManage
		appManageRouter.DELETE("deleteAppManageByIds", appManageApi.DeleteAppManageByIds) // 批量删除AppManage
		appManageRouter.PUT("updateAppManage", appManageApi.UpdateAppManage)              // 更新AppManage
	}
	{
		appManageRouterWithoutRecord.GET("findAppManage", appManageApi.FindAppManage)       // 根据ID获取AppManage
		appManageRouterWithoutRecord.GET("getAppManageList", appManageApi.GetAppManageList) // 获取AppManage列表
	}
}
