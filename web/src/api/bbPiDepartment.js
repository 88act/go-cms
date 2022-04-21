import service from '@/utils/request'

// @Tags BbPiDepartment
// @Summary 创建BbPiDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiDepartment true "创建BbPiDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiDepartment/createBbPiDepartment [post]
export const createBbPiDepartment = (data) => {
  return service({
    url: '/bbPiDepartment/createBbPiDepartment',
    method: 'post',
    data
  })
}

// @Tags BbPiDepartment
// @Summary 删除BbPiDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiDepartment true "删除BbPiDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiDepartment/deleteBbPiDepartment [delete]
export const deleteBbPiDepartment = (data) => {
  return service({
    url: '/bbPiDepartment/deleteBbPiDepartment',
    method: 'delete',
    data
  })
}

// @Tags BbPiDepartment
// @Summary 删除BbPiDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BbPiDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiDepartment/deleteBbPiDepartment [delete]
export const deleteBbPiDepartmentByIds = (data) => {
  return service({
    url: '/bbPiDepartment/deleteBbPiDepartmentByIds',
    method: 'delete',
    data
  })
}

// @Tags BbPiDepartment
// @Summary 更新BbPiDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiDepartment true "更新BbPiDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bbPiDepartment/updateBbPiDepartment [put]
export const updateBbPiDepartment = (data) => {
  return service({
    url: '/bbPiDepartment/updateBbPiDepartment',
    method: 'put',
    data
  })
}

// @Tags BbPiDepartment
// @Summary 用id查询BbPiDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BbPiDepartment true "用id查询BbPiDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bbPiDepartment/findBbPiDepartment [get]
export const findBbPiDepartment = (params) => {
  return service({
    url: '/bbPiDepartment/findBbPiDepartment',
    method: 'get',
    params
  })
}

// @Tags BbPiDepartment
// @Summary 分页获取BbPiDepartment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BbPiDepartment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiDepartment/getBbPiDepartmentList [get]
export const getBbPiDepartmentList = (params) => {
  return service({
    url: '/bbPiDepartment/getBbPiDepartmentList',
    method: 'get',
    params
  })
}


// @Tags BbPiDepartment
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiDepartment true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiDepartment/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/bbPiDepartment/quickEdit',
    method: 'post',
    data
  })
}

// @Tags BbPiDepartment
// @Summary 导出excelBbPiDepartment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelBbPiDepartment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiDepartment/getBbPiDepartmentexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/bbPiDepartment/excelList',
    method: 'get',
    params
  })
}