import service from '@/utils/request'

// @Tags GoodsBrand
// @Summary 创建GoodsBrand
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoodsBrand true "创建GoodsBrand"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodsBrand/createGoodsBrand [post]
export const createGoodsBrand = (data) => {
  return service({
    url: '/goodsBrand/createGoodsBrand',
    method: 'post',
    data
  })
}

// @Tags GoodsBrand
// @Summary 删除GoodsBrand
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoodsBrand true "删除GoodsBrand"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goodsBrand/deleteGoodsBrand [delete]
export const deleteGoodsBrand = (data) => {
  return service({
    url: '/goodsBrand/deleteGoodsBrand',
    method: 'delete',
    data
  })
}

// @Tags GoodsBrand
// @Summary 删除GoodsBrand
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除GoodsBrand"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goodsBrand/deleteGoodsBrand [delete]
export const deleteGoodsBrandByIds = (data) => {
  return service({
    url: '/goodsBrand/deleteGoodsBrandByIds',
    method: 'delete',
    data
  })
}

// @Tags GoodsBrand
// @Summary 更新GoodsBrand
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoodsBrand true "更新GoodsBrand"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /goodsBrand/updateGoodsBrand [put]
export const updateGoodsBrand = (data) => {
  return service({
    url: '/goodsBrand/updateGoodsBrand',
    method: 'put',
    data
  })
}

// @Tags GoodsBrand
// @Summary 用id查询GoodsBrand
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.GoodsBrand true "用id查询GoodsBrand"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /goodsBrand/findGoodsBrand [get]
export const findGoodsBrand = (params) => {
  return service({
    url: '/goodsBrand/findGoodsBrand',
    method: 'get',
    params
  })
}

// @Tags GoodsBrand
// @Summary 分页获取GoodsBrand列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取GoodsBrand列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodsBrand/getGoodsBrandList [get]
export const getGoodsBrandList = (params) => {
  return service({
    url: '/goodsBrand/getGoodsBrandList',
    method: 'get',
    params
  })
}
