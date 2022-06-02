package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/rss"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type RssFeedSearch struct{
    rss.RssFeed
    request.PageInfo
}
