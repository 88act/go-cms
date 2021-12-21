import service from '@/utils/request'

// @Tags K8sDeployments
// @Summary 创建K8sDeployments
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sDeployments true "创建K8sDeployments"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sDeployments/createK8sDeployments [post]
export const createK8sDeployments = (data) => {
  return service({
    url: '/k8sDeployments/createK8sDeployments',
    method: 'post',
    data
  })
}

// @Tags K8sDeployments
// @Summary 删除K8sDeployments
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sDeployments true "删除K8sDeployments"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sDeployments/deleteK8sDeployments [delete]
export const deleteK8sDeployments = (data) => {
  return service({
    url: '/k8sDeployments/deleteK8sDeployments',
    method: 'delete',
    data
  })
}

// @Tags K8sDeployments
// @Summary 删除K8sDeployments
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除K8sDeployments"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sDeployments/deleteK8sDeployments [delete]
export const deleteK8sDeploymentsByIds = (data) => {
  return service({
    url: '/k8sDeployments/deleteK8sDeploymentsByIds',
    method: 'delete',
    data
  })
}

// @Tags K8sDeployments
// @Summary 更新K8sDeployments
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sDeployments true "更新K8sDeployments"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /k8sDeployments/updateK8sDeployments [put]
export const updateK8sDeployments = (data) => {
  return service({
    url: '/k8sDeployments/updateK8sDeployments',
    method: 'put',
    data
  })
}

// @Tags K8sDeployments
// @Summary 用id查询K8sDeployments
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.K8sDeployments true "用id查询K8sDeployments"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /k8sDeployments/findK8sDeployments [get]
export const findK8sDeployments = (params) => {
  return service({
    url: '/k8sDeployments/findK8sDeployments',
    method: 'get',
    params
  })
}

// @Tags K8sDeployments
// @Summary 分页获取K8sDeployments列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取K8sDeployments列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sDeployments/getK8sDeploymentsList [get]
export const getK8sDeploymentsList = (params) => {
  return service({
    url: '/k8sDeployments/getK8sDeploymentsList',
    method: 'get',
    params
  })
}


// @Tags K8sDeployments
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sDeployments true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sDeployments/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/k8sDeployments/quickEdit',
    method: 'post',
    data
  })
}

// @Tags K8sDeployments
// @Summary 导出excelK8sDeployments列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelK8sDeployments列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sDeployments/getK8sDeploymentsexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/k8sDeployments/excelList',
    method: 'get',
    params
  })
}