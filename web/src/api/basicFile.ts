import { http,resData } from "@/utils/http";
 
// @Tags BasicFile
// @Summary 创建BasicFile
export const createBasicFile = (data?: object) => {
   return http.request<resData>("post", "/basicFile/createBasicFile", { data });  
}

// @Tags BasicFile
// @Summary 删除BasicFile
export const deleteBasicFile = (data?: object) => {
   return http.request<resData>("post", "/basicFile/deleteBasicFile", { data });  
}
 

// @Tags BasicFile
// @Summary 删除ByIdsBasicFile
export const deleteBasicFileByIds = (data?: object) => {
   return http.request<resData>("post", "/basicFile/deleteBasicFileByIds", { data });  
}
 

// @Tags BasicFile
// @Summary 更新BasicFile
export const updateBasicFile = (data?: object) => {
   return http.request<resData>("post", "/basicFile/updateBasicFile", { data });  
}

// @Tags BasicFile
// @Summary 快速编辑
export const quickEdit = (data?: object) => {
   return http.request<resData>("post", "/basicFile/quickEdit", { data });  
}
 

// @Tags BasicFile
// @Summary 用id查询BasicFile
export const findBasicFile = (params?: object) => {
   return http.request<resData>("get", "/basicFile/findBasicFile", { params });  
}
 

// @Tags BasicFile
// @Summary 分页获取BasicFile列表
export const getBasicFileList = (params?: object) => {
   return http.request<resData>("get", "/basicFile/getBasicFileList", { params });  
}
 

// @Tags BasicFile
// @Summary 导出excelBasicFile列表
export const excelList = (params?: object) => {
   return http.request<resData>("get", "/basicFile/excelList", { params });  
}
