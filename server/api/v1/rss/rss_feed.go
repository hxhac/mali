package rss

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/rss"
	rssReq "github.com/flipped-aurora/gin-vue-admin/server/model/rss/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"github.com/mmcdole/gofeed"
	"go.uber.org/zap"
)

type RssFeedApi struct{}

var rssFeedService = service.ServiceGroupApp.RssServiceGroup.RssFeedService

// CreateRssFeed 创建RssFeed
// @Tags RssFeed
// @Summary 创建RssFeed
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body rss.RssFeed true "创建RssFeed"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rssFeed/createRssFeed [post]
func (rssFeedApi *RssFeedApi) CreateRssFeed(c *gin.Context) {
	var rssFeed rss.RssFeed
	_ = c.ShouldBindJSON(&rssFeed)

	// rss名称等数据直接通过解析rss获得
	// 验证feed是否合法
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(rssFeed.Url)
	if err != nil {
		response.FailWithMessage("feed非法，创建失败", c)
		return
	}
	rssFeed.RssName = feed.Title
	rssFeed.SourceUrl = feed.Link

	if err := rssFeedService.CreateRssFeed(rssFeed); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRssFeed 删除RssFeed
// @Tags RssFeed
// @Summary 删除RssFeed
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body rss.RssFeed true "删除RssFeed"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rssFeed/deleteRssFeed [delete]
func (rssFeedApi *RssFeedApi) DeleteRssFeed(c *gin.Context) {
	var rssFeed rss.RssFeed
	_ = c.ShouldBindJSON(&rssFeed)
	if err := rssFeedService.DeleteRssFeed(rssFeed); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRssFeedByIds 批量删除RssFeed
// @Tags RssFeed
// @Summary 批量删除RssFeed
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RssFeed"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /rssFeed/deleteRssFeedByIds [delete]
func (rssFeedApi *RssFeedApi) DeleteRssFeedByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := rssFeedService.DeleteRssFeedByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRssFeed 更新RssFeed
// @Tags RssFeed
// @Summary 更新RssFeed
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body rss.RssFeed true "更新RssFeed"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rssFeed/updateRssFeed [put]
func (rssFeedApi *RssFeedApi) UpdateRssFeed(c *gin.Context) {
	var rssFeed rss.RssFeed
	_ = c.ShouldBindJSON(&rssFeed)
	if err := rssFeedService.UpdateRssFeed(rssFeed); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRssFeed 用id查询RssFeed
// @Tags RssFeed
// @Summary 用id查询RssFeed
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query rss.RssFeed true "用id查询RssFeed"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /rssFeed/findRssFeed [get]
func (rssFeedApi *RssFeedApi) FindRssFeed(c *gin.Context) {
	var rssFeed rss.RssFeed
	_ = c.ShouldBindQuery(&rssFeed)
	if err, rerssFeed := rssFeedService.GetRssFeed(rssFeed.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerssFeed": rerssFeed}, c)
	}
}

// GetRssFeedList 分页获取RssFeed列表
// @Tags RssFeed
// @Summary 分页获取RssFeed列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query rssReq.RssFeedSearch true "分页获取RssFeed列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rssFeed/getRssFeedList [get]
func (rssFeedApi *RssFeedApi) GetRssFeedList(c *gin.Context) {
	var pageInfo rssReq.RssFeedSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := rssFeedService.GetRssFeedInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
