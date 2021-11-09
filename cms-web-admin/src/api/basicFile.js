import service from '@/utils/request'

// @Tags BasicFile
// @Summary 创建BasicFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BasicFile true "创建BasicFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /basicFile/createBasicFile [post]
export const createBasicFile = (data) => {
  return service({
    url: '/basicFile/createBasicFile',
    method: 'post',
    data
  })
}

// @Tags BasicFile
// @Summary 删除BasicFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BasicFile true "删除BasicFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /basicFile/deleteBasicFile [delete]
export const deleteBasicFile = (data) => {
  return service({
    url: '/basicFile/deleteBasicFile',
    method: 'delete',
    data
  })
}

// @Tags BasicFile
// @Summary 删除BasicFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BasicFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /basicFile/deleteBasicFile [delete]
export const deleteBasicFileByIds = (data) => {
  return service({
    url: '/basicFile/deleteBasicFileByIds',
    method: 'delete',
    data
  })
}

// @Tags BasicFile
// @Summary 更新BasicFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BasicFile true "更新BasicFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /basicFile/updateBasicFile [put]
export const updateBasicFile = (data) => {
  return service({
    url: '/basicFile/updateBasicFile',
    method: 'put',
    data
  })
}

// @Tags BasicFile
// @Summary 用id查询BasicFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BasicFile true "用id查询BasicFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /basicFile/findBasicFile [get]
export const findBasicFile = (params) => {
  return service({
    url: '/basicFile/findBasicFile',
    method: 'get',
    params
  })
}

// @Tags BasicFile
// @Summary 分页获取BasicFile列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BasicFile列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /basicFile/getBasicFileList [get]
export const getBasicFileList = (params) => {
  return service({
    url: '/basicFile/getBasicFileList',
    method: 'get',
    params
  })
}


// @Tags BasicFile
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BasicFile true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /basicFile/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/basicFile/quickEdit',
    method: 'post',
    data
  })
}