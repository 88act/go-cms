import service from '@/utils/request'

// @Tags BbPiPerson
// @Summary 创建BbPiPerson
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiPerson true "创建BbPiPerson"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiPerson/createBbPiPerson [post]
export const createBbPiPerson = (data) => {
  return service({
    url: '/bbPiPerson/createBbPiPerson',
    method: 'post',
    data
  })
}

// @Tags BbPiPerson
// @Summary 删除BbPiPerson
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiPerson true "删除BbPiPerson"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiPerson/deleteBbPiPerson [delete]
export const deleteBbPiPerson = (data) => {
  return service({
    url: '/bbPiPerson/deleteBbPiPerson',
    method: 'delete',
    data
  })
}

// @Tags BbPiPerson
// @Summary 删除BbPiPerson
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BbPiPerson"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiPerson/deleteBbPiPerson [delete]
export const deleteBbPiPersonByIds = (data) => {
  return service({
    url: '/bbPiPerson/deleteBbPiPersonByIds',
    method: 'delete',
    data
  })
}

// @Tags BbPiPerson
// @Summary 更新BbPiPerson
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiPerson true "更新BbPiPerson"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bbPiPerson/updateBbPiPerson [put]
export const updateBbPiPerson = (data) => {
  return service({
    url: '/bbPiPerson/updateBbPiPerson',
    method: 'put',
    data
  })
}

// @Tags BbPiPerson
// @Summary 用id查询BbPiPerson
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BbPiPerson true "用id查询BbPiPerson"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bbPiPerson/findBbPiPerson [get]
export const findBbPiPerson = (params) => {
  return service({
    url: '/bbPiPerson/findBbPiPerson',
    method: 'get',
    params
  })
}

// @Tags BbPiPerson
// @Summary 分页获取BbPiPerson列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BbPiPerson列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiPerson/getBbPiPersonList [get]
export const getBbPiPersonList = (params) => {
  return service({
    url: '/bbPiPerson/getBbPiPersonList',
    method: 'get',
    params
  })
}


// @Tags BbPiPerson
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BbPiPerson true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiPerson/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/bbPiPerson/quickEdit',
    method: 'post',
    data
  })
}

// @Tags BbPiPerson
// @Summary 导出excelBbPiPerson列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelBbPiPerson列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiPerson/getBbPiPersonexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/bbPiPerson/excelList',
    method: 'get',
    params
  })
}