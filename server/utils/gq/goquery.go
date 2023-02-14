package gq

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
	"io"
	"net/http"

	"github.com/gogf/gf/net/ghttp"

	query "github.com/PuerkitoBio/goquery"
)

// FetchHTML 获取网页
func FetchHTML(url string) *query.Document {
	resp, err := ghttp.NewClient().Get(url)
	if err != nil {
		global.GVA_LOG.Error("function FetchHTML() failed", zap.String("url", url), zap.String("err", err.Error()))
		return &query.Document{}
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			global.GVA_LOG.Error("http close failed", zap.String("url", url), zap.String("err", err.Error()))
		}
	}(resp.Body)

	return DocQuery(resp.Response)
}

// PostHTML 发送表单请求
func PostHTML(url string, m map[string]any) *query.Document {
	resp, err := ghttp.NewClient().Post(url, m)
	if err != nil {
		global.GVA_LOG.Error("function PostHTML() failed", zap.String("url", url), zap.String("err", err.Error()))
		return nil
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			global.GVA_LOG.Error("function PostHTML() failed", zap.String("url", url), zap.String("err", err.Error()))
		}
	}(resp.Response.Body)

	return DocQuery(resp.Response)
}

// 请求goquery
func DocQuery(resp *http.Response) *query.Document {
	doc, err := query.NewDocumentFromReader(resp.Body)
	if err != nil {
		global.GVA_LOG.Error("function DocQuery() failed", zap.String("url", resp.Request.URL.String()), zap.String("err", err.Error()))
		return &query.Document{}
	}

	return doc
}
