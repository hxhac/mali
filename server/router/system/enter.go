package system

import "github.com/flipped-aurora/gin-vue-admin/server/router/rss"

type RouterGroup struct {
	ApiRouter
	JwtRouter
	SysRouter
	BaseRouter
	rss.RssRouter
	InitRouter
	MenuRouter
	UserRouter
	CasbinRouter
	AutoCodeRouter
	AuthorityRouter
	DictionaryRouter
	OperationRecordRouter
	DictionaryDetailRouter
	AuthorityBtnRouter
}
