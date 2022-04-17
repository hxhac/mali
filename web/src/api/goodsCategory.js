import service from '@/utils/request'

// @Tags GoodsCategory
// @Summary 创建GoodsCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoodsCategory true "创建GoodsCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodsCategory/createGoodsCategory [post]
export const createGoodsCategory = (data) => {
	return service({
		url: '/goodsCategory/createGoodsCategory',
		method: 'post',
		data
	})
}

// @Tags GoodsCategory
// @Summary 删除GoodsCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoodsCategory true "删除GoodsCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goodsCategory/deleteGoodsCategory [delete]
export const deleteGoodsCategory = (data) => {
	return service({
		url: '/goodsCategory/deleteGoodsCategory',
		method: 'delete',
		data
	})
}

// @Tags GoodsCategory
// @Summary 删除GoodsCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除GoodsCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goodsCategory/deleteGoodsCategory [delete]
export const deleteGoodsCategoryByIds = (data) => {
	return service({
		url: '/goodsCategory/deleteGoodsCategoryByIds',
		method: 'delete',
		data
	})
}

// @Tags GoodsCategory
// @Summary 更新GoodsCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoodsCategory true "更新GoodsCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /goodsCategory/updateGoodsCategory [put]
export const updateGoodsCategory = (data) => {
	return service({
		url: '/goodsCategory/updateGoodsCategory',
		method: 'put',
		data
	})
}

// @Tags GoodsCategory
// @Summary 用id查询GoodsCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.GoodsCategory true "用id查询GoodsCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /goodsCategory/findGoodsCategory [get]
export const findGoodsCategory = (params) => {
	return service({
		url: '/goodsCategory/findGoodsCategory',
		method: 'get',
		params
	})
}

// @Tags GoodsCategory
// @Summary 分页获取GoodsCategory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取GoodsCategory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodsCategory/getGoodsCategoryList [get]
export const getGoodsCategoryList = (params) => {
	return service({
		url: '/goodsCategory/getGoodsCategoryList',
		method: 'get',
		params
	})
}
