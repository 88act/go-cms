import service from '@/utils/request'

// @Tags BbPiStaff
// @Summary 创建BbPiStaff
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiStaff true "创建BbPiStaff"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiStaff/createBbPiStaff [post]
export const createBbPiStaff = (data) => {
  return service({
    url: '/bbPiStaff/createBbPiStaff',
    method: 'post',
    data
  })
}

// @Tags BbPiStaff
// @Summary 删除BbPiStaff
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiStaff true "删除BbPiStaff"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiStaff/deleteBbPiStaff [delete]
export const deleteBbPiStaff = (data) => {
  return service({
    url: '/bbPiStaff/deleteBbPiStaff',
    method: 'delete',
    data
  })
}

// @Tags BbPiStaff
// @Summary 删除BbPiStaff
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BbPiStaff"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiStaff/deleteBbPiStaff [delete]
export const deleteBbPiStaffByIds = (data) => {
  return service({
    url: '/bbPiStaff/deleteBbPiStaffByIds',
    method: 'delete',
    data
  })
}

// @Tags BbPiStaff
// @Summary 更新BbPiStaff
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiStaff true "更新BbPiStaff"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bbPiStaff/updateBbPiStaff [put]
export const updateBbPiStaff = (data) => {
  return service({
    url: '/bbPiStaff/updateBbPiStaff',
    method: 'put',
    data
  })
}

// @Tags BbPiStaff
// @Summary 用id查询BbPiStaff
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BbPiStaff true "用id查询BbPiStaff"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bbPiStaff/findBbPiStaff [get]
export const findBbPiStaff = (params) => {
  return service({
    url: '/bbPiStaff/findBbPiStaff',
    method: 'get',
    params
  })
}

// @Tags BbPiStaff
// @Summary 分页获取BbPiStaff列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BbPiStaff列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiStaff/getBbPiStaffList [get]
export const getBbPiStaffList = (params) => {
  return service({
    url: '/bbPiStaff/getBbPiStaffList',
    method: 'get',
    params
  })
}


// @Tags BbPiStaff
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiStaff true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiStaff/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/bbPiStaff/quickEdit',
    method: 'post',
    data
  })
}

// @Tags BbPiStaff
// @Summary 导出excelBbPiStaff列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelBbPiStaff列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiStaff/getBbPiStaffexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/bbPiStaff/excelList',
    method: 'get',
    params
  })
}