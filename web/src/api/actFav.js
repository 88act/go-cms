import service from '@/utils/request'

// @Tags ActFav
// @Summary 创建ActFav
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActFav true "创建ActFav"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actFav/createActFav [post]
export const createActFav = (data) => {
  return service({
    url: '/actFav/createActFav',
    method: 'post',
    data
  })
}

// @Tags ActFav
// @Summary 删除ActFav
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActFav true "删除ActFav"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /actFav/deleteActFav [delete]
export const deleteActFav = (data) => {
  return service({
    url: '/actFav/deleteActFav',
    method: 'delete',
    data
  })
}

// @Tags ActFav
// @Summary 删除ActFav
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ActFav"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /actFav/deleteActFav [delete]
export const deleteActFavByIds = (data) => {
  return service({
    url: '/actFav/deleteActFavByIds',
    method: 'delete',
    data
  })
}

// @Tags ActFav
// @Summary 更新ActFav
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActFav true "更新ActFav"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /actFav/updateActFav [put]
export const updateActFav = (data) => {
  return service({
    url: '/actFav/updateActFav',
    method: 'put',
    data
  })
}

// @Tags ActFav
// @Summary 用id查询ActFav
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ActFav true "用id查询ActFav"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /actFav/findActFav [get]
export const findActFav = (params) => {
  return service({
    url: '/actFav/findActFav',
    method: 'get',
    params
  })
}

// @Tags ActFav
// @Summary 分页获取ActFav列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ActFav列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actFav/getActFavList [get]
export const getActFavList = (params) => {
  return service({
    url: '/actFav/getActFavList',
    method: 'get',
    params
  })
}


// @Tags ActFav
// @Summary 快速编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActFav true "快速编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actFav/quickEdit [post]
export const quickEdit = (data) => {
  return service({
    url: '/actFav/quickEdit',
    method: 'post',
    data
  })
}

// @Tags ActFav
// @Summary 导出excelActFav列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页导出excelActFav列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /actFav/getActFavexcelList [get]
export const excelList = (params) => {
  return service({
    url: '/actFav/excelList',
    method: 'get',
    params
  })
}