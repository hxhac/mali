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

type GoodsBrandApi struct{}

var goodsBrandService = service.ServiceGroupApp.GoodsServiceGroup.GoodsBrandService

// CreateGoodsBrand 创建GoodsBrand
// @Tags GoodsBrand
// @Summary 创建GoodsBrand
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body goods.GoodsBrand true "创建GoodsBrand"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodsBrand/createGoodsBrand [post]
func (goodsBrandApi *GoodsBrandApi) CreateGoodsBrand(c *gin.Context) {
	var goodsBrand goods.GoodsBrand
	_ = c.ShouldBindJSON(&goodsBrand)
	if err := goodsBrandService.CreateGoodsBrand(goodsBrand); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteGoodsBrand 删除GoodsBrand
// @Tags GoodsBrand
// @Summary 删除GoodsBrand
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body goods.GoodsBrand true "删除GoodsBrand"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goodsBrand/deleteGoodsBrand [delete]
func (goodsBrandApi *GoodsBrandApi) DeleteGoodsBrand(c *gin.Context) {
	var goodsBrand goods.GoodsBrand
	_ = c.ShouldBindJSON(&goodsBrand)
	if err := goodsBrandService.DeleteGoodsBrand(goodsBrand); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteGoodsBrandByIds 批量删除GoodsBrand
// @Tags GoodsBrand
// @Summary 批量删除GoodsBrand
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除GoodsBrand"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /goodsBrand/deleteGoodsBrandByIds [delete]
func (goodsBrandApi *GoodsBrandApi) DeleteGoodsBrandByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := goodsBrandService.DeleteGoodsBrandByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateGoodsBrand 更新GoodsBrand
// @Tags GoodsBrand
// @Summary 更新GoodsBrand
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body goods.GoodsBrand true "更新GoodsBrand"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /goodsBrand/updateGoodsBrand [put]
func (goodsBrandApi *GoodsBrandApi) UpdateGoodsBrand(c *gin.Context) {
	var goodsBrand goods.GoodsBrand
	_ = c.ShouldBindJSON(&goodsBrand)
	if err := goodsBrandService.UpdateGoodsBrand(goodsBrand); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindGoodsBrand 用id查询GoodsBrand
// @Tags GoodsBrand
// @Summary 用id查询GoodsBrand
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query goods.GoodsBrand true "用id查询GoodsBrand"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /goodsBrand/findGoodsBrand [get]
func (goodsBrandApi *GoodsBrandApi) FindGoodsBrand(c *gin.Context) {
	var goodsBrand goods.GoodsBrand
	_ = c.ShouldBindQuery(&goodsBrand)
	if err, regoodsBrand := goodsBrandService.GetGoodsBrand(goodsBrand.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"regoodsBrand": regoodsBrand}, c)
	}
}

// GetGoodsBrandList 分页获取GoodsBrand列表
// @Tags GoodsBrand
// @Summary 分页获取GoodsBrand列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query goodsReq.GoodsBrandSearch true "分页获取GoodsBrand列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodsBrand/getGoodsBrandList [get]
func (goodsBrandApi *GoodsBrandApi) GetGoodsBrandList(c *gin.Context) {
	var pageInfo goodsReq.GoodsBrandSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := goodsBrandService.GetGoodsBrandInfoList(pageInfo); err != nil {
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
