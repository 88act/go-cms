import service from '@/utils/request'

// @Tags MemUser
// @Summary 创建MemUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MemUser true "创建MemUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memUser/createMemUser [post]
export const createMemUser = (data) => {
  return service({
    url: '/memUser/createMemUser',
    method: 'post',
    data
  })
}

// @Tags MemUser
// @Summary 删除MemUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MemUser true "删除MemUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /memUser/deleteMemUser [delete]
export const deleteMemUser = (data) => {
  return service({
    url: '/memUser/deleteMemUser',
    method: 'delete',
    data
  })
}

// @Tags MemUser
// @Summary 删除MemUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MemUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /memUser/deleteMemUser [delete]
export const deleteMemUserByIds = (data) => {
  return service({
    url: '/memUser/deleteMemUserByIds',
    method: 'delete',
    data
  })
}

// @Tags MemUser
// @Summary 更新MemUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MemUser true "更新MemUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /memUser/updateMemUser [put]
export const updateMemUser = (data) => {
  return service({
    url: '/memUser/updateMemUser',
    method: 'put',
    data
  })
}

// @Tags MemUser
// @Summary 用id查询MemUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MemUser true "用id查询MemUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /memUser/findMemUser [get]
export const findMemUser = (params) => {
  return service({
    url: '/memUser/findMemUser',
    method: 'get',
    params
  })
}

// @Tags MemUser
// @Summary 分页获取MemUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取MemUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memUser/getMemUserList [get]
export const getMemUserList = (params) => {
  return service({
    url: '/memUser/getMemUserList',
    method: 'get',
    params
  })
}


// @Tags MemUser
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MemUser true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memUser/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/memUser/quickEdit',
    method: 'post',
    data
  })
}

// @Tags MemUser
// @Summary 导出excelMemUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelMemUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memUser/getMemUserexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/memUser/excelList',
    method: 'get',
    params
  })
}