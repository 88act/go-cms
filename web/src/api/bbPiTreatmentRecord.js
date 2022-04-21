import service from '@/utils/request'

// @Tags BbPiTreatmentRecord
// @Summary 创建BbPiTreatmentRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiTreatmentRecord true "创建BbPiTreatmentRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiTreatmentRecord/createBbPiTreatmentRecord [post]
export const createBbPiTreatmentRecord = (data) => {
  return service({
    url: '/bbPiTreatmentRecord/createBbPiTreatmentRecord',
    method: 'post',
    data
  })
}

// @Tags BbPiTreatmentRecord
// @Summary 删除BbPiTreatmentRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiTreatmentRecord true "删除BbPiTreatmentRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiTreatmentRecord/deleteBbPiTreatmentRecord [delete]
export const deleteBbPiTreatmentRecord = (data) => {
  return service({
    url: '/bbPiTreatmentRecord/deleteBbPiTreatmentRecord',
    method: 'delete',
    data
  })
}

// @Tags BbPiTreatmentRecord
// @Summary 删除BbPiTreatmentRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BbPiTreatmentRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiTreatmentRecord/deleteBbPiTreatmentRecord [delete]
export const deleteBbPiTreatmentRecordByIds = (data) => {
  return service({
    url: '/bbPiTreatmentRecord/deleteBbPiTreatmentRecordByIds',
    method: 'delete',
    data
  })
}

// @Tags BbPiTreatmentRecord
// @Summary 更新BbPiTreatmentRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiTreatmentRecord true "更新BbPiTreatmentRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bbPiTreatmentRecord/updateBbPiTreatmentRecord [put]
export const updateBbPiTreatmentRecord = (data) => {
  return service({
    url: '/bbPiTreatmentRecord/updateBbPiTreatmentRecord',
    method: 'put',
    data
  })
}

// @Tags BbPiTreatmentRecord
// @Summary 用id查询BbPiTreatmentRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BbPiTreatmentRecord true "用id查询BbPiTreatmentRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bbPiTreatmentRecord/findBbPiTreatmentRecord [get]
export const findBbPiTreatmentRecord = (params) => {
  return service({
    url: '/bbPiTreatmentRecord/findBbPiTreatmentRecord',
    method: 'get',
    params
  })
}

// @Tags BbPiTreatmentRecord
// @Summary 分页获取BbPiTreatmentRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BbPiTreatmentRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiTreatmentRecord/getBbPiTreatmentRecordList [get]
export const getBbPiTreatmentRecordList = (params) => {
  return service({
    url: '/bbPiTreatmentRecord/getBbPiTreatmentRecordList',
    method: 'get',
    params
  })
}


// @Tags BbPiTreatmentRecord
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiTreatmentRecord true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiTreatmentRecord/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/bbPiTreatmentRecord/quickEdit',
    method: 'post',
    data
  })
}

// @Tags BbPiTreatmentRecord
// @Summary 导出excelBbPiTreatmentRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelBbPiTreatmentRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiTreatmentRecord/getBbPiTreatmentRecordexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/bbPiTreatmentRecord/excelList',
    method: 'get',
    params
  })
}