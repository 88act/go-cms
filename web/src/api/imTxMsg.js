import service from '@/utils/request'

// @Tags ImTxMsg
// @Summary 创建ImTxMsg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ImTxMsg true "创建ImTxMsg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxMsg/createImTxMsg [post]
export const createImTxMsg = (data) => {
  return service({
    url: '/imTxMsg/createImTxMsg',
    method: 'post',
    data
  })
}

// @Tags ImTxMsg
// @Summary 删除ImTxMsg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ImTxMsg true "删除ImTxMsg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /imTxMsg/deleteImTxMsg [delete]
export const deleteImTxMsg = (data) => {
  return service({
    url: '/imTxMsg/deleteImTxMsg',
    method: 'delete',
    data
  })
}

// @Tags ImTxMsg
// @Summary 删除ImTxMsg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ImTxMsg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /imTxMsg/deleteImTxMsg [delete]
export const deleteImTxMsgByIds = (data) => {
  return service({
    url: '/imTxMsg/deleteImTxMsgByIds',
    method: 'delete',
    data
  })
}

// @Tags ImTxMsg
// @Summary 更新ImTxMsg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ImTxMsg true "更新ImTxMsg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /imTxMsg/updateImTxMsg [put]
export const updateImTxMsg = (data) => {
  return service({
    url: '/imTxMsg/updateImTxMsg',
    method: 'put',
    data
  })
}

// @Tags ImTxMsg
// @Summary 用id查询ImTxMsg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ImTxMsg true "用id查询ImTxMsg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /imTxMsg/findImTxMsg [get]
export const findImTxMsg = (params) => {
  return service({
    url: '/imTxMsg/findImTxMsg',
    method: 'get',
    params
  })
}

// @Tags ImTxMsg
// @Summary 分页获取ImTxMsg列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ImTxMsg列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxMsg/getImTxMsgList [get]
export const getImTxMsgList = (params) => {
  return service({
    url: '/imTxMsg/getImTxMsgList',
    method: 'get',
    params
  })
}


// @Tags ImTxMsg
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ImTxMsg true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxMsg/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/imTxMsg/quickEdit',
    method: 'post',
    data
  })
}

// @Tags ImTxMsg
// @Summary 导出excelImTxMsg列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelImTxMsg列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imTxMsg/getImTxMsgexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/imTxMsg/excelList',
    method: 'get',
    params
  })
}