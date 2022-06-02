package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/goods"
	"github.com/flipped-aurora/gin-vue-admin/server/router/life"
	"github.com/flipped-aurora/gin-vue-admin/server/router/rss"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Daily   life.RouterGroup
	Goods   goods.RouterGroup
	Life    life.RouterGroup
	Rss     rss.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
