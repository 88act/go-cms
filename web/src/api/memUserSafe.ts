import { http,resData } from "@/utils/http";
 
// @Tags MemUserSafe
// @Summary 创建MemUserSafe
export const createMemUserSafe = (data?: object) => {
   return http.request<resData>("post", "/memUserSafe/createMemUserSafe", { data });  
}

// @Tags MemUserSafe
// @Summary 删除MemUserSafe
export const deleteMemUserSafe = (data?: object) => {
   return http.request<resData>("post", "/memUserSafe/deleteMemUserSafe", { data });  
}
 

// @Tags MemUserSafe
// @Summary 删除ByIdsMemUserSafe
export const deleteMemUserSafeByIds = (data?: object) => {
   return http.request<resData>("post", "/memUserSafe/deleteMemUserSafeByIds", { data });  
}
 

// @Tags MemUserSafe
// @Summary 更新MemUserSafe
export const updateMemUserSafe = (data?: object) => {
   return http.request<resData>("post", "/memUserSafe/updateMemUserSafe", { data });  
}

// @Tags MemUserSafe
// @Summary 快速编辑
export const quickEdit = (data?: object) => {
   return http.request<resData>("post", "/memUserSafe/quickEdit", { data });  
}
 

// @Tags MemUserSafe
// @Summary 用id查询MemUserSafe
export const findMemUserSafe = (params?: object) => {
   return http.request<resData>("get", "/memUserSafe/findMemUserSafe", { params });  
}
 

// @Tags MemUserSafe
// @Summary 分页获取MemUserSafe列表
export const getMemUserSafeList = (params?: object) => {
   return http.request<resData>("get", "/memUserSafe/getMemUserSafeList", { params });  
}
 

// @Tags MemUserSafe
// @Summary 导出excelMemUserSafe列表
export const excelList = (params?: object) => {
   return http.request<resData>("get", "/memUserSafe/excelList", { params });  
}
