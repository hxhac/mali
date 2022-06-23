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

type GoodsLabelApi struct{}

var goodsLabelService = service.ServiceGroupApp.GoodsServiceGroup.GoodsLabelService

// CreateGoodsLabel 创建GoodsLabel
// @Tags GoodsLabel
// @Summary 创建GoodsLabel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body goods.GoodsLabel true "创建GoodsLabel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodsLabel/createGoodsLabel [post]
func (goodsLabelApi *GoodsLabelApi) CreateGoodsLabel(c *gin.Context) {
	var goodsLabel goods.GoodsLabel
	_ = c.ShouldBindJSON(&goodsLabel)
	if err := goodsLabelService.CreateGoodsLabel(goodsLabel); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteGoodsLabel 删除GoodsLabel
// @Tags GoodsLabel
// @Summary 删除GoodsLabel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body goods.GoodsLabel true "删除GoodsLabel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goodsLabel/deleteGoodsLabel [delete]
func (goodsLabelApi *GoodsLabelApi) DeleteGoodsLabel(c *gin.Context) {
	var goodsLabel goods.GoodsLabel
	_ = c.ShouldBindJSON(&goodsLabel)
	if err := goodsLabelService.DeleteGoodsLabel(goodsLabel); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteGoodsLabelByIds 批量删除GoodsLabel
// @Tags GoodsLabel
// @Summary 批量删除GoodsLabel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除GoodsLabel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /goodsLabel/deleteGoodsLabelByIds [delete]
func (goodsLabelApi *GoodsLabelApi) DeleteGoodsLabelByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := goodsLabelService.DeleteGoodsLabelByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateGoodsLabel 更新GoodsLabel
// @Tags GoodsLabel
// @Summary 更新GoodsLabel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body goods.GoodsLabel true "更新GoodsLabel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /goodsLabel/updateGoodsLabel [put]
func (goodsLabelApi *GoodsLabelApi) UpdateGoodsLabel(c *gin.Context) {
	var goodsLabel goods.GoodsLabel
	_ = c.ShouldBindJSON(&goodsLabel)
	if err := goodsLabelService.UpdateGoodsLabel(goodsLabel); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindGoodsLabel 用id查询GoodsLabel
// @Tags GoodsLabel
// @Summary 用id查询GoodsLabel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query goods.GoodsLabel true "用id查询GoodsLabel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /goodsLabel/findGoodsLabel [get]
func (goodsLabelApi *GoodsLabelApi) FindGoodsLabel(c *gin.Context) {
	var goodsLabel goods.GoodsLabel
	_ = c.ShouldBindQuery(&goodsLabel)
	if err, regoodsLabel := goodsLabelService.GetGoodsLabel(goodsLabel.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"regoodsLabel": regoodsLabel}, c)
	}
}

// GetGoodsLabelList 分页获取GoodsLabel列表
// @Tags GoodsLabel
// @Summary 分页获取GoodsLabel列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query goodsReq.GoodsLabelSearch true "分页获取GoodsLabel列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodsLabel/getGoodsLabelList [get]
func (goodsLabelApi *GoodsLabelApi) GetGoodsLabelList(c *gin.Context) {
	var pageInfo goodsReq.GoodsLabelSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := goodsLabelService.GetGoodsLabelInfoList(pageInfo); err != nil {
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
