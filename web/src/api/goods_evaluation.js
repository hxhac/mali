import service from '@/utils/request'

// @Tags GoodsEvaluation
// @Summary 创建GoodsEvaluation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoodsEvaluation true "创建GoodsEvaluation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ge/createGoodsEvaluation [post]
export const createGoodsEvaluation = (data) => {
  return service({
    url: '/ge/createGoodsEvaluation',
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
// @Router /ge/deleteGoodsEvaluation [delete]
export const deleteGoodsEvaluation = (data) => {
  return service({
    url: '/ge/deleteGoodsEvaluation',
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
// @Router /ge/deleteGoodsEvaluation [delete]
export const deleteGoodsEvaluationByIds = (data) => {
  return service({
    url: '/ge/deleteGoodsEvaluationByIds',
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
// @Router /ge/updateGoodsEvaluation [put]
export const updateGoodsEvaluation = (data) => {
  return service({
    url: '/ge/updateGoodsEvaluation',
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
// @Router /ge/findGoodsEvaluation [get]
export const findGoodsEvaluation = (params) => {
  return service({
    url: '/ge/findGoodsEvaluation',
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
// @Router /ge/getGoodsEvaluationList [get]
export const getGoodsEvaluationList = (params) => {
  return service({
    url: '/ge/getGoodsEvaluationList',
    method: 'get',
    params
  })
}
