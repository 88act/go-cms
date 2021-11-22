import service from '@/utils/request'

// @Tags ColCollect
// @Summary 创建ColCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ColCollect true "创建ColCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /colCollect/createColCollect [post]
export const createColCollect = (data) => {
  return service({
    url: '/colCollect/createColCollect',
    method: 'post',
    data
  })
}

// @Tags ColCollect
// @Summary 删除ColCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ColCollect true "删除ColCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /colCollect/deleteColCollect [delete]
export const deleteColCollect = (data) => {
  return service({
    url: '/colCollect/deleteColCollect',
    method: 'delete',
    data
  })
}

// @Tags ColCollect
// @Summary 删除ColCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ColCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /colCollect/deleteColCollect [delete]
export const deleteColCollectByIds = (data) => {
  return service({
    url: '/colCollect/deleteColCollectByIds',
    method: 'delete',
    data
  })
}

// @Tags ColCollect
// @Summary 更新ColCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ColCollect true "更新ColCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /colCollect/updateColCollect [put]
export const updateColCollect = (data) => {
  return service({
    url: '/colCollect/updateColCollect',
    method: 'put',
    data
  })
}

// @Tags ColCollect
// @Summary 用id查询ColCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ColCollect true "用id查询ColCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /colCollect/findColCollect [get]
export const findColCollect = (params) => {
  return service({
    url: '/colCollect/findColCollect',
    method: 'get',
    params
  })
}

// @Tags ColCollect
// @Summary 分页获取ColCollect列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ColCollect列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /colCollect/getColCollectList [get]
export const getColCollectList = (params) => {
  return service({
    url: '/colCollect/getColCollectList',
    method: 'get',
    params
  })
}


// @Tags ColCollect
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ColCollect true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /colCollect/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/colCollect/quickEdit',
    method: 'post',
    data
  })
}

// @Tags ColCollect
// @Summary 导出excelColCollect列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelColCollect列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /colCollect/getColCollectexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/colCollect/excelList',
    method: 'get',
    params
  })
}


// @Tags ColCollect
// @Summary 启动或停止采集器
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ColCollect true "启动或停止采集器"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /colCollect/startOrStopCollect [get]
export const startOrStopCollect = (params) => {
  return service({
    url: '/colCollect/startOrStopCollect',
    method: 'get',
    params
  })
}