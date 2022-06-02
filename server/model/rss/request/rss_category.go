package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/rss"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type RssCategorySearch struct{
    rss.RssCategory
    request.PageInfo
}
