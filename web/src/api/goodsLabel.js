import service from '@/utils/request'

// @Tags GoodsLabel
// @Summary 创建GoodsLabel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoodsLabel true "创建GoodsLabel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodsLabel/createGoodsLabel [post]
export const createGoodsLabel = (data) => {
  return service({
    url: '/goodsLabel/createGoodsLabel',
    method: 'post',
    data
  })
}

// @Tags GoodsLabel
// @Summary 删除GoodsLabel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoodsLabel true "删除GoodsLabel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goodsLabel/deleteGoodsLabel [delete]
export const deleteGoodsLabel = (data) => {
  return service({
    url: '/goodsLabel/deleteGoodsLabel',
    method: 'delete',
    data
  })
}

// @Tags GoodsLabel
// @Summary 删除GoodsLabel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除GoodsLabel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goodsLabel/deleteGoodsLabel [delete]
export const deleteGoodsLabelByIds = (data) => {
  return service({
    url: '/goodsLabel/deleteGoodsLabelByIds',
    method: 'delete',
    data
  })
}

// @Tags GoodsLabel
// @Summary 更新GoodsLabel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoodsLabel true "更新GoodsLabel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /goodsLabel/updateGoodsLabel [put]
export const updateGoodsLabel = (data) => {
  return service({
    url: '/goodsLabel/updateGoodsLabel',
    method: 'put',
    data
  })
}

// @Tags GoodsLabel
// @Summary 用id查询GoodsLabel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.GoodsLabel true "用id查询GoodsLabel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /goodsLabel/findGoodsLabel [get]
export const findGoodsLabel = (params) => {
  return service({
    url: '/goodsLabel/findGoodsLabel',
    method: 'get',
    params
  })
}

// @Tags GoodsLabel
// @Summary 分页获取GoodsLabel列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取GoodsLabel列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodsLabel/getGoodsLabelList [get]
export const getGoodsLabelList = (params) => {
  return service({
    url: '/goodsLabel/getGoodsLabelList',
    method: 'get',
    params
  })
}
