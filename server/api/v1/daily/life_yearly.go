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

type LifeYearlyApi struct {
}

var lifeYearlyService = service.ServiceGroupApp.DailyServiceGroup.LifeYearlyService

// CreateLifeYearly 创建LifeYearly
// @Tags LifeYearly
// @Summary 创建LifeYearly
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body daily.LifeYearly true "创建LifeYearly"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lifeYearly/createLifeYearly [post]
func (lifeYearlyApi *LifeYearlyApi) CreateLifeYearly(c *gin.Context) {
	var lifeYearly daily.LifeYearly
	_ = c.ShouldBindJSON(&lifeYearly)
	if err := lifeYearlyService.CreateLifeYearly(lifeYearly); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLifeYearly 删除LifeYearly
// @Tags LifeYearly
// @Summary 删除LifeYearly
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body daily.LifeYearly true "删除LifeYearly"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lifeYearly/deleteLifeYearly [delete]
func (lifeYearlyApi *LifeYearlyApi) DeleteLifeYearly(c *gin.Context) {
	var lifeYearly daily.LifeYearly
	_ = c.ShouldBindJSON(&lifeYearly)
	if err := lifeYearlyService.DeleteLifeYearly(lifeYearly); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteLifeYearlyByIds 批量删除LifeYearly
// @Tags LifeYearly
// @Summary 批量删除LifeYearly
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LifeYearly"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /lifeYearly/deleteLifeYearlyByIds [delete]
func (lifeYearlyApi *LifeYearlyApi) DeleteLifeYearlyByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := lifeYearlyService.DeleteLifeYearlyByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateLifeYearly 更新LifeYearly
// @Tags LifeYearly
// @Summary 更新LifeYearly
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body daily.LifeYearly true "更新LifeYearly"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lifeYearly/updateLifeYearly [put]
func (lifeYearlyApi *LifeYearlyApi) UpdateLifeYearly(c *gin.Context) {
	var lifeYearly daily.LifeYearly
	_ = c.ShouldBindJSON(&lifeYearly)
	if err := lifeYearlyService.UpdateLifeYearly(lifeYearly); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLifeYearly 用id查询LifeYearly
// @Tags LifeYearly
// @Summary 用id查询LifeYearly
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query daily.LifeYearly true "用id查询LifeYearly"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lifeYearly/findLifeYearly [get]
func (lifeYearlyApi *LifeYearlyApi) FindLifeYearly(c *gin.Context) {
	var lifeYearly daily.LifeYearly
	_ = c.ShouldBindQuery(&lifeYearly)
	if err, relifeYearly := lifeYearlyService.GetLifeYearly(lifeYearly.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relifeYearly": relifeYearly}, c)
	}
}

// GetLifeYearlyList 分页获取LifeYearly列表
// @Tags LifeYearly
// @Summary 分页获取LifeYearly列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dailyReq.LifeYearlySearch true "分页获取LifeYearly列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lifeYearly/getLifeYearlyList [get]
func (lifeYearlyApi *LifeYearlyApi) GetLifeYearlyList(c *gin.Context) {
	var pageInfo dailyReq.LifeYearlySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := lifeYearlyService.GetLifeYearlyInfoList(pageInfo); err != nil {
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

func (lifeYearlyApi *LifeYearlyApi) GetLifeYearlyOptions(c *gin.Context) {
	column := c.Query("column")
	err, res := lifeYearlyService.GetLifeYearlyColumn(column)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: res,
		}, "获取成功", c)
	}
}
