import { http,resData } from "@/utils/http";
 
// @Tags MemDepart
// @Summary 创建MemDepart
export const createMemDepart = (data?: object) => {
   return http.request<resData>("post", "/memDepart/createMemDepart", { data });  
}

// @Tags MemDepart
// @Summary 删除MemDepart
export const deleteMemDepart = (data?: object) => {
   return http.request<resData>("post", "/memDepart/deleteMemDepart", { data });  
}
 

// @Tags MemDepart
// @Summary 删除ByIdsMemDepart
export const deleteMemDepartByIds = (data?: object) => {
   return http.request<resData>("post", "/memDepart/deleteMemDepartByIds", { data });  
}
 

// @Tags MemDepart
// @Summary 更新MemDepart
export const updateMemDepart = (data?: object) => {
   return http.request<resData>("post", "/memDepart/updateMemDepart", { data });  
}

// @Tags MemDepart
// @Summary 快速编辑
export const quickEdit = (data?: object) => {
   return http.request<resData>("post", "/memDepart/quickEdit", { data });  
}
 

// @Tags MemDepart
// @Summary 用id查询MemDepart
export const findMemDepart = (params?: object) => {
   return http.request<resData>("get", "/memDepart/findMemDepart", { params });  
}
 

// @Tags MemDepart
// @Summary 分页获取MemDepart列表
export const getMemDepartList = (params?: object) => {
   return http.request<resData>("get", "/memDepart/getMemDepartList", { params });  
}
 

// @Tags MemDepart
// @Summary 导出excelMemDepart列表
export const excelList = (params?: object) => {
   return http.request<resData>("get", "/memDepart/excelList", { params });  
}
