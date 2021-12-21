import service from '@/utils/request'

// @Tags ColHsj
// @Summary 创建ColHsj
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ColHsj true "创建ColHsj"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /colHsj/createColHsj [post]
export const createColHsj = (data) => {
  return service({
    url: '/colHsj/createColHsj',
    method: 'post',
    data
  })
}

// @Tags ColHsj
// @Summary 删除ColHsj
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ColHsj true "删除ColHsj"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /colHsj/deleteColHsj [delete]
export const deleteColHsj = (data) => {
  return service({
    url: '/colHsj/deleteColHsj',
    method: 'delete',
    data
  })
}

// @Tags ColHsj
// @Summary 删除ColHsj
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ColHsj"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /colHsj/deleteColHsj [delete]
export const deleteColHsjByIds = (data) => {
  return service({
    url: '/colHsj/deleteColHsjByIds',
    method: 'delete',
    data
  })
}

// @Tags ColHsj
// @Summary 更新ColHsj
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ColHsj true "更新ColHsj"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /colHsj/updateColHsj [put]
export const updateColHsj = (data) => {
  return service({
    url: '/colHsj/updateColHsj',
    method: 'put',
    data
  })
}

// @Tags ColHsj
// @Summary 用id查询ColHsj
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ColHsj true "用id查询ColHsj"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /colHsj/findColHsj [get]
export const findColHsj = (params) => {
  return service({
    url: '/colHsj/findColHsj',
    method: 'get',
    params
  })
}

// @Tags ColHsj
// @Summary 分页获取ColHsj列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ColHsj列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /colHsj/getColHsjList [get]
export const getColHsjList = (params) => {
  return service({
    url: '/colHsj/getColHsjList',
    method: 'get',
    params
  })
}


// @Tags ColHsj
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ColHsj true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /colHsj/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/colHsj/quickEdit',
    method: 'post',
    data
  })
}

// @Tags ColHsj
// @Summary 导出excelColHsj列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelColHsj列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /colHsj/getColHsjexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/colHsj/excelList',
    method: 'get',
    params
  })
}