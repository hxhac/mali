import service from '@/utils/request'

// @Tags RssFeed
// @Summary 创建RssFeed
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RssFeed true "创建RssFeed"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rssFeed/createRssFeed [post]
export const createRssFeed = (data) => {
  return service({
    url: '/rssFeed/createRssFeed',
    method: 'post',
    data
  })
}

// @Tags RssFeed
// @Summary 删除RssFeed
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RssFeed true "删除RssFeed"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rssFeed/deleteRssFeed [delete]
export const deleteRssFeed = (data) => {
  return service({
    url: '/rssFeed/deleteRssFeed',
    method: 'delete',
    data
  })
}

// @Tags RssFeed
// @Summary 删除RssFeed
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RssFeed"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rssFeed/deleteRssFeed [delete]
export const deleteRssFeedByIds = (data) => {
  return service({
    url: '/rssFeed/deleteRssFeedByIds',
    method: 'delete',
    data
  })
}

// @Tags RssFeed
// @Summary 更新RssFeed
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RssFeed true "更新RssFeed"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rssFeed/updateRssFeed [put]
export const updateRssFeed = (data) => {
  return service({
    url: '/rssFeed/updateRssFeed',
    method: 'put',
    data
  })
}

// @Tags RssFeed
// @Summary 用id查询RssFeed
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.RssFeed true "用id查询RssFeed"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /rssFeed/findRssFeed [get]
export const findRssFeed = (params) => {
  return service({
    url: '/rssFeed/findRssFeed',
    method: 'get',
    params
  })
}

// @Tags RssFeed
// @Summary 分页获取RssFeed列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取RssFeed列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rssFeed/getRssFeedList [get]
export const getRssFeedList = (params) => {
  return service({
    url: '/rssFeed/getRssFeedList',
    method: 'get',
    params
  })
}
