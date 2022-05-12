import service from '@/utils/request'

// @Tags MemAddress
// @Summary 创建MemAddress
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MemAddress true "创建MemAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memAddress/createMemAddress [post]
export const createMemAddress = (data) => {
  return service({
    url: '/memAddress/createMemAddress',
    method: 'post',
    data
  })
}

// @Tags MemAddress
// @Summary 删除MemAddress
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MemAddress true "删除MemAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /memAddress/deleteMemAddress [delete]
export const deleteMemAddress = (data) => {
  return service({
    url: '/memAddress/deleteMemAddress',
    method: 'delete',
    data
  })
}

// @Tags MemAddress
// @Summary 删除MemAddress
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MemAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /memAddress/deleteMemAddress [delete]
export const deleteMemAddressByIds = (data) => {
  return service({
    url: '/memAddress/deleteMemAddressByIds',
    method: 'delete',
    data
  })
}

// @Tags MemAddress
// @Summary 更新MemAddress
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MemAddress true "更新MemAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /memAddress/updateMemAddress [put]
export const updateMemAddress = (data) => {
  return service({
    url: '/memAddress/updateMemAddress',
    method: 'put',
    data
  })
}

// @Tags MemAddress
// @Summary 用id查询MemAddress
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MemAddress true "用id查询MemAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /memAddress/findMemAddress [get]
export const findMemAddress = (params) => {
  return service({
    url: '/memAddress/findMemAddress',
    method: 'get',
    params
  })
}

// @Tags MemAddress
// @Summary 分页获取MemAddress列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取MemAddress列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memAddress/getMemAddressList [get]
export const getMemAddressList = (params) => {
  return service({
    url: '/memAddress/getMemAddressList',
    method: 'get',
    params
  })
}


// @Tags MemAddress
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MemAddress true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memAddress/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/memAddress/quickEdit',
    method: 'post',
    data
  })
}

// @Tags MemAddress
// @Summary 导出excelMemAddress列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelMemAddress列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memAddress/getMemAddressexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/memAddress/excelList',
    method: 'get',
    params
  })
}