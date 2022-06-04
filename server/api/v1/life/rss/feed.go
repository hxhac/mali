package rss

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	resp "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/feeds"
	"github.com/mmcdole/gofeed"
)

const (
	DefaultAuthor    = "jf"
	FeedLimitPerFeed = 99
	TimeoutSeconds   = 60
)

var rssCategoryService = service.ServiceGroupApp.RssServiceGroup.RssCategoryService

// FeedRss xxx
func (RssApi) FeedRss(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	err, r := rssCategoryService.GetRssCategoryByUUID(uuid)
	if err != nil {
		return
	}

	_, urls := rssCategoryService.GetRssURLs(r.Uuid)
	allFeeds := fetchUrls(urls)
	// 合并所有feed
	feed := &feeds.Feed{
		Title:       fmt.Sprintf("%s - %s", "rss", r.CateName),
		Link:        &feeds.Link{Href: GetURL(ctx.Request)},
		Description: r.Remark,
		Author: &feeds.Author{
			Name: r.Author,
		},
		Created: time.Now(),
	}

	// 根据发布时间排序
	sort.Sort(sort.Reverse(byPublished(allFeeds)))
	// limit_per_feed := FeedLimitPerFeed
	seen := make(map[string]bool)
	for _, sourceFeed := range allFeeds {
		// for _, item := range sourceFeed.Items[:limit_per_feed] {
		for _, item := range sourceFeed.Items {
			// 判断title是否命中关键字
			if !strings.Contains(r.Title, "") {
				if seen[item.Link] {
					continue
				}
				created := item.PublishedParsed
				if created == nil {
					created = item.UpdatedParsed
				}
				feed.Items = append(feed.Items, &feeds.Item{
					Title:       item.Title,
					Link:        &feeds.Link{Href: item.Link},
					Description: item.Description,
					Author:      &feeds.Author{Name: getAuthor(sourceFeed)},
					Created:     *created,
					Content:     item.Content,
				})
				seen[item.Link] = true
			}
		}
	}

	res, err := feed.ToAtom()
	if err != nil {
		return
	}
	resp.SendXML(ctx, res)
}

func fetchUrl(url string, ch chan<- *gofeed.Feed) {
	log.Printf("Fetching URL: %v\n", url)
	fp := gofeed.NewParser()
	// fp.Client = &http.Client{
	// 	Timeout: time.Duration(TimeoutSeconds) * time.Second,
	// }
	feed, err := fp.ParseURL(url)
	if err == nil {
		ch <- feed
	} else {
		log.Printf("Error on URL: %v (%v)", url, err)
		ch <- nil
	}
}

//
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
func getAuthor(feed *gofeed.Feed) string {
	// if feed.Authors != nil {
	// 	return feed.Authors[0].Name
	// }
	if feed.Items[0].Authors != nil {
		return feed.Items[0].Authors[0].Name
	}
	log.Printf("Could not determine author for %v", feed.Link)
	return DefaultAuthor
}
