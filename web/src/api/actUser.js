import service from '@/utils/request'

// @Tags ActUser
// @Summary 创建ActUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActUser true "创建ActUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actUser/createActUser [post]
export const createActUser = (data) => {
  return service({
    url: '/actUser/createActUser',
    method: 'post',
    data
  })
}

// @Tags ActUser
// @Summary 删除ActUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActUser true "删除ActUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /actUser/deleteActUser [delete]
export const deleteActUser = (data) => {
  return service({
    url: '/actUser/deleteActUser',
    method: 'delete',
    data
  })
}

// @Tags ActUser
// @Summary 删除ActUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ActUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /actUser/deleteActUser [delete]
export const deleteActUserByIds = (data) => {
  return service({
    url: '/actUser/deleteActUserByIds',
    method: 'delete',
    data
  })
}

// @Tags ActUser
// @Summary 更新ActUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActUser true "更新ActUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /actUser/updateActUser [put]
export const updateActUser = (data) => {
  return service({
    url: '/actUser/updateActUser',
    method: 'put',
    data
  })
}

// @Tags ActUser
// @Summary 用id查询ActUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ActUser true "用id查询ActUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /actUser/findActUser [get]
export const findActUser = (params) => {
  return service({
    url: '/actUser/findActUser',
    method: 'get',
    params
  })
}

// @Tags ActUser
// @Summary 分页获取ActUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ActUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actUser/getActUserList [get]
export const getActUserList = (params) => {
  return service({
    url: '/actUser/getActUserList',
    method: 'get',
    params
  })
}


// @Tags ActUser
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActUser true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actUser/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/actUser/quickEdit',
    method: 'post',
    data
  })
}

// @Tags ActUser
// @Summary 导出excelActUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelActUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actUser/getActUserexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/actUser/excelList',
    method: 'get',
    params
  })
}