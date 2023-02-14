package gq

import (
	"github.com/chanyipiaomiao/hlog"
	"io"
	"net/http"

	"github.com/gogf/gf/net/ghttp"

	query "github.com/PuerkitoBio/goquery"
)

// FetchHTML 获取网页
func FetchHTML(url string) *query.Document {
	resp, err := ghttp.NewClient().Get(url)
	if err != nil {
		hlog.Error(hlog.D{"url": url}, "http request failed")
		return &query.Document{}
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			hlog.Error(hlog.D{"url": url}, "http close failed")
		}
	}(resp.Body)

	return DocQuery(resp.Response)
}

// PostHTML 发送表单请求
func PostHTML(url string, m map[string]any) *query.Document {
	resp, err := ghttp.NewClient().Post(url, m)
	if err != nil {
		return nil
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			hlog.Error(hlog.D{"url": url}, "http close failed")
		}
	}(resp.Response.Body)

	return DocQuery(resp.Response)
}

// 请求goquery
func DocQuery(resp *http.Response) *query.Document {
	doc, err := query.NewDocumentFromReader(resp.Body)
	if err != nil {
		hlog.Error(hlog.D{"url": resp.Request.URL.String()}, "goquery failed")
		return &query.Document{}
	}

	return doc
}
