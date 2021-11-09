import service from '@/utils/request'

export const preview = (data) => {
  return service({
    url: '/superBuilder/preview',
    method: 'post',
    data
  })
}

export const createTemp = (data) => {
  return service({
    url: '/superBuilder/createTemp',
    method: 'post',
    data,
    responseType: 'blob'
  })
}

// @Tags SysApi
// @Summary 获取当前所有数据库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /superBuilder/getDatabase [get]
export const getDB = () => {
  return service({
    url: '/superBuilder/getDB',
    method: 'get'
  })
}

// @Tags SysApi
// @Summary 获取当前数据库所有表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /superBuilder/getTables [get]
export const getTable = (params) => {
  return service({
    url: '/superBuilder/getTables',
    method: 'get',
    params
  })
}

// @Tags SysApi
// @Summary 获取当前数据库所有表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /superBuilder/getColumn [get]
export const getColumn = (params) => {
  return service({
    url: '/superBuilder/getColumn',
    method: 'get',
    params
  })
}

export const getSysHistory = (data) => {
  return service({
    url: '/superBuilder/getSysHistory',
    method: 'post',
    data
  })
}

export const rollback = (data) => {
  return service({
    url: '/superBuilder/rollback',
    method: 'post',
    data
  })
}

export const getMeta = (data) => {
  return service({
    url: '/superBuilder/getMeta',
    method: 'post',
    data
  })
}

export const delSysHistory = (data) => {
  return service({
    url: '/superBuilder/delSysHistory',
    method: 'post',
    data
  })
}
