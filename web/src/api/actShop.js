import service from '@/utils/request'

// @Tags ActShop
// @Summary 创建ActShop
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActShop true "创建ActShop"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actShop/createActShop [post]
export const createActShop = (data) => {
  return service({
    url: '/actShop/createActShop',
    method: 'post',
    data
  })
}

// @Tags ActShop
// @Summary 删除ActShop
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActShop true "删除ActShop"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /actShop/deleteActShop [delete]
export const deleteActShop = (data) => {
  return service({
    url: '/actShop/deleteActShop',
    method: 'delete',
    data
  })
}

// @Tags ActShop
// @Summary 删除ActShop
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ActShop"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /actShop/deleteActShop [delete]
export const deleteActShopByIds = (data) => {
  return service({
    url: '/actShop/deleteActShopByIds',
    method: 'delete',
    data
  })
}

// @Tags ActShop
// @Summary 更新ActShop
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActShop true "更新ActShop"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /actShop/updateActShop [put]
export const updateActShop = (data) => {
  return service({
    url: '/actShop/updateActShop',
    method: 'put',
    data
  })
}

// @Tags ActShop
// @Summary 用id查询ActShop
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ActShop true "用id查询ActShop"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /actShop/findActShop [get]
export const findActShop = (params) => {
  return service({
    url: '/actShop/findActShop',
    method: 'get',
    params
  })
}

// @Tags ActShop
// @Summary 分页获取ActShop列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ActShop列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actShop/getActShopList [get]
export const getActShopList = (params) => {
  return service({
    url: '/actShop/getActShopList',
    method: 'get',
    params
  })
}


// @Tags ActShop
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActShop true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actShop/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/actShop/quickEdit',
    method: 'post',
    data
  })
}

// @Tags ActShop
// @Summary 导出excelActShop列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelActShop列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actShop/getActShopexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/actShop/excelList',
    method: 'get',
    params
  })
}