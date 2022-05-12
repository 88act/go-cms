import service from '@/utils/request'

// @Tags ActSelect
// @Summary 创建ActSelect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActSelect true "创建ActSelect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actSelect/createActSelect [post]
export const createActSelect = (data) => {
  return service({
    url: '/actSelect/createActSelect',
    method: 'post',
    data
  })
}

// @Tags ActSelect
// @Summary 删除ActSelect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActSelect true "删除ActSelect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /actSelect/deleteActSelect [delete]
export const deleteActSelect = (data) => {
  return service({
    url: '/actSelect/deleteActSelect',
    method: 'delete',
    data
  })
}

// @Tags ActSelect
// @Summary 删除ActSelect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ActSelect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /actSelect/deleteActSelect [delete]
export const deleteActSelectByIds = (data) => {
  return service({
    url: '/actSelect/deleteActSelectByIds',
    method: 'delete',
    data
  })
}

// @Tags ActSelect
// @Summary 更新ActSelect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActSelect true "更新ActSelect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /actSelect/updateActSelect [put]
export const updateActSelect = (data) => {
  return service({
    url: '/actSelect/updateActSelect',
    method: 'put',
    data
  })
}

// @Tags ActSelect
// @Summary 用id查询ActSelect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ActSelect true "用id查询ActSelect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /actSelect/findActSelect [get]
export const findActSelect = (params) => {
  return service({
    url: '/actSelect/findActSelect',
    method: 'get',
    params
  })
}

// @Tags ActSelect
// @Summary 分页获取ActSelect列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ActSelect列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actSelect/getActSelectList [get]
export const getActSelectList = (params) => {
  return service({
    url: '/actSelect/getActSelectList',
    method: 'get',
    params
  })
}


// @Tags ActSelect
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActSelect true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actSelect/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/actSelect/quickEdit',
    method: 'post',
    data
  })
}

// @Tags ActSelect
// @Summary 导出excelActSelect列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelActSelect列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actSelect/getActSelectexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/actSelect/excelList',
    method: 'get',
    params
  })
}