import service from '@/utils/request'

// @Tags BasicArea
// @Summary 创建BasicArea
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BasicArea true "创建BasicArea"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /basicArea/createBasicArea [post]
export const createBasicArea = (data) => {
  return service({
    url: '/basicArea/createBasicArea',
    method: 'post',
    data
  })
}

// @Tags BasicArea
// @Summary 删除BasicArea
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BasicArea true "删除BasicArea"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /basicArea/deleteBasicArea [delete]
export const deleteBasicArea = (data) => {
  return service({
    url: '/basicArea/deleteBasicArea',
    method: 'delete',
    data
  })
}

// @Tags BasicArea
// @Summary 删除BasicArea
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BasicArea"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /basicArea/deleteBasicArea [delete]
export const deleteBasicAreaByIds = (data) => {
  return service({
    url: '/basicArea/deleteBasicAreaByIds',
    method: 'delete',
    data
  })
}

// @Tags BasicArea
// @Summary 更新BasicArea
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BasicArea true "更新BasicArea"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /basicArea/updateBasicArea [put]
export const updateBasicArea = (data) => {
  return service({
    url: '/basicArea/updateBasicArea',
    method: 'put',
    data
  })
}

// @Tags BasicArea
// @Summary 用id查询BasicArea
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BasicArea true "用id查询BasicArea"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /basicArea/findBasicArea [get]
export const findBasicArea = (params) => {
  return service({
    url: '/basicArea/findBasicArea',
    method: 'get',
    params
  })
}

// @Tags BasicArea
// @Summary 分页获取BasicArea列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BasicArea列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /basicArea/getBasicAreaList [get]
export const getBasicAreaList = (params) => {
  return service({
    url: '/basicArea/getBasicAreaList',
    method: 'get',
    params
  })
}


// @Tags BasicArea
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BasicArea true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /basicArea/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/basicArea/quickEdit',
    method: 'post',
    data
  })
}

// @Tags BasicArea
// @Summary 导出excelBasicArea列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelBasicArea列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /basicArea/getBasicAreaexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/basicArea/excelList',
    method: 'get',
    params
  })
}