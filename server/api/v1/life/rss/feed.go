package rss

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	resp "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	rssModel "github.com/flipped-aurora/gin-vue-admin/server/model/rss"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	ht "github.com/flipped-aurora/gin-vue-admin/server/utils/helper/time"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/rss"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon"
	"github.com/mmcdole/gofeed"
	"go.uber.org/zap"
	"log"
	"time"
)

const (
	DefaultAuthor = "jeffcott"
)

var (
	rssCategoryService = service.ServiceGroupApp.RssServiceGroup.RssCategoryService
	rssFeedService     = service.ServiceGroupApp.RssServiceGroup.RssFeedService
)

func (RssApi) FeedRss(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	err, r := rssCategoryService.GetRssCategoryByUUID(uuid)
	if err != nil {
		return
	}

	feed := &rss.Feed{
		Title: rss.Title{
			Prefix: "rss",
			Name:   r.CateName,
		},
		Author:      r.Author,
		URL:         GetURL(ctx.Request),
		UpdatedTime: ht.GetToday(),
	}

	// 如果禁止，则停止更新
	// 判断当前是否处于更新时间范围内
	// 不需要单独处理IsUpdate，因为默认有内容（只需要处理nil的情况）
	isCron := isRssCategoryCron(r.Cron, carbon.Now().IsFriday(), r.UpdateTimeStub)
	if *r.IsUpdate || isCron && !*r.IsMute {
		// 正常更新的情况
		_, urls := rssCategoryService.GetRssURLs(r.Uuid)
		allFeeds := fetchURLs(urls)

		res := rss.Rss(feed, feeds(allFeeds))
		resp.SendXML(ctx, res)
		return
	}

	res := rss.Rss(feed, nil)
	resp.SendXML(ctx, res)
}

func feeds(allFeeds []*gofeed.Feed) (ret []*rss.Item) {
	// 移除所有没有时间戳字段的feed
	filteredFeeds := allFeeds[:0]
	for _, feed := range allFeeds {
		if feed.PublishedParsed != nil || feed.UpdatedParsed != nil {
			filteredFeeds = append(filteredFeeds, feed)
		} else {
			global.GVA_LOG.Error("function feeds() Failed, missing params published_time, cause error", zap.String("url", feed.FeedLink))
		}
	}
	// 根据发布时间排序
	//bp := byPublished(filteredFeeds)
	//sort.Sort(sort.Reverse(bp))
	//seen := make(map[string]bool)

	for _, sourceFeed := range allFeeds {

		// 通过feed的源url匹配，并修改"最后更新时间"
		var rssFeed rssModel.RssFeed
		rssFeed.SourceUrl = sourceFeed.Link
		// 判断updated_time是否存在
		if sourceFeed.UpdatedParsed != nil {
			rssFeed.UpdatedAt = *sourceFeed.UpdatedParsed
			err := rssFeedService.UpdateUpdatedTime(rssFeed)
			global.GVA_LOG.Error("function Feeds() failed", zap.String("url", rssFeed.SourceUrl), zap.String("err", err.Error()))
		}

		for _, item := range sourceFeed.Items {
			//if seen[item.Link] {
			//	continue
			//}
			// 判断
			created := ht.GetToday()
			if item.UpdatedParsed != nil {
				created = *item.UpdatedParsed
			}
			if item.PublishedParsed != nil {
				created = *item.PublishedParsed
			}
			if item.UpdatedParsed == nil && item.PublishedParsed == nil {
				created = time.Now()
			}
			ret = append(ret, &rss.Item{
				Title:       item.Title,
				URL:         item.Link,
				Description: item.Description,
				Author:      getAuthor(item.Author),
				Contents:    item.Content,
				UpdatedTime: created,
				ID:          item.GUID,
			})
			//seen[item.Link] = true
		}
	}
	return
}

// isRssCategoryCron
func isRssCategoryCron(cronTime string, isWeekDay bool, nn string) bool {
	checkCron := CheckCronNow(cronTime, isWeekDay)
	checkTime := ht.CheckTimeLimit(nn)

	return checkCron && checkTime
}

// fetchURL
func fetchURL(url string, ch chan<- *gofeed.Feed) {
	log.Printf("Fetching URL: %v\n", url)
	fp := gofeed.NewParser()

	feed, err := fp.ParseURL(url)
	if err == nil {
		ch <- feed
	} else {
		log.Printf("Error on URL: %v (%v)", url, err)
		ch <- nil
	}
}

// fetchURLs
func fetchURLs(urls []string) []*gofeed.Feed {
	allFeeds := make([]*gofeed.Feed, 0)
	ch := make(chan *gofeed.Feed)
	for _, url := range urls {
		go fetchURL(url, ch)
	}
	for range urls {
		feed := <-ch
		if feed != nil {
			allFeeds = append(allFeeds, feed)
		}
	}
	return allFeeds
}

//type byPublished []*gofeed.Feed

//func (s byPublished) Len() int {
//	return len(s)
//}
//
//func (s byPublished) Swap(i, j int) {
//	s[i], s[j] = s[j], s[i]
//}
//
//func (s byPublished) Less(i, j int) bool {
//	date1 := s[i].Items[0].PublishedParsed
//	if date1 == nil {
//		date1 = s[i].Items[0].UpdatedParsed
//	}
//	date2 := s[j].Items[0].PublishedParsed
//	if date2 == nil {
//		date2 = s[j].Items[0].UpdatedParsed
//	}
//	return date1.Before(*date2)
//}

// 获取item的author
// TODO
func getAuthor(author *gofeed.Person) string {
	if author != nil {
		return author.Name
	}
	return DefaultAuthor
}
