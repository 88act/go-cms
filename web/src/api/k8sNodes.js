import service from '@/utils/request'

// @Tags K8sNodes
// @Summary 创建K8sNodes
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sNodes true "创建K8sNodes"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sNodes/createK8sNodes [post]
export const createK8sNodes = (data) => {
  return service({
    url: '/k8sNodes/createK8sNodes',
    method: 'post',
    data
  })
}

// @Tags K8sNodes
// @Summary 删除K8sNodes
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sNodes true "删除K8sNodes"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sNodes/deleteK8sNodes [delete]
export const deleteK8sNodes = (data) => {
  return service({
    url: '/k8sNodes/deleteK8sNodes',
    method: 'delete',
    data
  })
}

// @Tags K8sNodes
// @Summary 删除K8sNodes
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除K8sNodes"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sNodes/deleteK8sNodes [delete]
export const deleteK8sNodesByIds = (data) => {
  return service({
    url: '/k8sNodes/deleteK8sNodesByIds',
    method: 'delete',
    data
  })
}

// @Tags K8sNodes
// @Summary 更新K8sNodes
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sNodes true "更新K8sNodes"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /k8sNodes/updateK8sNodes [put]
export const updateK8sNodes = (data) => {
  return service({
    url: '/k8sNodes/updateK8sNodes',
    method: 'put',
    data
  })
}

// @Tags K8sNodes
// @Summary 用id查询K8sNodes
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.K8sNodes true "用id查询K8sNodes"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /k8sNodes/findK8sNodes [get]
export const findK8sNodes = (params) => {
  return service({
    url: '/k8sNodes/findK8sNodes',
    method: 'get',
    params
  })
}

// @Tags K8sNodes
// @Summary 分页获取K8sNodes列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取K8sNodes列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sNodes/getK8sNodesList [get]
export const getK8sNodesList = (params) => {
  return service({
    url: '/k8sNodes/getK8sNodesList',
    method: 'get',
    params
  })
}


// @Tags K8sNodes
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sNodes true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sNodes/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/k8sNodes/quickEdit',
    method: 'post',
    data
  })
}

// @Tags K8sNodes
// @Summary 导出excelK8sNodes列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelK8sNodes列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sNodes/getK8sNodesexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/k8sNodes/excelList',
    method: 'get',
    params
  })
}