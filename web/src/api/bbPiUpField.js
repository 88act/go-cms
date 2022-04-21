import service from '@/utils/request'

// @Tags BbPiUpField
// @Summary 创建BbPiUpField
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiUpField true "创建BbPiUpField"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiUpField/createBbPiUpField [post]
export const createBbPiUpField = (data) => {
  return service({
    url: '/bbPiUpField/createBbPiUpField',
    method: 'post',
    data
  })
}

// @Tags BbPiUpField
// @Summary 删除BbPiUpField
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiUpField true "删除BbPiUpField"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiUpField/deleteBbPiUpField [delete]
export const deleteBbPiUpField = (data) => {
  return service({
    url: '/bbPiUpField/deleteBbPiUpField',
    method: 'delete',
    data
  })
}

// @Tags BbPiUpField
// @Summary 删除BbPiUpField
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BbPiUpField"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiUpField/deleteBbPiUpField [delete]
export const deleteBbPiUpFieldByIds = (data) => {
  return service({
    url: '/bbPiUpField/deleteBbPiUpFieldByIds',
    method: 'delete',
    data
  })
}

// @Tags BbPiUpField
// @Summary 更新BbPiUpField
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiUpField true "更新BbPiUpField"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bbPiUpField/updateBbPiUpField [put]
export const updateBbPiUpField = (data) => {
  return service({
    url: '/bbPiUpField/updateBbPiUpField',
    method: 'put',
    data
  })
}

// @Tags BbPiUpField
// @Summary 用id查询BbPiUpField
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BbPiUpField true "用id查询BbPiUpField"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bbPiUpField/findBbPiUpField [get]
export const findBbPiUpField = (params) => {
  return service({
    url: '/bbPiUpField/findBbPiUpField',
    method: 'get',
    params
  })
}

// @Tags BbPiUpField
// @Summary 分页获取BbPiUpField列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BbPiUpField列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiUpField/getBbPiUpFieldList [get]
export const getBbPiUpFieldList = (params) => {
  return service({
    url: '/bbPiUpField/getBbPiUpFieldList',
    method: 'get',
    params
  })
}


// @Tags BbPiUpField
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiUpField true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiUpField/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/bbPiUpField/quickEdit',
    method: 'post',
    data
  })
}

// @Tags BbPiUpField
// @Summary 导出excelBbPiUpField列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelBbPiUpField列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiUpField/getBbPiUpFieldexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/bbPiUpField/excelList',
    method: 'get',
    params
  })
}