import service from '@/utils/request'

// @Tags AppManage
// @Summary 创建AppManage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppManage true "创建AppManage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appManage/createAppManage [post]
export const createAppManage = (data) => {
  return service({
    url: '/appManage/createAppManage',
    method: 'post',
    data
  })
}

// @Tags AppManage
// @Summary 删除AppManage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppManage true "删除AppManage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appManage/deleteAppManage [delete]
export const deleteAppManage = (data) => {
  return service({
    url: '/appManage/deleteAppManage',
    method: 'delete',
    data
  })
}

// @Tags AppManage
// @Summary 删除AppManage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppManage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appManage/deleteAppManage [delete]
export const deleteAppManageByIds = (data) => {
  return service({
    url: '/appManage/deleteAppManageByIds',
    method: 'delete',
    data
  })
}

// @Tags AppManage
// @Summary 更新AppManage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppManage true "更新AppManage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appManage/updateAppManage [put]
export const updateAppManage = (data) => {
  return service({
    url: '/appManage/updateAppManage',
    method: 'put',
    data
  })
}

// @Tags AppManage
// @Summary 用id查询AppManage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppManage true "用id查询AppManage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appManage/findAppManage [get]
export const findAppManage = (params) => {
  return service({
    url: '/appManage/findAppManage',
    method: 'get',
    params
  })
}

// @Tags AppManage
// @Summary 分页获取AppManage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AppManage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appManage/getAppManageList [get]
export const getAppManageList = (params) => {
  return service({
    url: '/appManage/getAppManageList',
    method: 'get',
    params
  })
}
