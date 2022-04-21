import service from '@/utils/request'

// @Tags BbPiInstitution
// @Summary 创建BbPiInstitution
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiInstitution true "创建BbPiInstitution"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiInstitution/createBbPiInstitution [post]
export const createBbPiInstitution = (data) => {
  return service({
    url: '/bbPiInstitution/createBbPiInstitution',
    method: 'post',
    data
  })
}

// @Tags BbPiInstitution
// @Summary 删除BbPiInstitution
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiInstitution true "删除BbPiInstitution"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiInstitution/deleteBbPiInstitution [delete]
export const deleteBbPiInstitution = (data) => {
  return service({
    url: '/bbPiInstitution/deleteBbPiInstitution',
    method: 'delete',
    data
  })
}

// @Tags BbPiInstitution
// @Summary 删除BbPiInstitution
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BbPiInstitution"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiInstitution/deleteBbPiInstitution [delete]
export const deleteBbPiInstitutionByIds = (data) => {
  return service({
    url: '/bbPiInstitution/deleteBbPiInstitutionByIds',
    method: 'delete',
    data
  })
}

// @Tags BbPiInstitution
// @Summary 更新BbPiInstitution
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiInstitution true "更新BbPiInstitution"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bbPiInstitution/updateBbPiInstitution [put]
export const updateBbPiInstitution = (data) => {
  return service({
    url: '/bbPiInstitution/updateBbPiInstitution',
    method: 'put',
    data
  })
}

// @Tags BbPiInstitution
// @Summary 用id查询BbPiInstitution
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BbPiInstitution true "用id查询BbPiInstitution"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bbPiInstitution/findBbPiInstitution [get]
export const findBbPiInstitution = (params) => {
  return service({
    url: '/bbPiInstitution/findBbPiInstitution',
    method: 'get',
    params
  })
}

// @Tags BbPiInstitution
// @Summary 分页获取BbPiInstitution列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BbPiInstitution列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiInstitution/getBbPiInstitutionList [get]
export const getBbPiInstitutionList = (params) => {
  return service({
    url: '/bbPiInstitution/getBbPiInstitutionList',
    method: 'get',
    params
  })
}


// @Tags BbPiInstitution
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiInstitution true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiInstitution/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/bbPiInstitution/quickEdit',
    method: 'post',
    data
  })
}

// @Tags BbPiInstitution
// @Summary 导出excelBbPiInstitution列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelBbPiInstitution列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiInstitution/getBbPiInstitutionexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/bbPiInstitution/excelList',
    method: 'get',
    params
  })
}