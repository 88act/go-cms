import service from '@/utils/request'

// @Tags MemUserSafe
// @Summary 创建MemUserSafe
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MemUserSafe true "创建MemUserSafe"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memUserSafe/createMemUserSafe [post]
export const createMemUserSafe = (data) => {
  return service({
    url: '/memUserSafe/createMemUserSafe',
    method: 'post',
    data
  })
}

// @Tags MemUserSafe
// @Summary 删除MemUserSafe
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MemUserSafe true "删除MemUserSafe"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /memUserSafe/deleteMemUserSafe [delete]
export const deleteMemUserSafe = (data) => {
  return service({
    url: '/memUserSafe/deleteMemUserSafe',
    method: 'delete',
    data
  })
}

// @Tags MemUserSafe
// @Summary 删除MemUserSafe
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MemUserSafe"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /memUserSafe/deleteMemUserSafe [delete]
export const deleteMemUserSafeByIds = (data) => {
  return service({
    url: '/memUserSafe/deleteMemUserSafeByIds',
    method: 'delete',
    data
  })
}

// @Tags MemUserSafe
// @Summary 更新MemUserSafe
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MemUserSafe true "更新MemUserSafe"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /memUserSafe/updateMemUserSafe [put]
export const updateMemUserSafe = (data) => {
  return service({
    url: '/memUserSafe/updateMemUserSafe',
    method: 'put',
    data
  })
}

// @Tags MemUserSafe
// @Summary 用id查询MemUserSafe
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MemUserSafe true "用id查询MemUserSafe"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /memUserSafe/findMemUserSafe [get]
export const findMemUserSafe = (params) => {
  return service({
    url: '/memUserSafe/findMemUserSafe',
    method: 'get',
    params
  })
}

// @Tags MemUserSafe
// @Summary 分页获取MemUserSafe列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取MemUserSafe列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memUserSafe/getMemUserSafeList [get]
export const getMemUserSafeList = (params) => {
  return service({
    url: '/memUserSafe/getMemUserSafeList',
    method: 'get',
    params
  })
}


// @Tags MemUserSafe
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MemUserSafe true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memUserSafe/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/memUserSafe/quickEdit',
    method: 'post',
    data
  })
}

// @Tags MemUserSafe
// @Summary 导出excelMemUserSafe列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelMemUserSafe列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memUserSafe/getMemUserSafeexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/memUserSafe/excelList',
    method: 'get',
    params
  })
}