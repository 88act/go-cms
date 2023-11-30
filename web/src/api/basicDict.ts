import { http,resData } from "@/utils/http";
 
// @Tags BasicDict
// @Summary 创建BasicDict
export const createBasicDict = (data?: object) => {
   return http.request<resData>("post", "/basicDict/createBasicDict", { data });  
}

// @Tags BasicDict
// @Summary 删除BasicDict
export const deleteBasicDict = (data?: object) => {
   return http.request<resData>("post", "/basicDict/deleteBasicDict", { data });  
}
 

// @Tags BasicDict
// @Summary 删除ByIdsBasicDict
export const deleteBasicDictByIds = (data?: object) => {
   return http.request<resData>("post", "/basicDict/deleteBasicDictByIds", { data });  
}
 

// @Tags BasicDict
// @Summary 更新BasicDict
export const updateBasicDict = (data?: object) => {
   return http.request<resData>("post", "/basicDict/updateBasicDict", { data });  
}

// @Tags BasicDict
// @Summary 快速编辑
export const quickEdit = (data?: object) => {
   return http.request<resData>("post", "/basicDict/quickEdit", { data });  
}
 

// @Tags BasicDict
// @Summary 用id查询BasicDict
export const findBasicDict = (params?: object) => {
   return http.request<resData>("get", "/basicDict/findBasicDict", { params });  
}
 

// @Tags BasicDict
// @Summary 分页获取BasicDict列表
export const getBasicDictList = (params?: object) => {
   return http.request<resData>("get", "/basicDict/getBasicDictList", { params });  
}
 

// @Tags BasicDict
// @Summary 导出excelBasicDict列表
export const excelList = (params?: object) => {
   return http.request<resData>("get", "/basicDict/excelList", { params });  
}
