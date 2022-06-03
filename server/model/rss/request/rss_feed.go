package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/rss"
)

type RssFeedSearch struct {
	rss.RssFeed
	request.PageInfo
}
