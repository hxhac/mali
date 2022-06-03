package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/rss"
)

type RssCategorySearch struct {
	rss.RssCategory
	request.PageInfo
}
