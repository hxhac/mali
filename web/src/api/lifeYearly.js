import service from '@/utils/request'

// @Tags LifeYearly
// @Summary 创建LifeYearly
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LifeYearly true "创建LifeYearly"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lifeYearly/createLifeYearly [post]
export const createLifeYearly = (data) => {
  return service({
    url: '/lifeYearly/createLifeYearly',
    method: 'post',
    data
  })
}

// @Tags LifeYearly
// @Summary 删除LifeYearly
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LifeYearly true "删除LifeYearly"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lifeYearly/deleteLifeYearly [delete]
export const deleteLifeYearly = (data) => {
  return service({
    url: '/lifeYearly/deleteLifeYearly',
    method: 'delete',
    data
  })
}

// @Tags LifeYearly
// @Summary 删除LifeYearly
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LifeYearly"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lifeYearly/deleteLifeYearly [delete]
export const deleteLifeYearlyByIds = (data) => {
  return service({
    url: '/lifeYearly/deleteLifeYearlyByIds',
    method: 'delete',
    data
  })
}

// @Tags LifeYearly
// @Summary 更新LifeYearly
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LifeYearly true "更新LifeYearly"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lifeYearly/updateLifeYearly [put]
export const updateLifeYearly = (data) => {
  return service({
    url: '/lifeYearly/updateLifeYearly',
    method: 'put',
    data
  })
}

// @Tags LifeYearly
// @Summary 用id查询LifeYearly
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.LifeYearly true "用id查询LifeYearly"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lifeYearly/findLifeYearly [get]
export const findLifeYearly = (params) => {
  return service({
    url: '/lifeYearly/findLifeYearly',
    method: 'get',
    params
  })
}

// @Tags LifeYearly
// @Summary 分页获取LifeYearly列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LifeYearly列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lifeYearly/getLifeYearlyList [get]
export const getLifeYearlyList = (params) => {
  return service({
    url: '/lifeYearly/getLifeYearlyList',
    method: 'get',
    params
  })
}
