import service from '@/utils/request'

// @Tags ImTxMsgFile
// @Summary 创建ImTxMsgFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ImTxMsgFile true "创建ImTxMsgFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxMsgFile/createImTxMsgFile [post]
export const createImTxMsgFile = (data) => {
  return service({
    url: '/imTxMsgFile/createImTxMsgFile',
    method: 'post',
    data
  })
}

// @Tags ImTxMsgFile
// @Summary 删除ImTxMsgFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ImTxMsgFile true "删除ImTxMsgFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /imTxMsgFile/deleteImTxMsgFile [delete]
export const deleteImTxMsgFile = (data) => {
  return service({
    url: '/imTxMsgFile/deleteImTxMsgFile',
    method: 'delete',
    data
  })
}

// @Tags ImTxMsgFile
// @Summary 删除ImTxMsgFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ImTxMsgFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /imTxMsgFile/deleteImTxMsgFile [delete]
export const deleteImTxMsgFileByIds = (data) => {
  return service({
    url: '/imTxMsgFile/deleteImTxMsgFileByIds',
    method: 'delete',
    data
  })
}

// @Tags ImTxMsgFile
// @Summary 更新ImTxMsgFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ImTxMsgFile true "更新ImTxMsgFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /imTxMsgFile/updateImTxMsgFile [put]
export const updateImTxMsgFile = (data) => {
  return service({
    url: '/imTxMsgFile/updateImTxMsgFile',
    method: 'put',
    data
  })
}

// @Tags ImTxMsgFile
// @Summary 用id查询ImTxMsgFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ImTxMsgFile true "用id查询ImTxMsgFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /imTxMsgFile/findImTxMsgFile [get]
export const findImTxMsgFile = (params) => {
  return service({
    url: '/imTxMsgFile/findImTxMsgFile',
    method: 'get',
    params
  })
}

// @Tags ImTxMsgFile
// @Summary 分页获取ImTxMsgFile列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ImTxMsgFile列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxMsgFile/getImTxMsgFileList [get]
export const getImTxMsgFileList = (params) => {
  return service({
    url: '/imTxMsgFile/getImTxMsgFileList',
    method: 'get',
    params
  })
}


// @Tags ImTxMsgFile
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ImTxMsgFile true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxMsgFile/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/imTxMsgFile/quickEdit',
    method: 'post',
    data
  })
}

// @Tags ImTxMsgFile
// @Summary 导出excelImTxMsgFile列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelImTxMsgFile列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxMsgFile/getImTxMsgFileexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/imTxMsgFile/excelList',
    method: 'get',
    params
  })
}