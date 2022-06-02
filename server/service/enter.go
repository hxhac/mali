package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/goods"
	"github.com/flipped-aurora/gin-vue-admin/server/service/life"
	"github.com/flipped-aurora/gin-vue-admin/server/service/rss"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	DailyServiceGroup   life.ServiceGroup
	GoodsServiceGroup   goods.ServiceGroup
	LifeServiceGroup    life.ServiceGroup
	RssServiceGroup     rss.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
