package daily

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/daily"
	dailyReq "github.com/flipped-aurora/gin-vue-admin/server/model/daily/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LifeEverydayApi struct {
}

var lifeEverydayService = service.ServiceGroupApp.DailyServiceGroup.LifeEverydayService

// CreateLifeEveryday 创建LifeEveryday
// @Tags LifeEveryday
// @Summary 创建LifeEveryday
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body daily.LifeEveryday true "创建LifeEveryday"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lifeEveryday/createLifeEveryday [post]
func (lifeEverydayApi *LifeEverydayApi) CreateLifeEveryday(c *gin.Context) {
	var lifeEveryday daily.LifeEveryday
	_ = c.ShouldBindJSON(&lifeEveryday)
	if err := lifeEverydayService.CreateLifeEveryday(lifeEveryday); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLifeEveryday 删除LifeEveryday
// @Tags LifeEveryday
// @Summary 删除LifeEveryday
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body daily.LifeEveryday true "删除LifeEveryday"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lifeEveryday/deleteLifeEveryday [delete]
func (lifeEverydayApi *LifeEverydayApi) DeleteLifeEveryday(c *gin.Context) {
	var lifeEveryday daily.LifeEveryday
	_ = c.ShouldBindJSON(&lifeEveryday)
	if err := lifeEverydayService.DeleteLifeEveryday(lifeEveryday); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteLifeEverydayByIds 批量删除LifeEveryday
// @Tags LifeEveryday
// @Summary 批量删除LifeEveryday
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LifeEveryday"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /lifeEveryday/deleteLifeEverydayByIds [delete]
func (lifeEverydayApi *LifeEverydayApi) DeleteLifeEverydayByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := lifeEverydayService.DeleteLifeEverydayByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateLifeEveryday 更新LifeEveryday
// @Tags LifeEveryday
// @Summary 更新LifeEveryday
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body daily.LifeEveryday true "更新LifeEveryday"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lifeEveryday/updateLifeEveryday [put]
func (lifeEverydayApi *LifeEverydayApi) UpdateLifeEveryday(c *gin.Context) {
	var lifeEveryday daily.LifeEveryday
	_ = c.ShouldBindJSON(&lifeEveryday)
	if err := lifeEverydayService.UpdateLifeEveryday(lifeEveryday); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLifeEveryday 用id查询LifeEveryday
// @Tags LifeEveryday
// @Summary 用id查询LifeEveryday
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query daily.LifeEveryday true "用id查询LifeEveryday"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lifeEveryday/findLifeEveryday [get]
func (lifeEverydayApi *LifeEverydayApi) FindLifeEveryday(c *gin.Context) {
	var lifeEveryday daily.LifeEveryday
	_ = c.ShouldBindQuery(&lifeEveryday)
	if err, relifeEveryday := lifeEverydayService.GetLifeEveryday(lifeEveryday.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relifeEveryday": relifeEveryday}, c)
	}
}

// GetLifeEverydayList 分页获取LifeEveryday列表
// @Tags LifeEveryday
// @Summary 分页获取LifeEveryday列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dailyReq.LifeEverydaySearch true "分页获取LifeEveryday列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lifeEveryday/getLifeEverydayList [get]
func (lifeEverydayApi *LifeEverydayApi) GetLifeEverydayList(c *gin.Context) {
	var pageInfo dailyReq.LifeEverydaySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := lifeEverydayService.GetLifeEverydayInfoList(pageInfo); err != nil {
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
