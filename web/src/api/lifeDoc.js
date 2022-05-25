import service from '@/utils/request'

// @Tags LifeDoc
// @Summary 创建LifeDoc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LifeDoc true "创建LifeDoc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lifeDoc/createLifeDoc [post]
export const createLifeDoc = (data) => {
  return service({
    url: '/lifeDoc/createLifeDoc',
    method: 'post',
    data
  })
}

// @Tags LifeDoc
// @Summary 删除LifeDoc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LifeDoc true "删除LifeDoc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lifeDoc/deleteLifeDoc [delete]
export const deleteLifeDoc = (data) => {
  return service({
    url: '/lifeDoc/deleteLifeDoc',
    method: 'delete',
    data
  })
}

// @Tags LifeDoc
// @Summary 删除LifeDoc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LifeDoc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lifeDoc/deleteLifeDoc [delete]
export const deleteLifeDocByIds = (data) => {
  return service({
    url: '/lifeDoc/deleteLifeDocByIds',
    method: 'delete',
    data
  })
}

// @Tags LifeDoc
// @Summary 更新LifeDoc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LifeDoc true "更新LifeDoc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lifeDoc/updateLifeDoc [put]
export const updateLifeDoc = (data) => {
  return service({
    url: '/lifeDoc/updateLifeDoc',
    method: 'put',
    data
  })
}

// @Tags LifeDoc
// @Summary 用id查询LifeDoc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.LifeDoc true "用id查询LifeDoc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lifeDoc/findLifeDoc [get]
export const findLifeDoc = (params) => {
  return service({
    url: '/lifeDoc/findLifeDoc',
    method: 'get',
    params
  })
}

// @Tags LifeDoc
// @Summary 分页获取LifeDoc列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LifeDoc列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lifeDoc/getLifeDocList [get]
export const getLifeDocList = (params) => {
  return service({
    url: '/lifeDoc/getLifeDocList',
    method: 'get',
    params
  })
}
