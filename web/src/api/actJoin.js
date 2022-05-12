import service from '@/utils/request'

// @Tags ActJoin
// @Summary 创建ActJoin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActJoin true "创建ActJoin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actJoin/createActJoin [post]
export const createActJoin = (data) => {
  return service({
    url: '/actJoin/createActJoin',
    method: 'post',
    data
  })
}

// @Tags ActJoin
// @Summary 删除ActJoin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActJoin true "删除ActJoin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /actJoin/deleteActJoin [delete]
export const deleteActJoin = (data) => {
  return service({
    url: '/actJoin/deleteActJoin',
    method: 'delete',
    data
  })
}

// @Tags ActJoin
// @Summary 删除ActJoin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ActJoin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /actJoin/deleteActJoin [delete]
export const deleteActJoinByIds = (data) => {
  return service({
    url: '/actJoin/deleteActJoinByIds',
    method: 'delete',
    data
  })
}

// @Tags ActJoin
// @Summary 更新ActJoin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActJoin true "更新ActJoin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /actJoin/updateActJoin [put]
export const updateActJoin = (data) => {
  return service({
    url: '/actJoin/updateActJoin',
    method: 'put',
    data
  })
}

// @Tags ActJoin
// @Summary 用id查询ActJoin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ActJoin true "用id查询ActJoin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /actJoin/findActJoin [get]
export const findActJoin = (params) => {
  return service({
    url: '/actJoin/findActJoin',
    method: 'get',
    params
  })
}

// @Tags ActJoin
// @Summary 分页获取ActJoin列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ActJoin列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actJoin/getActJoinList [get]
export const getActJoinList = (params) => {
  return service({
    url: '/actJoin/getActJoinList',
    method: 'get',
    params
  })
}


// @Tags ActJoin
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActJoin true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actJoin/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/actJoin/quickEdit',
    method: 'post',
    data
  })
}

// @Tags ActJoin
// @Summary 导出excelActJoin列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelActJoin列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actJoin/getActJoinexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/actJoin/excelList',
    method: 'get',
    params
  })
}