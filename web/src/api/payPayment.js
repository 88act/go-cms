import service from '@/utils/request'

// @Tags PayPayment
// @Summary 创建PayPayment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayPayment true "创建PayPayment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payPayment/createPayPayment [post]
export const createPayPayment = (data) => {
  return service({
    url: '/payPayment/createPayPayment',
    method: 'post',
    data
  })
}

// @Tags PayPayment
// @Summary 删除PayPayment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayPayment true "删除PayPayment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payPayment/deletePayPayment [delete]
export const deletePayPayment = (data) => {
  return service({
    url: '/payPayment/deletePayPayment',
    method: 'delete',
    data
  })
}

// @Tags PayPayment
// @Summary 删除PayPayment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PayPayment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payPayment/deletePayPayment [delete]
export const deletePayPaymentByIds = (data) => {
  return service({
    url: '/payPayment/deletePayPaymentByIds',
    method: 'delete',
    data
  })
}

// @Tags PayPayment
// @Summary 更新PayPayment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayPayment true "更新PayPayment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /payPayment/updatePayPayment [put]
export const updatePayPayment = (data) => {
  return service({
    url: '/payPayment/updatePayPayment',
    method: 'put',
    data
  })
}

// @Tags PayPayment
// @Summary 用id查询PayPayment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PayPayment true "用id查询PayPayment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /payPayment/findPayPayment [get]
export const findPayPayment = (params) => {
  return service({
    url: '/payPayment/findPayPayment',
    method: 'get',
    params
  })
}

// @Tags PayPayment
// @Summary 分页获取PayPayment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取PayPayment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payPayment/getPayPaymentList [get]
export const getPayPaymentList = (params) => {
  return service({
    url: '/payPayment/getPayPaymentList',
    method: 'get',
    params
  })
}


// @Tags PayPayment
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayPayment true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payPayment/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/payPayment/quickEdit',
    method: 'post',
    data
  })
}

// @Tags PayPayment
// @Summary 导出excelPayPayment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelPayPayment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payPayment/getPayPaymentexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/payPayment/excelList',
    method: 'get',
    params
  })
}