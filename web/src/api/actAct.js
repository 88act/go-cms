import service from '@/utils/request'

// @Tags ActAct
// @Summary 创建ActAct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActAct true "创建ActAct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actAct/createActAct [post]
export const createActAct = (data) => {
  return service({
    url: '/actAct/createActAct',
    method: 'post',
    data
  })
}

// @Tags ActAct
// @Summary 删除ActAct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActAct true "删除ActAct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /actAct/deleteActAct [delete]
export const deleteActAct = (data) => {
  return service({
    url: '/actAct/deleteActAct',
    method: 'delete',
    data
  })
}

// @Tags ActAct
// @Summary 删除ActAct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ActAct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /actAct/deleteActAct [delete]
export const deleteActActByIds = (data) => {
  return service({
    url: '/actAct/deleteActActByIds',
    method: 'delete',
    data
  })
}

// @Tags ActAct
// @Summary 更新ActAct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActAct true "更新ActAct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /actAct/updateActAct [put]
export const updateActAct = (data) => {
  return service({
    url: '/actAct/updateActAct',
    method: 'put',
    data
  })
}

// @Tags ActAct
// @Summary 用id查询ActAct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ActAct true "用id查询ActAct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /actAct/findActAct [get]
export const findActAct = (params) => {
  return service({
    url: '/actAct/findActAct',
    method: 'get',
    params
  })
}

// @Tags ActAct
// @Summary 分页获取ActAct列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ActAct列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actAct/getActActList [get]
export const getActActList = (params) => {
  return service({
    url: '/actAct/getActActList',
    method: 'get',
    params
  })
}


// @Tags ActAct
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActAct true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actAct/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/actAct/quickEdit',
    method: 'post',
    data
  })
}

// @Tags ActAct
// @Summary 导出excelActAct列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelActAct列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actAct/getActActexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/actAct/excelList',
    method: 'get',
    params
  })
}