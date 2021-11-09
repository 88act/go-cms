import service from '@/utils/request'

// @Tags CmsArticle
// @Summary 创建CmsArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsArticle true "创建CmsArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsArticle/createCmsArticle [post]
export const createCmsArticle = (data) => {
  return service({
    url: '/cmsArticle/createCmsArticle',
    method: 'post',
    data
  })
}

// @Tags CmsArticle
// @Summary 删除CmsArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsArticle true "删除CmsArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsArticle/deleteCmsArticle [delete]
export const deleteCmsArticle = (data) => {
  return service({
    url: '/cmsArticle/deleteCmsArticle',
    method: 'delete',
    data
  })
}

// @Tags CmsArticle
// @Summary 删除CmsArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CmsArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsArticle/deleteCmsArticle [delete]
export const deleteCmsArticleByIds = (data) => {
  return service({
    url: '/cmsArticle/deleteCmsArticleByIds',
    method: 'delete',
    data
  })
}

// @Tags CmsArticle
// @Summary 更新CmsArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsArticle true "更新CmsArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsArticle/updateCmsArticle [put]
export const updateCmsArticle = (data) => {
  return service({
    url: '/cmsArticle/updateCmsArticle',
    method: 'put',
    data
  })
}

// @Tags CmsArticle
// @Summary 用id查询CmsArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CmsArticle true "用id查询CmsArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsArticle/findCmsArticle [get]
export const findCmsArticle = (params) => {
  return service({
    url: '/cmsArticle/findCmsArticle',
    method: 'get',
    params
  })
}

// @Tags CmsArticle
// @Summary 分页获取CmsArticle列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CmsArticle列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsArticle/getCmsArticleList [get]
export const getCmsArticleList = (params) => {
  return service({
    url: '/cmsArticle/getCmsArticleList',
    method: 'get',
    params
  })
}


// @Tags CmsArticle
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsArticle true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsArticle/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/cmsArticle/quickEdit',
    method: 'post',
    data
  })
}