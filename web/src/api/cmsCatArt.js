import service from '@/utils/request'

// @Tags CmsCatArt
// @Summary 创建CmsCatArt
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsCatArt true "创建CmsCatArt"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsCatArt/createCmsCatArt [post]
export const createCmsCatArt = (data) => {
  return service({
    url: '/cmsCatArt/createCmsCatArt',
    method: 'post',
    data
  })
}

// @Tags CmsCatArt
// @Summary 删除CmsCatArt
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsCatArt true "删除CmsCatArt"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsCatArt/deleteCmsCatArt [delete]
export const deleteCmsCatArt = (data) => {
  return service({
    url: '/cmsCatArt/deleteCmsCatArt',
    method: 'delete',
    data
  })
}

// @Tags CmsCatArt
// @Summary 删除CmsCatArt
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CmsCatArt"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsCatArt/deleteCmsCatArt [delete]
export const deleteCmsCatArtByIds = (data) => {
  return service({
    url: '/cmsCatArt/deleteCmsCatArtByIds',
    method: 'delete',
    data
  })
}

// @Tags CmsCatArt
// @Summary 更新CmsCatArt
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsCatArt true "更新CmsCatArt"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsCatArt/updateCmsCatArt [put]
export const updateCmsCatArt = (data) => {
  return service({
    url: '/cmsCatArt/updateCmsCatArt',
    method: 'put',
    data
  })
}

// @Tags CmsCatArt
// @Summary 用id查询CmsCatArt
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CmsCatArt true "用id查询CmsCatArt"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsCatArt/findCmsCatArt [get]
export const findCmsCatArt = (params) => {
  return service({
    url: '/cmsCatArt/findCmsCatArt',
    method: 'get',
    params
  })
}

// @Tags CmsCatArt
// @Summary 分页获取CmsCatArt列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CmsCatArt列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsCatArt/getCmsCatArtList [get]
export const getCmsCatArtList = (params) => {
  return service({
    url: '/cmsCatArt/getCmsCatArtList',
    method: 'get',
    params
  })
}


// @Tags CmsCatArt
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsCatArt true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsCatArt/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/cmsCatArt/quickEdit',
    method: 'post',
    data
  })
}