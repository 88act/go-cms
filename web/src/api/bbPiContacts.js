import service from '@/utils/request'

// @Tags BbPiContacts
// @Summary 创建BbPiContacts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiContacts true "创建BbPiContacts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiContacts/createBbPiContacts [post]
export const createBbPiContacts = (data) => {
  return service({
    url: '/bbPiContacts/createBbPiContacts',
    method: 'post',
    data
  })
}

// @Tags BbPiContacts
// @Summary 删除BbPiContacts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiContacts true "删除BbPiContacts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiContacts/deleteBbPiContacts [delete]
export const deleteBbPiContacts = (data) => {
  return service({
    url: '/bbPiContacts/deleteBbPiContacts',
    method: 'delete',
    data
  })
}

// @Tags BbPiContacts
// @Summary 删除BbPiContacts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BbPiContacts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiContacts/deleteBbPiContacts [delete]
export const deleteBbPiContactsByIds = (data) => {
  return service({
    url: '/bbPiContacts/deleteBbPiContactsByIds',
    method: 'delete',
    data
  })
}

// @Tags BbPiContacts
// @Summary 更新BbPiContacts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiContacts true "更新BbPiContacts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bbPiContacts/updateBbPiContacts [put]
export const updateBbPiContacts = (data) => {
  return service({
    url: '/bbPiContacts/updateBbPiContacts',
    method: 'put',
    data
  })
}

// @Tags BbPiContacts
// @Summary 用id查询BbPiContacts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BbPiContacts true "用id查询BbPiContacts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bbPiContacts/findBbPiContacts [get]
export const findBbPiContacts = (params) => {
  return service({
    url: '/bbPiContacts/findBbPiContacts',
    method: 'get',
    params
  })
}

// @Tags BbPiContacts
// @Summary 分页获取BbPiContacts列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BbPiContacts列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiContacts/getBbPiContactsList [get]
export const getBbPiContactsList = (params) => {
  return service({
    url: '/bbPiContacts/getBbPiContactsList',
    method: 'get',
    params
  })
}


// @Tags BbPiContacts
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiContacts true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiContacts/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/bbPiContacts/quickEdit',
    method: 'post',
    data
  })
}

// @Tags BbPiContacts
// @Summary 导出excelBbPiContacts列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelBbPiContacts列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiContacts/getBbPiContactsexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/bbPiContacts/excelList',
    method: 'get',
    params
  })
}