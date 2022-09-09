package rss

import (
	"fmt"
	"log"
	"sort"

	"github.com/golang-module/carbon"
	"gorm.io/gorm"

	resp "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	rssModel "github.com/flipped-aurora/gin-vue-admin/server/model/rss"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/helper/time"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/rss"
	"github.com/gin-gonic/gin"
	"github.com/mmcdole/gofeed"
)

const (
	DefaultAuthor    = ""
	FeedLimitPerFeed = 99
	TimeoutSeconds   = 60
	TimeLimit        = 30 // 默认晚上8点后30min
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
		UpdatedTime: time.GetToday(),
	}

	// 如果禁止，则停止更新
	// 判断当前是否处于更新时间范围内
	// 不需要单独处理IsUpdate，因为默认有内容（只需要处理nil的情况）
	isCron := isRssCategoryCron(r.Cron, carbon.Now().IsFriday(), r.UpdateTimeStub)
	if *r.IsUpdate || isCron && !*r.IsMute {
		// 正常更新的情况
		_, urls := rssCategoryService.GetRssURLs(r.Uuid)
		allFeeds := fetchUrls(urls)

		res := rss.Rss(feed, feeds(allFeeds))
		resp.SendXML(ctx, res)
		return
	}

	res := rss.Rss(feed, nil)
	resp.SendXML(ctx, res)
}

func feeds(allFeeds []*gofeed.Feed) []rss.Item {
	// 根据发布时间排序
	// TODO 无法处理没有pubDate参数的rss源
	sort.Sort(sort.Reverse(byPublished(allFeeds)))
	seen := make(map[string]bool)
	ret := []rss.Item{}

	for _, sourceFeed := range allFeeds {

		// 通过feed的源url匹配，并修改"最后更新时间"
		var rssFeed rssModel.RssFeed
		rssFeed.SourceUrl = sourceFeed.Link
		rssFeed.LastUpdated = gorm.DeletedAt{Time: *sourceFeed.UpdatedParsed, Valid: true}
		err := rssFeedService.UpdateUpdatedTime(rssFeed)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		for _, item := range sourceFeed.Items {
			// 判断title是否命中关键字
			if seen[item.Link] {
				continue
			}
			created := item.PublishedParsed
			if created == nil {
				created = item.UpdatedParsed
			}
			ret = append(ret, rss.Item{
				Title:       item.Title,
				URL:         item.Link,
				Description: item.Description,
				Author:      getAuthor(item.Author),
				Contents:    item.Content,
				UpdatedTime: *created,
				ID:          item.GUID,
			})
			seen[item.Link] = true
		}
	}
	return ret
}

// isRssCategoryCron
func isRssCategoryCron(cronTime string, isWeekDay bool, nn string) bool {
	checkCron := CheckCronNow(cronTime, isWeekDay)
	checkTime := time.CheckTimeLimit(nn)

	return checkCron && checkTime
}

func fetchUrl(url string, ch chan<- *gofeed.Feed) {
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

func fetchUrls(urls []string) []*gofeed.Feed {
	allFeeds := make([]*gofeed.Feed, 0)
	ch := make(chan *gofeed.Feed)
	for _, url := range urls {
		go fetchUrl(url, ch)
	}
	for range urls {
		feed := <-ch
		if feed != nil {
			allFeeds = append(allFeeds, feed)
		}
	}
	return allFeeds
}

// TODO: there must be a shorter syntax for this
type byPublished []*gofeed.Feed

func (s byPublished) Len() int {
	return len(s)
}

func (s byPublished) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byPublished) Less(i, j int) bool {
	date1 := s[i].Items[0].PublishedParsed
	if date1 == nil {
		date1 = s[i].Items[0].UpdatedParsed
	}
	date2 := s[j].Items[0].PublishedParsed
	if date2 == nil {
		date2 = s[j].Items[0].UpdatedParsed
	}
	return date1.Before(*date2)
}

// 获取item的author
// TODO
func getAuthor(author *gofeed.Person) string {
	if author != nil {
		return author.Name
	}
	return DefaultAuthor
}
