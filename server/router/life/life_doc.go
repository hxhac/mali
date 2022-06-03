package life

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LifeDocRouter struct{}

// InitLifeDocRouter 初始化 LifeDoc 路由信息
func (s *LifeDocRouter) InitLifeDocRouter(Router *gin.RouterGroup) {
	lifeDocRouter := Router.Group("lifeDoc").Use(middleware.OperationRecord())
	lifeDocRouterWithoutRecord := Router.Group("lifeDoc")
	lifeDocApi := v1.ApiGroupApp.LifeApiGroup.LifeDocApi
	{
		lifeDocRouter.POST("createLifeDoc", lifeDocApi.CreateLifeDoc)             // 新建LifeDoc
		lifeDocRouter.DELETE("deleteLifeDoc", lifeDocApi.DeleteLifeDoc)           // 删除LifeDoc
		lifeDocRouter.DELETE("deleteLifeDocByIds", lifeDocApi.DeleteLifeDocByIds) // 批量删除LifeDoc
		lifeDocRouter.PUT("updateLifeDoc", lifeDocApi.UpdateLifeDoc)              // 更新LifeDoc
	}
	{
		lifeDocRouterWithoutRecord.GET("findLifeDoc", lifeDocApi.FindLifeDoc)       // 根据ID获取LifeDoc
		lifeDocRouterWithoutRecord.GET("getLifeDocList", lifeDocApi.GetLifeDocList) // 获取LifeDoc列表
	}
}
