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

type GoodsCategoryApi struct{}

var goodsCategoryService = service.ServiceGroupApp.GoodsServiceGroup.GoodsCategoryService

// CreateGoodsCategory 创建GoodsCategory
// @Tags GoodsCategory
// @Summary 创建GoodsCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body goods.GoodsCategory true "创建GoodsCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodsCategory/createGoodsCategory [post]
func (goodsCategoryApi *GoodsCategoryApi) CreateGoodsCategory(c *gin.Context) {
	var goodsCategory goods.GoodsCategory
	_ = c.ShouldBindJSON(&goodsCategory)
	if err := goodsCategoryService.CreateGoodsCategory(goodsCategory); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteGoodsCategory 删除GoodsCategory
// @Tags GoodsCategory
// @Summary 删除GoodsCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body goods.GoodsCategory true "删除GoodsCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goodsCategory/deleteGoodsCategory [delete]
func (goodsCategoryApi *GoodsCategoryApi) DeleteGoodsCategory(c *gin.Context) {
	var goodsCategory goods.GoodsCategory
	_ = c.ShouldBindJSON(&goodsCategory)
	if err := goodsCategoryService.DeleteGoodsCategory(goodsCategory); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteGoodsCategoryByIds 批量删除GoodsCategory
// @Tags GoodsCategory
// @Summary 批量删除GoodsCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除GoodsCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /goodsCategory/deleteGoodsCategoryByIds [delete]
func (goodsCategoryApi *GoodsCategoryApi) DeleteGoodsCategoryByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := goodsCategoryService.DeleteGoodsCategoryByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateGoodsCategory 更新GoodsCategory
// @Tags GoodsCategory
// @Summary 更新GoodsCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body goods.GoodsCategory true "更新GoodsCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /goodsCategory/updateGoodsCategory [put]
func (goodsCategoryApi *GoodsCategoryApi) UpdateGoodsCategory(c *gin.Context) {
	var goodsCategory goods.GoodsCategory
	_ = c.ShouldBindJSON(&goodsCategory)
	if err := goodsCategoryService.UpdateGoodsCategory(goodsCategory); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindGoodsCategory 用id查询GoodsCategory
// @Tags GoodsCategory
// @Summary 用id查询GoodsCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query goods.GoodsCategory true "用id查询GoodsCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /goodsCategory/findGoodsCategory [get]
func (goodsCategoryApi *GoodsCategoryApi) FindGoodsCategory(c *gin.Context) {
	var goodsCategory goods.GoodsCategory
	_ = c.ShouldBindQuery(&goodsCategory)
	if err, regoodsCategory := goodsCategoryService.GetGoodsCategory(goodsCategory.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"regoodsCategory": regoodsCategory}, c)
	}
}

// GetGoodsCategoryList 分页获取GoodsCategory列表
// @Tags GoodsCategory
// @Summary 分页获取GoodsCategory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query goodsReq.GoodsCategorySearch true "分页获取GoodsCategory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodsCategory/getGoodsCategoryList [get]
func (goodsCategoryApi *GoodsCategoryApi) GetGoodsCategoryList(c *gin.Context) {
	var pageInfo goodsReq.GoodsCategorySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := goodsCategoryService.GetGoodsCategoryInfoList(pageInfo); err != nil {
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
