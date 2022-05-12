import service from '@/utils/request'

// @Tags MemUserLog
// @Summary 创建MemUserLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MemUserLog true "创建MemUserLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memUserLog/createMemUserLog [post]
export const createMemUserLog = (data) => {
  return service({
    url: '/memUserLog/createMemUserLog',
    method: 'post',
    data
  })
}

// @Tags MemUserLog
// @Summary 删除MemUserLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MemUserLog true "删除MemUserLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /memUserLog/deleteMemUserLog [delete]
export const deleteMemUserLog = (data) => {
  return service({
    url: '/memUserLog/deleteMemUserLog',
    method: 'delete',
    data
  })
}

// @Tags MemUserLog
// @Summary 删除MemUserLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MemUserLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /memUserLog/deleteMemUserLog [delete]
export const deleteMemUserLogByIds = (data) => {
  return service({
    url: '/memUserLog/deleteMemUserLogByIds',
    method: 'delete',
    data
  })
}

// @Tags MemUserLog
// @Summary 更新MemUserLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MemUserLog true "更新MemUserLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /memUserLog/updateMemUserLog [put]
export const updateMemUserLog = (data) => {
  return service({
    url: '/memUserLog/updateMemUserLog',
    method: 'put',
    data
  })
}

// @Tags MemUserLog
// @Summary 用id查询MemUserLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MemUserLog true "用id查询MemUserLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /memUserLog/findMemUserLog [get]
export const findMemUserLog = (params) => {
  return service({
    url: '/memUserLog/findMemUserLog',
    method: 'get',
    params
  })
}

// @Tags MemUserLog
// @Summary 分页获取MemUserLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取MemUserLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memUserLog/getMemUserLogList [get]
export const getMemUserLogList = (params) => {
  return service({
    url: '/memUserLog/getMemUserLogList',
    method: 'get',
    params
  })
}


// @Tags MemUserLog
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MemUserLog true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memUserLog/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/memUserLog/quickEdit',
    method: 'post',
    data
  })
}

// @Tags MemUserLog
// @Summary 导出excelMemUserLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelMemUserLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memUserLog/getMemUserLogexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/memUserLog/excelList',
    method: 'get',
    params
  })
}