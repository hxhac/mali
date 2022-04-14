import service from '@/utils/request'

// @Tags GoodsEvaluation
// @Summary 创建GoodsEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoodsEvaluation true "创建GoodsEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodsEvaluation/createGoodsEvaluation [post]
export const createGoodsEvaluation = (data) => {
  return service({
    url: '/goodsEvaluation/createGoodsEvaluation',
    method: 'post',
    data
  })
}

// @Tags GoodsEvaluation
// @Summary 删除GoodsEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoodsEvaluation true "删除GoodsEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goodsEvaluation/deleteGoodsEvaluation [delete]
export const deleteGoodsEvaluation = (data) => {
  return service({
    url: '/goodsEvaluation/deleteGoodsEvaluation',
    method: 'delete',
    data
  })
}

// @Tags GoodsEvaluation
// @Summary 删除GoodsEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除GoodsEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goodsEvaluation/deleteGoodsEvaluation [delete]
export const deleteGoodsEvaluationByIds = (data) => {
  return service({
    url: '/goodsEvaluation/deleteGoodsEvaluationByIds',
    method: 'delete',
    data
  })
}

// @Tags GoodsEvaluation
// @Summary 更新GoodsEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoodsEvaluation true "更新GoodsEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /goodsEvaluation/updateGoodsEvaluation [put]
export const updateGoodsEvaluation = (data) => {
  return service({
    url: '/goodsEvaluation/updateGoodsEvaluation',
    method: 'put',
    data
  })
}

// @Tags GoodsEvaluation
// @Summary 用id查询GoodsEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.GoodsEvaluation true "用id查询GoodsEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /goodsEvaluation/findGoodsEvaluation [get]
export const findGoodsEvaluation = (params) => {
  return service({
    url: '/goodsEvaluation/findGoodsEvaluation',
    method: 'get',
    params
  })
}

// @Tags GoodsEvaluation
// @Summary 分页获取GoodsEvaluation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取GoodsEvaluation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodsEvaluation/getGoodsEvaluationList [get]
export const getGoodsEvaluationList = (params) => {
  return service({
    url: '/goodsEvaluation/getGoodsEvaluationList',
    method: 'get',
    params
  })
}
