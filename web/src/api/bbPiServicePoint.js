import service from '@/utils/request'

// @Tags BbPiServicePoint
// @Summary 创建BbPiServicePoint
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiServicePoint true "创建BbPiServicePoint"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiServicePoint/createBbPiServicePoint [post]
export const createBbPiServicePoint = (data) => {
  return service({
    url: '/bbPiServicePoint/createBbPiServicePoint',
    method: 'post',
    data
  })
}

// @Tags BbPiServicePoint
// @Summary 删除BbPiServicePoint
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiServicePoint true "删除BbPiServicePoint"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiServicePoint/deleteBbPiServicePoint [delete]
export const deleteBbPiServicePoint = (data) => {
  return service({
    url: '/bbPiServicePoint/deleteBbPiServicePoint',
    method: 'delete',
    data
  })
}

// @Tags BbPiServicePoint
// @Summary 删除BbPiServicePoint
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BbPiServicePoint"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiServicePoint/deleteBbPiServicePoint [delete]
export const deleteBbPiServicePointByIds = (data) => {
  return service({
    url: '/bbPiServicePoint/deleteBbPiServicePointByIds',
    method: 'delete',
    data
  })
}

// @Tags BbPiServicePoint
// @Summary 更新BbPiServicePoint
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiServicePoint true "更新BbPiServicePoint"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bbPiServicePoint/updateBbPiServicePoint [put]
export const updateBbPiServicePoint = (data) => {
  return service({
    url: '/bbPiServicePoint/updateBbPiServicePoint',
    method: 'put',
    data
  })
}

// @Tags BbPiServicePoint
// @Summary 用id查询BbPiServicePoint
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BbPiServicePoint true "用id查询BbPiServicePoint"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bbPiServicePoint/findBbPiServicePoint [get]
export const findBbPiServicePoint = (params) => {
  return service({
    url: '/bbPiServicePoint/findBbPiServicePoint',
    method: 'get',
    params
  })
}

// @Tags BbPiServicePoint
// @Summary 分页获取BbPiServicePoint列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BbPiServicePoint列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiServicePoint/getBbPiServicePointList [get]
export const getBbPiServicePointList = (params) => {
  return service({
    url: '/bbPiServicePoint/getBbPiServicePointList',
    method: 'get',
    params
  })
}


// @Tags BbPiServicePoint
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiServicePoint true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiServicePoint/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/bbPiServicePoint/quickEdit',
    method: 'post',
    data
  })
}

// @Tags BbPiServicePoint
// @Summary 导出excelBbPiServicePoint列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelBbPiServicePoint列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiServicePoint/getBbPiServicePointexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/bbPiServicePoint/excelList',
    method: 'get',
    params
  })
}