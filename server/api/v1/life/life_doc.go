package life

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/life"
	lifeReq "github.com/flipped-aurora/gin-vue-admin/server/model/life/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LifeDocApi struct {
}

var lifeDocService = service.ServiceGroupApp.LifeServiceGroup.LifeDocService

// CreateLifeDoc 创建LifeDoc
// @Tags LifeDoc
// @Summary 创建LifeDoc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body life.LifeDoc true "创建LifeDoc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lifeDoc/createLifeDoc [post]
func (lifeDocApi *LifeDocApi) CreateLifeDoc(c *gin.Context) {
	var lifeDoc life.LifeDoc
	_ = c.ShouldBindJSON(&lifeDoc)
	if err := lifeDocService.CreateLifeDoc(lifeDoc); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLifeDoc 删除LifeDoc
// @Tags LifeDoc
// @Summary 删除LifeDoc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body life.LifeDoc true "删除LifeDoc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lifeDoc/deleteLifeDoc [delete]
func (lifeDocApi *LifeDocApi) DeleteLifeDoc(c *gin.Context) {
	var lifeDoc life.LifeDoc
	_ = c.ShouldBindJSON(&lifeDoc)
	if err := lifeDocService.DeleteLifeDoc(lifeDoc); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteLifeDocByIds 批量删除LifeDoc
// @Tags LifeDoc
// @Summary 批量删除LifeDoc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LifeDoc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /lifeDoc/deleteLifeDocByIds [delete]
func (lifeDocApi *LifeDocApi) DeleteLifeDocByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := lifeDocService.DeleteLifeDocByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateLifeDoc 更新LifeDoc
// @Tags LifeDoc
// @Summary 更新LifeDoc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body life.LifeDoc true "更新LifeDoc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lifeDoc/updateLifeDoc [put]
func (lifeDocApi *LifeDocApi) UpdateLifeDoc(c *gin.Context) {
	var lifeDoc life.LifeDoc
	_ = c.ShouldBindJSON(&lifeDoc)
	if err := lifeDocService.UpdateLifeDoc(lifeDoc); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLifeDoc 用id查询LifeDoc
// @Tags LifeDoc
// @Summary 用id查询LifeDoc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query life.LifeDoc true "用id查询LifeDoc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lifeDoc/findLifeDoc [get]
func (lifeDocApi *LifeDocApi) FindLifeDoc(c *gin.Context) {
	var lifeDoc life.LifeDoc
	_ = c.ShouldBindQuery(&lifeDoc)
	if err, relifeDoc := lifeDocService.GetLifeDoc(lifeDoc.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relifeDoc": relifeDoc}, c)
	}
}

// GetLifeDocList 分页获取LifeDoc列表
// @Tags LifeDoc
// @Summary 分页获取LifeDoc列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lifeReq.LifeDocSearch true "分页获取LifeDoc列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lifeDoc/getLifeDocList [get]
func (lifeDocApi *LifeDocApi) GetLifeDocList(c *gin.Context) {
	var pageInfo lifeReq.LifeDocSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := lifeDocService.GetLifeDocInfoList(pageInfo); err != nil {
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
