import service from '@/utils/request'

// @Tags CmsAd
// @Summary 创建CmsAd
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsAd true "创建CmsAd"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsAd/createCmsAd [post]
export const createCmsAd = (data) => {
  return service({
    url: '/cmsAd/createCmsAd',
    method: 'post',
    data
  })
}

// @Tags CmsAd
// @Summary 删除CmsAd
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsAd true "删除CmsAd"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsAd/deleteCmsAd [delete]
export const deleteCmsAd = (data) => {
  return service({
    url: '/cmsAd/deleteCmsAd',
    method: 'delete',
    data
  })
}

// @Tags CmsAd
// @Summary 删除CmsAd
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CmsAd"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsAd/deleteCmsAd [delete]
export const deleteCmsAdByIds = (data) => {
  return service({
    url: '/cmsAd/deleteCmsAdByIds',
    method: 'delete',
    data
  })
}

// @Tags CmsAd
// @Summary 更新CmsAd
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsAd true "更新CmsAd"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsAd/updateCmsAd [put]
export const updateCmsAd = (data) => {
  return service({
    url: '/cmsAd/updateCmsAd',
    method: 'put',
    data
  })
}

// @Tags CmsAd
// @Summary 用id查询CmsAd
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CmsAd true "用id查询CmsAd"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsAd/findCmsAd [get]
export const findCmsAd = (params) => {
  return service({
    url: '/cmsAd/findCmsAd',
    method: 'get',
    params
  })
}

// @Tags CmsAd
// @Summary 分页获取CmsAd列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CmsAd列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsAd/getCmsAdList [get]
export const getCmsAdList = (params) => {
  return service({
    url: '/cmsAd/getCmsAdList',
    method: 'get',
    params
  })
}


// @Tags CmsAd
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsAd true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsAd/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/cmsAd/quickEdit',
    method: 'post',
    data
  })
}