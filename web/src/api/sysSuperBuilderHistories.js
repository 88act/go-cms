import service from '@/utils/request'

// @Tags SysSuperBuilderHistories
// @Summary 创建SysSuperBuilderHistories
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysSuperBuilderHistories true "创建SysSuperBuilderHistories"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysSuperBuilderHistories/createSysSuperBuilderHistories [post]
export const createSysSuperBuilderHistories = (data) => {
  return service({
    url: '/sysSuperBuilderHistories/createSysSuperBuilderHistories',
    method: 'post',
    data
  })
}

// @Tags SysSuperBuilderHistories
// @Summary 删除SysSuperBuilderHistories
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysSuperBuilderHistories true "删除SysSuperBuilderHistories"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysSuperBuilderHistories/deleteSysSuperBuilderHistories [delete]
export const deleteSysSuperBuilderHistories = (data) => {
  return service({
    url: '/sysSuperBuilderHistories/deleteSysSuperBuilderHistories',
    method: 'delete',
    data
  })
}

// @Tags SysSuperBuilderHistories
// @Summary 删除SysSuperBuilderHistories
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysSuperBuilderHistories"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysSuperBuilderHistories/deleteSysSuperBuilderHistories [delete]
export const deleteSysSuperBuilderHistoriesByIds = (data) => {
  return service({
    url: '/sysSuperBuilderHistories/deleteSysSuperBuilderHistoriesByIds',
    method: 'delete',
    data
  })
}

// @Tags SysSuperBuilderHistories
// @Summary 更新SysSuperBuilderHistories
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysSuperBuilderHistories true "更新SysSuperBuilderHistories"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysSuperBuilderHistories/updateSysSuperBuilderHistories [put]
export const updateSysSuperBuilderHistories = (data) => {
  return service({
    url: '/sysSuperBuilderHistories/updateSysSuperBuilderHistories',
    method: 'put',
    data
  })
}

// @Tags SysSuperBuilderHistories
// @Summary 用id查询SysSuperBuilderHistories
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SysSuperBuilderHistories true "用id查询SysSuperBuilderHistories"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysSuperBuilderHistories/findSysSuperBuilderHistories [get]
export const findSysSuperBuilderHistories = (params) => {
  return service({
    url: '/sysSuperBuilderHistories/findSysSuperBuilderHistories',
    method: 'get',
    params
  })
}

// @Tags SysSuperBuilderHistories
// @Summary 分页获取SysSuperBuilderHistories列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取SysSuperBuilderHistories列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysSuperBuilderHistories/getSysSuperBuilderHistoriesList [get]
export const getSysSuperBuilderHistoriesList = (params) => {
  return service({
    url: '/sysSuperBuilderHistories/getSysSuperBuilderHistoriesList',
    method: 'get',
    params
  })
}


// @Tags SysSuperBuilderHistories
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysSuperBuilderHistories true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysSuperBuilderHistories/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/sysSuperBuilderHistories/quickEdit',
    method: 'post',
    data
  })
}

// @Tags SysSuperBuilderHistories
// @Summary 导出excelSysSuperBuilderHistories列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelSysSuperBuilderHistories列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysSuperBuilderHistories/getSysSuperBuilderHistoriesexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/sysSuperBuilderHistories/excelList',
    method: 'get',
    params
  })
}