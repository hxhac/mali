import service from '@/utils/request'

// @Tags LifeEveryday
// @Summary 创建LifeEveryday
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LifeEveryday true "创建LifeEveryday"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lifeEveryday/createLifeEveryday [post]
export const createLifeEveryday = (data) => {
  return service({
    url: '/lifeEveryday/createLifeEveryday',
    method: 'post',
    data
  })
}

// @Tags LifeEveryday
// @Summary 删除LifeEveryday
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LifeEveryday true "删除LifeEveryday"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lifeEveryday/deleteLifeEveryday [delete]
export const deleteLifeEveryday = (data) => {
  return service({
    url: '/lifeEveryday/deleteLifeEveryday',
    method: 'delete',
    data
  })
}

// @Tags LifeEveryday
// @Summary 删除LifeEveryday
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LifeEveryday"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lifeEveryday/deleteLifeEveryday [delete]
export const deleteLifeEverydayByIds = (data) => {
  return service({
    url: '/lifeEveryday/deleteLifeEverydayByIds',
    method: 'delete',
    data
  })
}

// @Tags LifeEveryday
// @Summary 更新LifeEveryday
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LifeEveryday true "更新LifeEveryday"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lifeEveryday/updateLifeEveryday [put]
export const updateLifeEveryday = (data) => {
  return service({
    url: '/lifeEveryday/updateLifeEveryday',
    method: 'put',
    data
  })
}

// @Tags LifeEveryday
// @Summary 用id查询LifeEveryday
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.LifeEveryday true "用id查询LifeEveryday"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lifeEveryday/findLifeEveryday [get]
export const findLifeEveryday = (params) => {
  return service({
    url: '/lifeEveryday/findLifeEveryday',
    method: 'get',
    params
  })
}

// @Tags LifeEveryday
// @Summary 分页获取LifeEveryday列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LifeEveryday列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lifeEveryday/getLifeEverydayList [get]
export const getLifeEverydayList = (params) => {
  return service({
    url: '/lifeEveryday/getLifeEverydayList',
    method: 'get',
    params
  })
}
