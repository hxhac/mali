package goods

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/goods"
	goodsReq "github.com/flipped-aurora/gin-vue-admin/server/model/goods/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GoodsEvaluationApi struct{}

var goodsEvaluationService = service.ServiceGroupApp.GoodsServiceGroup.GoodsEvaluationService

// CreateGoodsEvaluation 创建GoodsEvaluation
// @Tags GoodsEvaluation
// @Summary 创建GoodsEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body goods.GoodsEvaluation true "创建GoodsEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodsEvaluation/createGoodsEvaluation [post]
func (goodsEvaluationApi *GoodsEvaluationApi) CreateGoodsEvaluation(c *gin.Context) {
	var goodsEvaluation goods.GoodsEvaluation
	_ = c.ShouldBindJSON(&goodsEvaluation)
	if err := goodsEvaluationService.CreateGoodsEvaluation(goodsEvaluation); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteGoodsEvaluation 删除GoodsEvaluation
// @Tags GoodsEvaluation
// @Summary 删除GoodsEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body goods.GoodsEvaluation true "删除GoodsEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goodsEvaluation/deleteGoodsEvaluation [delete]
func (goodsEvaluationApi *GoodsEvaluationApi) DeleteGoodsEvaluation(c *gin.Context) {
	var goodsEvaluation goods.GoodsEvaluation
	_ = c.ShouldBindJSON(&goodsEvaluation)
	if err := goodsEvaluationService.DeleteGoodsEvaluation(goodsEvaluation); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteGoodsEvaluationByIds 批量删除GoodsEvaluation
// @Tags GoodsEvaluation
// @Summary 批量删除GoodsEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除GoodsEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /goodsEvaluation/deleteGoodsEvaluationByIds [delete]
func (goodsEvaluationApi *GoodsEvaluationApi) DeleteGoodsEvaluationByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := goodsEvaluationService.DeleteGoodsEvaluationByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateGoodsEvaluation 更新GoodsEvaluation
// @Tags GoodsEvaluation
// @Summary 更新GoodsEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body goods.GoodsEvaluation true "更新GoodsEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /goodsEvaluation/updateGoodsEvaluation [put]
func (goodsEvaluationApi *GoodsEvaluationApi) UpdateGoodsEvaluation(c *gin.Context) {
	var goodsEvaluation goods.GoodsEvaluation
	_ = c.ShouldBindJSON(&goodsEvaluation)
	if err := goodsEvaluationService.UpdateGoodsEvaluation(goodsEvaluation); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindGoodsEvaluation 用id查询GoodsEvaluation
// @Tags GoodsEvaluation
// @Summary 用id查询GoodsEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query goods.GoodsEvaluation true "用id查询GoodsEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /goodsEvaluation/findGoodsEvaluation [get]
func (goodsEvaluationApi *GoodsEvaluationApi) FindGoodsEvaluation(c *gin.Context) {
	var goodsEvaluation goods.GoodsEvaluation
	_ = c.ShouldBindQuery(&goodsEvaluation)
	if err, regoodsEvaluation := goodsEvaluationService.GetGoodsEvaluation(goodsEvaluation.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"regoodsEvaluation": regoodsEvaluation}, c)
	}
}

// GetGoodsEvaluationList 分页获取GoodsEvaluation列表
// @Tags GoodsEvaluation
// @Summary 分页获取GoodsEvaluation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query goodsReq.GoodsEvaluationSearch true "分页获取GoodsEvaluation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodsEvaluation/getGoodsEvaluationList [get]
func (goodsEvaluationApi *GoodsEvaluationApi) GetGoodsEvaluationList(c *gin.Context) {
	var pageInfo goodsReq.GoodsEvaluationSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := goodsEvaluationService.GetGoodsEvaluationInfoList(pageInfo); err != nil {
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

// 获取options数据
func (GoodsEvaluationApi) GetGoodsEvaluationOptions(c *gin.Context) {
	column := c.Query("column")
	err, res := goodsEvaluationService.GetGoodsEvaluationColumn(column)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: res,
		}, "获取成功", c)
	}
}
