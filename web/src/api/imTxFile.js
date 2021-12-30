import service from '@/utils/request'

// @Tags ImTxFile
// @Summary 创建ImTxFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ImTxFile true "创建ImTxFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxFile/createImTxFile [post]
export const createImTxFile = (data) => {
  return service({
    url: '/imTxFile/createImTxFile',
    method: 'post',
    data
  })
}

// @Tags ImTxFile
// @Summary 删除ImTxFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ImTxFile true "删除ImTxFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /imTxFile/deleteImTxFile [delete]
export const deleteImTxFile = (data) => {
  return service({
    url: '/imTxFile/deleteImTxFile',
    method: 'delete',
    data
  })
}

// @Tags ImTxFile
// @Summary 删除ImTxFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ImTxFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /imTxFile/deleteImTxFile [delete]
export const deleteImTxFileByIds = (data) => {
  return service({
    url: '/imTxFile/deleteImTxFileByIds',
    method: 'delete',
    data
  })
}

// @Tags ImTxFile
// @Summary 更新ImTxFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ImTxFile true "更新ImTxFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /imTxFile/updateImTxFile [put]
export const updateImTxFile = (data) => {
  return service({
    url: '/imTxFile/updateImTxFile',
    method: 'put',
    data
  })
}

// @Tags ImTxFile
// @Summary 用id查询ImTxFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ImTxFile true "用id查询ImTxFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /imTxFile/findImTxFile [get]
export const findImTxFile = (params) => {
  return service({
    url: '/imTxFile/findImTxFile',
    method: 'get',
    params
  })
}

// @Tags ImTxFile
// @Summary 分页获取ImTxFile列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ImTxFile列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxFile/getImTxFileList [get]
export const getImTxFileList = (params) => {
  return service({
    url: '/imTxFile/getImTxFileList',
    method: 'get',
    params
  })
}


// @Tags ImTxFile
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ImTxFile true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxFile/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/imTxFile/quickEdit',
    method: 'post',
    data
  })
}

// @Tags ImTxFile
// @Summary 导出excelImTxFile列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelImTxFile列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxFile/getImTxFileexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/imTxFile/excelList',
    method: 'get',
    params
  })
}