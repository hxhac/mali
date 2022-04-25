package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/daily"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/goods"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Daily   daily.RouterGroup
	Goods   goods.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
