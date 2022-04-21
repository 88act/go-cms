import service from '@/utils/request'

// @Tags BbPiTreatmentDiagnose
// @Summary 创建BbPiTreatmentDiagnose
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiTreatmentDiagnose true "创建BbPiTreatmentDiagnose"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiTreatmentDiagnose/createBbPiTreatmentDiagnose [post]
export const createBbPiTreatmentDiagnose = (data) => {
  return service({
    url: '/bbPiTreatmentDiagnose/createBbPiTreatmentDiagnose',
    method: 'post',
    data
  })
}

// @Tags BbPiTreatmentDiagnose
// @Summary 删除BbPiTreatmentDiagnose
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiTreatmentDiagnose true "删除BbPiTreatmentDiagnose"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiTreatmentDiagnose/deleteBbPiTreatmentDiagnose [delete]
export const deleteBbPiTreatmentDiagnose = (data) => {
  return service({
    url: '/bbPiTreatmentDiagnose/deleteBbPiTreatmentDiagnose',
    method: 'delete',
    data
  })
}

// @Tags BbPiTreatmentDiagnose
// @Summary 删除BbPiTreatmentDiagnose
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BbPiTreatmentDiagnose"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiTreatmentDiagnose/deleteBbPiTreatmentDiagnose [delete]
export const deleteBbPiTreatmentDiagnoseByIds = (data) => {
  return service({
    url: '/bbPiTreatmentDiagnose/deleteBbPiTreatmentDiagnoseByIds',
    method: 'delete',
    data
  })
}

// @Tags BbPiTreatmentDiagnose
// @Summary 更新BbPiTreatmentDiagnose
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiTreatmentDiagnose true "更新BbPiTreatmentDiagnose"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bbPiTreatmentDiagnose/updateBbPiTreatmentDiagnose [put]
export const updateBbPiTreatmentDiagnose = (data) => {
  return service({
    url: '/bbPiTreatmentDiagnose/updateBbPiTreatmentDiagnose',
    method: 'put',
    data
  })
}

// @Tags BbPiTreatmentDiagnose
// @Summary 用id查询BbPiTreatmentDiagnose
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BbPiTreatmentDiagnose true "用id查询BbPiTreatmentDiagnose"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bbPiTreatmentDiagnose/findBbPiTreatmentDiagnose [get]
export const findBbPiTreatmentDiagnose = (params) => {
  return service({
    url: '/bbPiTreatmentDiagnose/findBbPiTreatmentDiagnose',
    method: 'get',
    params
  })
}

// @Tags BbPiTreatmentDiagnose
// @Summary 分页获取BbPiTreatmentDiagnose列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BbPiTreatmentDiagnose列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiTreatmentDiagnose/getBbPiTreatmentDiagnoseList [get]
export const getBbPiTreatmentDiagnoseList = (params) => {
  return service({
    url: '/bbPiTreatmentDiagnose/getBbPiTreatmentDiagnoseList',
    method: 'get',
    params
  })
}


// @Tags BbPiTreatmentDiagnose
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiTreatmentDiagnose true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiTreatmentDiagnose/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/bbPiTreatmentDiagnose/quickEdit',
    method: 'post',
    data
  })
}

// @Tags BbPiTreatmentDiagnose
// @Summary 导出excelBbPiTreatmentDiagnose列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelBbPiTreatmentDiagnose列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiTreatmentDiagnose/getBbPiTreatmentDiagnoseexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/bbPiTreatmentDiagnose/excelList',
    method: 'get',
    params
  })
}