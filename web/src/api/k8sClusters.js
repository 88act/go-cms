import service from '@/utils/request'

// @Tags K8sClusters
// @Summary 创建K8sClusters
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sClusters true "创建K8sClusters"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sClusters/createK8sClusters [post]
export const createK8sClusters = (data) => {
  return service({
    url: '/k8sClusters/createK8sClusters',
    method: 'post',
    data
  })
}

// @Tags K8sClusters
// @Summary 删除K8sClusters
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sClusters true "删除K8sClusters"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sClusters/deleteK8sClusters [delete]
export const deleteK8sClusters = (data) => {
  return service({
    url: '/k8sClusters/deleteK8sClusters',
    method: 'delete',
    data
  })
}

// @Tags K8sClusters
// @Summary 删除K8sClusters
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除K8sClusters"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sClusters/deleteK8sClusters [delete]
export const deleteK8sClustersByIds = (data) => {
  return service({
    url: '/k8sClusters/deleteK8sClustersByIds',
    method: 'delete',
    data
  })
}

// @Tags K8sClusters
// @Summary 更新K8sClusters
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sClusters true "更新K8sClusters"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /k8sClusters/updateK8sClusters [put]
export const updateK8sClusters = (data) => {
  return service({
    url: '/k8sClusters/updateK8sClusters',
    method: 'put',
    data
  })
}

// @Tags K8sClusters
// @Summary 用id查询K8sClusters
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.K8sClusters true "用id查询K8sClusters"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /k8sClusters/findK8sClusters [get]
export const findK8sClusters = (params) => {
  return service({
    url: '/k8sClusters/findK8sClusters',
    method: 'get',
    params
  })
}

// @Tags K8sClusters
// @Summary 分页获取K8sClusters列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取K8sClusters列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sClusters/getK8sClustersList [get]
export const getK8sClustersList = (params) => {
  return service({
    url: '/k8sClusters/getK8sClustersList',
    method: 'get',
    params
  })
}


// @Tags K8sClusters
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sClusters true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sClusters/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/k8sClusters/quickEdit',
    method: 'post',
    data
  })
}

// @Tags K8sClusters
// @Summary 导出excelK8sClusters列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelK8sClusters列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sClusters/getK8sClustersexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/k8sClusters/excelList',
    method: 'get',
    params
  })
}