import service from '@/utils/request'

// @Tags BbPiTreatmentReferral
// @Summary 创建BbPiTreatmentReferral
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiTreatmentReferral true "创建BbPiTreatmentReferral"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiTreatmentReferral/createBbPiTreatmentReferral [post]
export const createBbPiTreatmentReferral = (data) => {
  return service({
    url: '/bbPiTreatmentReferral/createBbPiTreatmentReferral',
    method: 'post',
    data
  })
}

// @Tags BbPiTreatmentReferral
// @Summary 删除BbPiTreatmentReferral
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiTreatmentReferral true "删除BbPiTreatmentReferral"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiTreatmentReferral/deleteBbPiTreatmentReferral [delete]
export const deleteBbPiTreatmentReferral = (data) => {
  return service({
    url: '/bbPiTreatmentReferral/deleteBbPiTreatmentReferral',
    method: 'delete',
    data
  })
}

// @Tags BbPiTreatmentReferral
// @Summary 删除BbPiTreatmentReferral
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BbPiTreatmentReferral"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiTreatmentReferral/deleteBbPiTreatmentReferral [delete]
export const deleteBbPiTreatmentReferralByIds = (data) => {
  return service({
    url: '/bbPiTreatmentReferral/deleteBbPiTreatmentReferralByIds',
    method: 'delete',
    data
  })
}

// @Tags BbPiTreatmentReferral
// @Summary 更新BbPiTreatmentReferral
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiTreatmentReferral true "更新BbPiTreatmentReferral"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bbPiTreatmentReferral/updateBbPiTreatmentReferral [put]
export const updateBbPiTreatmentReferral = (data) => {
  return service({
    url: '/bbPiTreatmentReferral/updateBbPiTreatmentReferral',
    method: 'put',
    data
  })
}

// @Tags BbPiTreatmentReferral
// @Summary 用id查询BbPiTreatmentReferral
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BbPiTreatmentReferral true "用id查询BbPiTreatmentReferral"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bbPiTreatmentReferral/findBbPiTreatmentReferral [get]
export const findBbPiTreatmentReferral = (params) => {
  return service({
    url: '/bbPiTreatmentReferral/findBbPiTreatmentReferral',
    method: 'get',
    params
  })
}

// @Tags BbPiTreatmentReferral
// @Summary 分页获取BbPiTreatmentReferral列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BbPiTreatmentReferral列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiTreatmentReferral/getBbPiTreatmentReferralList [get]
export const getBbPiTreatmentReferralList = (params) => {
  return service({
    url: '/bbPiTreatmentReferral/getBbPiTreatmentReferralList',
    method: 'get',
    params
  })
}


// @Tags BbPiTreatmentReferral
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiTreatmentReferral true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiTreatmentReferral/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/bbPiTreatmentReferral/quickEdit',
    method: 'post',
    data
  })
}

// @Tags BbPiTreatmentReferral
// @Summary 导出excelBbPiTreatmentReferral列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelBbPiTreatmentReferral列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiTreatmentReferral/getBbPiTreatmentReferralexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/bbPiTreatmentReferral/excelList',
    method: 'get',
    params
  })
}