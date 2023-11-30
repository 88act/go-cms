import { http,resData } from "@/utils/http";
 
// @Tags MemLogs
// @Summary 创建MemLogs
export const createMemLogs = (data?: object) => {
   return http.request<resData>("post", "/memLogs/createMemLogs", { data });  
}

// @Tags MemLogs
// @Summary 删除MemLogs
export const deleteMemLogs = (data?: object) => {
   return http.request<resData>("post", "/memLogs/deleteMemLogs", { data });  
}
 

// @Tags MemLogs
// @Summary 删除ByIdsMemLogs
export const deleteMemLogsByIds = (data?: object) => {
   return http.request<resData>("post", "/memLogs/deleteMemLogsByIds", { data });  
}
 

// @Tags MemLogs
// @Summary 更新MemLogs
export const updateMemLogs = (data?: object) => {
   return http.request<resData>("post", "/memLogs/updateMemLogs", { data });  
}

// @Tags MemLogs
// @Summary 快速编辑
export const quickEdit = (data?: object) => {
   return http.request<resData>("post", "/memLogs/quickEdit", { data });  
}
 

// @Tags MemLogs
// @Summary 用id查询MemLogs
export const findMemLogs = (params?: object) => {
   return http.request<resData>("get", "/memLogs/findMemLogs", { params });  
}
 

// @Tags MemLogs
// @Summary 分页获取MemLogs列表
export const getMemLogsList = (params?: object) => {
   return http.request<resData>("get", "/memLogs/getMemLogsList", { params });  
}
 

// @Tags MemLogs
// @Summary 导出excelMemLogs列表
export const excelList = (params?: object) => {
   return http.request<resData>("get", "/memLogs/excelList", { params });  
}
