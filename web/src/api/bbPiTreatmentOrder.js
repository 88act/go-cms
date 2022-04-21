import service from '@/utils/request'

// @Tags BbPiTreatmentOrder
// @Summary 创建BbPiTreatmentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiTreatmentOrder true "创建BbPiTreatmentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiTreatmentOrder/createBbPiTreatmentOrder [post]
export const createBbPiTreatmentOrder = (data) => {
  return service({
    url: '/bbPiTreatmentOrder/createBbPiTreatmentOrder',
    method: 'post',
    data
  })
}

// @Tags BbPiTreatmentOrder
// @Summary 删除BbPiTreatmentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiTreatmentOrder true "删除BbPiTreatmentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiTreatmentOrder/deleteBbPiTreatmentOrder [delete]
export const deleteBbPiTreatmentOrder = (data) => {
  return service({
    url: '/bbPiTreatmentOrder/deleteBbPiTreatmentOrder',
    method: 'delete',
    data
  })
}

// @Tags BbPiTreatmentOrder
// @Summary 删除BbPiTreatmentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BbPiTreatmentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiTreatmentOrder/deleteBbPiTreatmentOrder [delete]
export const deleteBbPiTreatmentOrderByIds = (data) => {
  return service({
    url: '/bbPiTreatmentOrder/deleteBbPiTreatmentOrderByIds',
    method: 'delete',
    data
  })
}

// @Tags BbPiTreatmentOrder
// @Summary 更新BbPiTreatmentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiTreatmentOrder true "更新BbPiTreatmentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bbPiTreatmentOrder/updateBbPiTreatmentOrder [put]
export const updateBbPiTreatmentOrder = (data) => {
  return service({
    url: '/bbPiTreatmentOrder/updateBbPiTreatmentOrder',
    method: 'put',
    data
  })
}

// @Tags BbPiTreatmentOrder
// @Summary 用id查询BbPiTreatmentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BbPiTreatmentOrder true "用id查询BbPiTreatmentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bbPiTreatmentOrder/findBbPiTreatmentOrder [get]
export const findBbPiTreatmentOrder = (params) => {
  return service({
    url: '/bbPiTreatmentOrder/findBbPiTreatmentOrder',
    method: 'get',
    params
  })
}

// @Tags BbPiTreatmentOrder
// @Summary 分页获取BbPiTreatmentOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BbPiTreatmentOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiTreatmentOrder/getBbPiTreatmentOrderList [get]
export const getBbPiTreatmentOrderList = (params) => {
  return service({
    url: '/bbPiTreatmentOrder/getBbPiTreatmentOrderList',
    method: 'get',
    params
  })
}


// @Tags BbPiTreatmentOrder
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiTreatmentOrder true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiTreatmentOrder/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/bbPiTreatmentOrder/quickEdit',
    method: 'post',
    data
  })
}

// @Tags BbPiTreatmentOrder
// @Summary 导出excelBbPiTreatmentOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelBbPiTreatmentOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiTreatmentOrder/getBbPiTreatmentOrderexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/bbPiTreatmentOrder/excelList',
    method: 'get',
    params
  })
}