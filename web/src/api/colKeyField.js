import service from '@/utils/request'

// @Tags ColKeyField
// @Summary 创建ColKeyField
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ColKeyField true "创建ColKeyField"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /colKeyField/createColKeyField [post]
export const createColKeyField = (data) => {
  return service({
    url: '/colKeyField/createColKeyField',
    method: 'post',
    data
  })
}

// @Tags ColKeyField
// @Summary 删除ColKeyField
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ColKeyField true "删除ColKeyField"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /colKeyField/deleteColKeyField [delete]
export const deleteColKeyField = (data) => {
  return service({
    url: '/colKeyField/deleteColKeyField',
    method: 'delete',
    data
  })
}

// @Tags ColKeyField
// @Summary 删除ColKeyField
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ColKeyField"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /colKeyField/deleteColKeyField [delete]
export const deleteColKeyFieldByIds = (data) => {
  return service({
    url: '/colKeyField/deleteColKeyFieldByIds',
    method: 'delete',
    data
  })
}

// @Tags ColKeyField
// @Summary 更新ColKeyField
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ColKeyField true "更新ColKeyField"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /colKeyField/updateColKeyField [put]
export const updateColKeyField = (data) => {
  return service({
    url: '/colKeyField/updateColKeyField',
    method: 'put',
    data
  })
}

// @Tags ColKeyField
// @Summary 用id查询ColKeyField
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ColKeyField true "用id查询ColKeyField"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /colKeyField/findColKeyField [get]
export const findColKeyField = (params) => {
  return service({
    url: '/colKeyField/findColKeyField',
    method: 'get',
    params
  })
}

// @Tags ColKeyField
// @Summary 分页获取ColKeyField列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ColKeyField列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /colKeyField/getColKeyFieldList [get]
export const getColKeyFieldList = (params) => {
  return service({
    url: '/colKeyField/getColKeyFieldList',
    method: 'get',
    params
  })
}


// @Tags ColKeyField
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ColKeyField true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /colKeyField/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/colKeyField/quickEdit',
    method: 'post',
    data
  })
}

// @Tags ColKeyField
// @Summary 导出excelColKeyField列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelColKeyField列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /colKeyField/getColKeyFieldexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/colKeyField/excelList',
    method: 'get',
    params
  })
}