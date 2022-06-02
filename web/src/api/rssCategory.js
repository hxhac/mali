import service from '@/utils/request'

// @Tags RssCategory
// @Summary 创建RssCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RssCategory true "创建RssCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rssCategory/createRssCategory [post]
export const createRssCategory = (data) => {
  return service({
    url: '/rssCategory/createRssCategory',
    method: 'post',
    data
  })
}

// @Tags RssCategory
// @Summary 删除RssCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RssCategory true "删除RssCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rssCategory/deleteRssCategory [delete]
export const deleteRssCategory = (data) => {
  return service({
    url: '/rssCategory/deleteRssCategory',
    method: 'delete',
    data
  })
}

// @Tags RssCategory
// @Summary 删除RssCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RssCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rssCategory/deleteRssCategory [delete]
export const deleteRssCategoryByIds = (data) => {
  return service({
    url: '/rssCategory/deleteRssCategoryByIds',
    method: 'delete',
    data
  })
}

// @Tags RssCategory
// @Summary 更新RssCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RssCategory true "更新RssCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rssCategory/updateRssCategory [put]
export const updateRssCategory = (data) => {
  return service({
    url: '/rssCategory/updateRssCategory',
    method: 'put',
    data
  })
}

// @Tags RssCategory
// @Summary 用id查询RssCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.RssCategory true "用id查询RssCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /rssCategory/findRssCategory [get]
export const findRssCategory = (params) => {
  return service({
    url: '/rssCategory/findRssCategory',
    method: 'get',
    params
  })
}

// @Tags RssCategory
// @Summary 分页获取RssCategory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取RssCategory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rssCategory/getRssCategoryList [get]
export const getRssCategoryList = (params) => {
  return service({
    url: '/rssCategory/getRssCategoryList',
    method: 'get',
    params
  })
}
