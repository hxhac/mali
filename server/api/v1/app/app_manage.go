package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	xxxReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AppManageApi struct {
}

var appManageService = service.ServiceGroupApp.AppServiceGroup.AppManageService

// CreateAppManage 创建AppManage
// @Tags AppManage
// @Summary 创建AppManage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xxx.AppManage true "创建AppManage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appManage/createAppManage [post]
func (appManageApi *AppManageApi) CreateAppManage(c *gin.Context) {
	var appManage app.AppManage
	_ = c.ShouldBindJSON(&appManage)
	if err := appManageService.CreateAppManage(appManage); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppManage 删除AppManage
// @Tags AppManage
// @Summary 删除AppManage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xxx.AppManage true "删除AppManage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appManage/deleteAppManage [delete]
func (appManageApi *AppManageApi) DeleteAppManage(c *gin.Context) {
	var appManage app.AppManage
	_ = c.ShouldBindJSON(&appManage)
	if err := appManageService.DeleteAppManage(appManage); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppManageByIds 批量删除AppManage
// @Tags AppManage
// @Summary 批量删除AppManage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppManage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appManage/deleteAppManageByIds [delete]
func (appManageApi *AppManageApi) DeleteAppManageByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := appManageService.DeleteAppManageByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppManage 更新AppManage
// @Tags AppManage
// @Summary 更新AppManage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xxx.AppManage true "更新AppManage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appManage/updateAppManage [put]
func (appManageApi *AppManageApi) UpdateAppManage(c *gin.Context) {
	var appManage app.AppManage
	_ = c.ShouldBindJSON(&appManage)
	if err := appManageService.UpdateAppManage(appManage); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppManage 用id查询AppManage
// @Tags AppManage
// @Summary 用id查询AppManage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query xxx.AppManage true "用id查询AppManage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appManage/findAppManage [get]
func (appManageApi *AppManageApi) FindAppManage(c *gin.Context) {
	var appManage app.AppManage
	_ = c.ShouldBindQuery(&appManage)
	if err, reappManage := appManageService.GetAppManage(appManage.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappManage": reappManage}, c)
	}
}

// GetAppManageList 分页获取AppManage列表
// @Tags AppManage
// @Summary 分页获取AppManage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query xxxReq.AppManageSearch true "分页获取AppManage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appManage/getAppManageList [get]
func (appManageApi *AppManageApi) GetAppManageList(c *gin.Context) {
	var pageInfo xxxReq.AppManageSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := appManageService.GetAppManageInfoList(pageInfo); err != nil {
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
