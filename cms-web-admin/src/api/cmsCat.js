import service from '@/utils/request'

// @Tags CmsCat
// @Summary 创建CmsCat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsCat true "创建CmsCat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsCat/createCmsCat [post]
export const createCmsCat = (data) => {
  return service({
    url: '/cmsCat/createCmsCat',
    method: 'post',
    data
  })
}

// @Tags CmsCat
// @Summary 删除CmsCat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsCat true "删除CmsCat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsCat/deleteCmsCat [delete]
export const deleteCmsCat = (data) => {
  return service({
    url: '/cmsCat/deleteCmsCat',
    method: 'delete',
    data
  })
}

// @Tags CmsCat
// @Summary 删除CmsCat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CmsCat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsCat/deleteCmsCat [delete]
export const deleteCmsCatByIds = (data) => {
  return service({
    url: '/cmsCat/deleteCmsCatByIds',
    method: 'delete',
    data
  })
}

// @Tags CmsCat
// @Summary 更新CmsCat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsCat true "更新CmsCat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsCat/updateCmsCat [put]
export const updateCmsCat = (data) => {
  return service({
    url: '/cmsCat/updateCmsCat',
    method: 'put',
    data
  })
}

// @Tags CmsCat
// @Summary 用id查询CmsCat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CmsCat true "用id查询CmsCat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsCat/findCmsCat [get]
export const findCmsCat = (params) => {
  return service({
    url: '/cmsCat/findCmsCat',
    method: 'get',
    params
  })
}

// @Tags CmsCat
// @Summary 分页获取CmsCat列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CmsCat列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsCat/getCmsCatList [get]
export const getCmsCatList = (params) => {
  return service({
    url: '/cmsCat/getCmsCatList',
    method: 'get',
    params
  })
}


// @Tags CmsCat
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmsCat true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsCat/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/cmsCat/quickEdit',
    method: 'post',
    data
  })
}