import service from '@/utils/request'

// @Tags BbPiInstitutionBusiness
// @Summary 创建BbPiInstitutionBusiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiInstitutionBusiness true "创建BbPiInstitutionBusiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiInstitutionBusiness/createBbPiInstitutionBusiness [post]
export const createBbPiInstitutionBusiness = (data) => {
  return service({
    url: '/bbPiInstitutionBusiness/createBbPiInstitutionBusiness',
    method: 'post',
    data
  })
}

// @Tags BbPiInstitutionBusiness
// @Summary 删除BbPiInstitutionBusiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiInstitutionBusiness true "删除BbPiInstitutionBusiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiInstitutionBusiness/deleteBbPiInstitutionBusiness [delete]
export const deleteBbPiInstitutionBusiness = (data) => {
  return service({
    url: '/bbPiInstitutionBusiness/deleteBbPiInstitutionBusiness',
    method: 'delete',
    data
  })
}

// @Tags BbPiInstitutionBusiness
// @Summary 删除BbPiInstitutionBusiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BbPiInstitutionBusiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiInstitutionBusiness/deleteBbPiInstitutionBusiness [delete]
export const deleteBbPiInstitutionBusinessByIds = (data) => {
  return service({
    url: '/bbPiInstitutionBusiness/deleteBbPiInstitutionBusinessByIds',
    method: 'delete',
    data
  })
}

// @Tags BbPiInstitutionBusiness
// @Summary 更新BbPiInstitutionBusiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiInstitutionBusiness true "更新BbPiInstitutionBusiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bbPiInstitutionBusiness/updateBbPiInstitutionBusiness [put]
export const updateBbPiInstitutionBusiness = (data) => {
  return service({
    url: '/bbPiInstitutionBusiness/updateBbPiInstitutionBusiness',
    method: 'put',
    data
  })
}

// @Tags BbPiInstitutionBusiness
// @Summary 用id查询BbPiInstitutionBusiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BbPiInstitutionBusiness true "用id查询BbPiInstitutionBusiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bbPiInstitutionBusiness/findBbPiInstitutionBusiness [get]
export const findBbPiInstitutionBusiness = (params) => {
  return service({
    url: '/bbPiInstitutionBusiness/findBbPiInstitutionBusiness',
    method: 'get',
    params
  })
}

// @Tags BbPiInstitutionBusiness
// @Summary 分页获取BbPiInstitutionBusiness列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BbPiInstitutionBusiness列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiInstitutionBusiness/getBbPiInstitutionBusinessList [get]
export const getBbPiInstitutionBusinessList = (params) => {
  return service({
    url: '/bbPiInstitutionBusiness/getBbPiInstitutionBusinessList',
    method: 'get',
    params
  })
}


// @Tags BbPiInstitutionBusiness
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiInstitutionBusiness true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiInstitutionBusiness/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/bbPiInstitutionBusiness/quickEdit',
    method: 'post',
    data
  })
}

// @Tags BbPiInstitutionBusiness
// @Summary 导出excelBbPiInstitutionBusiness列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelBbPiInstitutionBusiness列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiInstitutionBusiness/getBbPiInstitutionBusinessexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/bbPiInstitutionBusiness/excelList',
    method: 'get',
    params
  })
}