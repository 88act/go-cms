import service from '@/utils/request'

// @Tags BbPiDevice
// @Summary 创建BbPiDevice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiDevice true "创建BbPiDevice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiDevice/createBbPiDevice [post]
export const createBbPiDevice = (data) => {
  return service({
    url: '/bbPiDevice/createBbPiDevice',
    method: 'post',
    data
  })
}

// @Tags BbPiDevice
// @Summary 删除BbPiDevice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiDevice true "删除BbPiDevice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiDevice/deleteBbPiDevice [delete]
export const deleteBbPiDevice = (data) => {
  return service({
    url: '/bbPiDevice/deleteBbPiDevice',
    method: 'delete',
    data
  })
}

// @Tags BbPiDevice
// @Summary 删除BbPiDevice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BbPiDevice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiDevice/deleteBbPiDevice [delete]
export const deleteBbPiDeviceByIds = (data) => {
  return service({
    url: '/bbPiDevice/deleteBbPiDeviceByIds',
    method: 'delete',
    data
  })
}

// @Tags BbPiDevice
// @Summary 更新BbPiDevice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiDevice true "更新BbPiDevice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bbPiDevice/updateBbPiDevice [put]
export const updateBbPiDevice = (data) => {
  return service({
    url: '/bbPiDevice/updateBbPiDevice',
    method: 'put',
    data
  })
}

// @Tags BbPiDevice
// @Summary 用id查询BbPiDevice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BbPiDevice true "用id查询BbPiDevice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bbPiDevice/findBbPiDevice [get]
export const findBbPiDevice = (params) => {
  return service({
    url: '/bbPiDevice/findBbPiDevice',
    method: 'get',
    params
  })
}

// @Tags BbPiDevice
// @Summary 分页获取BbPiDevice列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BbPiDevice列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiDevice/getBbPiDeviceList [get]
export const getBbPiDeviceList = (params) => {
  return service({
    url: '/bbPiDevice/getBbPiDeviceList',
    method: 'get',
    params
  })
}


// @Tags BbPiDevice
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiDevice true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiDevice/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/bbPiDevice/quickEdit',
    method: 'post',
    data
  })
}

// @Tags BbPiDevice
// @Summary 导出excelBbPiDevice列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelBbPiDevice列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiDevice/getBbPiDeviceexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/bbPiDevice/excelList',
    method: 'get',
    params
  })
}