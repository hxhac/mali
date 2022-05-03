package rss

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/utils/helper/html"

	"github.com/flipped-aurora/gin-vue-admin/server/utils/helper/time"

	resp "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/rss"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/os/gtime"
)

type RssApi struct{}

var lifeEverydayService = service.ServiceGroupApp.DailyServiceGroup.LifeEverydayService

func (RssApi) HabitEverydayRss(ctx *gin.Context) {
	res := rss.Rss(&rss.Feed{
		Title: rss.Title{
			Prefix: "life",
			Name:   "每日习惯everyday",
		},
		Author:      "lry",
		URL:         GetURL(ctx.Request),
		UpdatedTime: time.GetToday(),
	}, everyday())

	resp.SendXML(ctx, res)
}

func everyday() []rss.Item {
	ret := []rss.Item{}
	items, _ := lifeEverydayService.FindAll()

	for _, item := range items {
		dateTime := time.CheckDateTime(item.TimeStub)
		formatTime := dateTime.Format("H:i")
		prefix := item.Prefix

		if time.CheckDateTime(item.TimeStub).Before(gtime.Now()) {
			ret = append(ret, rss.Item{
				Title:       fmt.Sprintf("(从%s%s开始)%s", prefix, formatTime, item.Task),
				Contents:    fmt.Sprintf("%s\n%s", html.Md2HTML(item.Remark), html.Md2HTML(item.More)),
				UpdatedTime: dateTime.Time,
				ID:          rss.GenDateID("habit-routine", item.Task),
			})
		}
	}
	return ret
}
