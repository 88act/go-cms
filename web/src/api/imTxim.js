import service from '@/utils/request'

// @Tags ImTxim
// @Summary 创建ImTxim
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ImTxim true "创建ImTxim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxim/createImTxim [post]
export const createImTxim = (data) => {
  return service({
    url: '/imTxim/createImTxim',
    method: 'post',
    data
  })
}

// @Tags ImTxim
// @Summary 删除ImTxim
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ImTxim true "删除ImTxim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /imTxim/deleteImTxim [delete]
export const deleteImTxim = (data) => {
  return service({
    url: '/imTxim/deleteImTxim',
    method: 'delete',
    data
  })
}

// @Tags ImTxim
// @Summary 删除ImTxim
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ImTxim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /imTxim/deleteImTxim [delete]
export const deleteImTximByIds = (data) => {
  return service({
    url: '/imTxim/deleteImTximByIds',
    method: 'delete',
    data
  })
}

// @Tags ImTxim
// @Summary 更新ImTxim
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ImTxim true "更新ImTxim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /imTxim/updateImTxim [put]
export const updateImTxim = (data) => {
  return service({
    url: '/imTxim/updateImTxim',
    method: 'put',
    data
  })
}

// @Tags ImTxim
// @Summary 用id查询ImTxim
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ImTxim true "用id查询ImTxim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /imTxim/findImTxim [get]
export const findImTxim = (params) => {
  return service({
    url: '/imTxim/findImTxim',
    method: 'get',
    params
  })
}

// @Tags ImTxim
// @Summary 分页获取ImTxim列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ImTxim列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxim/getImTximList [get]
export const getImTximList = (params) => {
  return service({
    url: '/imTxim/getImTximList',
    method: 'get',
    params
  })
}


// @Tags ImTxim
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ImTxim true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxim/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/imTxim/quickEdit',
    method: 'post',
    data
  })
}

// @Tags ImTxim
// @Summary 导出excelImTxim列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelImTxim列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxim/getImTximexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/imTxim/excelList',
    method: 'get',
    params
  })
}


// @Tags ColCollect
// @Summary 启动或停止采集器
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ColCollect true "启动或停止采集器"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /imTxim/startOrStopCollect [get]
export const startOrStopCollect = (params) => {
  return service({
    url: '/imTxim/startOrStopCollect',
    method: 'get',
    params
  })
}