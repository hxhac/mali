package rss

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/rss"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    rssReq "github.com/flipped-aurora/gin-vue-admin/server/model/rss/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type RssCategoryApi struct {
}

var rssCategoryService = service.ServiceGroupApp.RssServiceGroup.RssCategoryService


// CreateRssCategory 创建RssCategory
// @Tags RssCategory
// @Summary 创建RssCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body rss.RssCategory true "创建RssCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rssCategory/createRssCategory [post]
func (rssCategoryApi *RssCategoryApi) CreateRssCategory(c *gin.Context) {
	var rssCategory rss.RssCategory
	_ = c.ShouldBindJSON(&rssCategory)
	if err := rssCategoryService.CreateRssCategory(rssCategory); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRssCategory 删除RssCategory
// @Tags RssCategory
// @Summary 删除RssCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body rss.RssCategory true "删除RssCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rssCategory/deleteRssCategory [delete]
func (rssCategoryApi *RssCategoryApi) DeleteRssCategory(c *gin.Context) {
	var rssCategory rss.RssCategory
	_ = c.ShouldBindJSON(&rssCategory)
	if err := rssCategoryService.DeleteRssCategory(rssCategory); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRssCategoryByIds 批量删除RssCategory
// @Tags RssCategory
// @Summary 批量删除RssCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RssCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /rssCategory/deleteRssCategoryByIds [delete]
func (rssCategoryApi *RssCategoryApi) DeleteRssCategoryByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := rssCategoryService.DeleteRssCategoryByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRssCategory 更新RssCategory
// @Tags RssCategory
// @Summary 更新RssCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body rss.RssCategory true "更新RssCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rssCategory/updateRssCategory [put]
func (rssCategoryApi *RssCategoryApi) UpdateRssCategory(c *gin.Context) {
	var rssCategory rss.RssCategory
	_ = c.ShouldBindJSON(&rssCategory)
	if err := rssCategoryService.UpdateRssCategory(rssCategory); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRssCategory 用id查询RssCategory
// @Tags RssCategory
// @Summary 用id查询RssCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query rss.RssCategory true "用id查询RssCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /rssCategory/findRssCategory [get]
func (rssCategoryApi *RssCategoryApi) FindRssCategory(c *gin.Context) {
	var rssCategory rss.RssCategory
	_ = c.ShouldBindQuery(&rssCategory)
	if err, rerssCategory := rssCategoryService.GetRssCategory(rssCategory.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerssCategory": rerssCategory}, c)
	}
}

// GetRssCategoryList 分页获取RssCategory列表
// @Tags RssCategory
// @Summary 分页获取RssCategory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query rssReq.RssCategorySearch true "分页获取RssCategory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rssCategory/getRssCategoryList [get]
func (rssCategoryApi *RssCategoryApi) GetRssCategoryList(c *gin.Context) {
	var pageInfo rssReq.RssCategorySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := rssCategoryService.GetRssCategoryInfoList(pageInfo); err != nil {
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
