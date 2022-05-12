import service from '@/utils/request'

// @Tags OrderOrder
// @Summary 创建OrderOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderOrder true "创建OrderOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderOrder/createOrderOrder [post]
export const createOrderOrder = (data) => {
  return service({
    url: '/orderOrder/createOrderOrder',
    method: 'post',
    data
  })
}

// @Tags OrderOrder
// @Summary 删除OrderOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderOrder true "删除OrderOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orderOrder/deleteOrderOrder [delete]
export const deleteOrderOrder = (data) => {
  return service({
    url: '/orderOrder/deleteOrderOrder',
    method: 'delete',
    data
  })
}

// @Tags OrderOrder
// @Summary 删除OrderOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OrderOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orderOrder/deleteOrderOrder [delete]
export const deleteOrderOrderByIds = (data) => {
  return service({
    url: '/orderOrder/deleteOrderOrderByIds',
    method: 'delete',
    data
  })
}

// @Tags OrderOrder
// @Summary 更新OrderOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderOrder true "更新OrderOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /orderOrder/updateOrderOrder [put]
export const updateOrderOrder = (data) => {
  return service({
    url: '/orderOrder/updateOrderOrder',
    method: 'put',
    data
  })
}

// @Tags OrderOrder
// @Summary 用id查询OrderOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OrderOrder true "用id查询OrderOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /orderOrder/findOrderOrder [get]
export const findOrderOrder = (params) => {
  return service({
    url: '/orderOrder/findOrderOrder',
    method: 'get',
    params
  })
}

// @Tags OrderOrder
// @Summary 分页获取OrderOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取OrderOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderOrder/getOrderOrderList [get]
export const getOrderOrderList = (params) => {
  return service({
    url: '/orderOrder/getOrderOrderList',
    method: 'get',
    params
  })
}


// @Tags OrderOrder
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderOrder true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderOrder/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/orderOrder/quickEdit',
    method: 'post',
    data
  })
}

// @Tags OrderOrder
// @Summary 导出excelOrderOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelOrderOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderOrder/getOrderOrderexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/orderOrder/excelList',
    method: 'get',
    params
  })
}