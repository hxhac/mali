package rss

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/helper/html"

	"github.com/flipped-aurora/gin-vue-admin/server/utils/helper/iter"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/text/gstr"

	htime "github.com/flipped-aurora/gin-vue-admin/server/utils/helper/time"

	resp "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/rss"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/os/gtime"
	"github.com/golang-module/carbon"
)

var numberReg = regexp.MustCompile(`\d+`)

var lifeYearlyService = service.ServiceGroupApp.DailyServiceGroup.LifeYearlyService

func (RssApi) HabitYearlyRss(ctx *gin.Context) {
	res := rss.Rss(&rss.Feed{
		Title: rss.Title{
			Prefix: "life",
			Name:   "生活习惯yearly",
		},
		Author:      "lry",
		URL:         GetURL(ctx.Request),
		UpdatedTime: htime.GetToday(),
	}, yearly())

	resp.SendXML(ctx, res)
}

func yearly() []*rss.Item {
	ret := []*rss.Item{}
	items, _ := lifeYearlyService.FindAll()

	for _, item := range items {
		if !item.IsPause && CheckCronNowDefault(item.Cron) {
			title := fmt.Sprintf("[%s] - [%s] - %s", item.Prefix, gtime.Date(), item.Task)
			ret = append(ret, &rss.Item{
				Title:       title,
				Contents:    fmt.Sprintf("%s\n%s", html.Md2HTML(item.Remark), html.Md2HTML(item.More)),
				UpdatedTime: htime.GetToday(),
				ID:          rss.GenDateID("habit-notify", item.Task),
			})
		}
	}

	return ret
}

// CheckCronNowDefault
func CheckCronNowDefault(cronTime string) bool {
	return CheckCronNow(cronTime, carbon.Now().IsSaturday())
}

// CheckCronSpecifiedTime
func CheckCronSpecifiedTime(tt time.Time, cronTime string) bool {
	return CheckCron(carbon.Time2Carbon(tt), cronTime, true)
}

// CheckCronNow
func CheckCronNow(cronTime string, isWeekDay bool) bool {
	cb := carbon.Now()
	return CheckCron(cb, cronTime, isWeekDay)
}

// CheckCron
// isWeekDay 是否为周六
func CheckCron(cb carbon.Carbon, cronTime string, isWeekDay bool) bool {
	dayOfYear := cb.DayOfYear()
	dayOfMonth := cb.DayOfMonth()
	weekOfYear := cb.WeekOfYear()
	monthOfYear := cb.MonthOfYear()

	isJanuary := cb.IsJanuary()
	isMatched, number := ExtractTimeNumber(cronTime)

	// 判断daily
	if gstr.Contains(cronTime, "daily") && ((dayOfYear-1)%number == 0 || dayOfYear == 1) {
		return true
	}
	// 判断weekly
	wn := (weekOfYear - 1) % number
	if gstr.Contains(cronTime, "weekly") && isWeekDay && (wn == 0 || !isMatched) {
		return true
	}
	// 判断monthly
	months := GetMonths(number)
	isContains := garray.NewIntArrayFrom(months).Contains(monthOfYear)
	if gstr.Contains(cronTime, "monthly") && dayOfMonth == 1 && isContains {
		return true
	}
	// @yearly
	if gstr.Contains(cronTime, "yearly") && isJanuary && dayOfMonth == 1 {
		return true
	}

	return false
}

func GetURL(r *http.Request) string {
	scheme := "http://"

	if r.TLS != nil {
		scheme = "https://"
	}
	return scheme + r.Host + r.RequestURI
}

// ExtractTimeNumber 正则提取数字
// isMatched
func ExtractTimeNumber(t string) (isMatched bool, number int) {
	isMatched = numberReg.MatchString(t)
	if !isMatched {
		return isMatched, 1
	}
	number, _ = strconv.Atoi(numberReg.FindString(t))
	return
}

// GetMonths 间隔取值
func GetMonths(step int) (res []int) {
	for v := range iter.N(1, carbon.MonthsPerYear+1, step) {
		res = append(res, v)
	}
	return
}
