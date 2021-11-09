import service from '@/utils/request'

// @Tags CmsAdSeat
// @Summary 创建CmsAdSeat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsAdSeat true "创建CmsAdSeat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsAdSeat/createCmsAdSeat [post]
export const createCmsAdSeat = (data) => {
  return service({
    url: '/cmsAdSeat/createCmsAdSeat',
    method: 'post',
    data
  })
}

// @Tags CmsAdSeat
// @Summary 删除CmsAdSeat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsAdSeat true "删除CmsAdSeat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsAdSeat/deleteCmsAdSeat [delete]
export const deleteCmsAdSeat = (data) => {
  return service({
    url: '/cmsAdSeat/deleteCmsAdSeat',
    method: 'delete',
    data
  })
}

// @Tags CmsAdSeat
// @Summary 删除CmsAdSeat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CmsAdSeat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsAdSeat/deleteCmsAdSeat [delete]
export const deleteCmsAdSeatByIds = (data) => {
  return service({
    url: '/cmsAdSeat/deleteCmsAdSeatByIds',
    method: 'delete',
    data
  })
}

// @Tags CmsAdSeat
// @Summary 更新CmsAdSeat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsAdSeat true "更新CmsAdSeat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsAdSeat/updateCmsAdSeat [put]
export const updateCmsAdSeat = (data) => {
  return service({
    url: '/cmsAdSeat/updateCmsAdSeat',
    method: 'put',
    data
  })
}

// @Tags CmsAdSeat
// @Summary 用id查询CmsAdSeat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CmsAdSeat true "用id查询CmsAdSeat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsAdSeat/findCmsAdSeat [get]
export const findCmsAdSeat = (params) => {
  return service({
    url: '/cmsAdSeat/findCmsAdSeat',
    method: 'get',
    params
  })
}

// @Tags CmsAdSeat
// @Summary 分页获取CmsAdSeat列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CmsAdSeat列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsAdSeat/getCmsAdSeatList [get]
export const getCmsAdSeatList = (params) => {
  return service({
    url: '/cmsAdSeat/getCmsAdSeatList',
    method: 'get',
    params
  })
}


// @Tags CmsAdSeat
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsAdSeat true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsAdSeat/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/cmsAdSeat/quickEdit',
    method: 'post',
    data
  })
}