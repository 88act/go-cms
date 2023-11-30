import { http,resData } from "@/utils/http";
 
// @Tags BasicFileCat
// @Summary 创建BasicFileCat
export const createBasicFileCat = (data?: object) => {
   return http.request<resData>("post", "/basicFileCat/createBasicFileCat", { data });  
}

// @Tags BasicFileCat
// @Summary 删除BasicFileCat
export const deleteBasicFileCat = (data?: object) => {
   return http.request<resData>("post", "/basicFileCat/deleteBasicFileCat", { data });  
}
 

// @Tags BasicFileCat
// @Summary 删除ByIdsBasicFileCat
export const deleteBasicFileCatByIds = (data?: object) => {
   return http.request<resData>("post", "/basicFileCat/deleteBasicFileCatByIds", { data });  
}
 

// @Tags BasicFileCat
// @Summary 更新BasicFileCat
export const updateBasicFileCat = (data?: object) => {
   return http.request<resData>("post", "/basicFileCat/updateBasicFileCat", { data });  
}

// @Tags BasicFileCat
// @Summary 快速编辑
export const quickEdit = (data?: object) => {
   return http.request<resData>("post", "/basicFileCat/quickEdit", { data });  
}
 

// @Tags BasicFileCat
// @Summary 用id查询BasicFileCat
export const findBasicFileCat = (params?: object) => {
   return http.request<resData>("get", "/basicFileCat/findBasicFileCat", { params });  
}
 

// @Tags BasicFileCat
// @Summary 分页获取BasicFileCat列表
export const getBasicFileCatList = (params?: object) => {
   return http.request<resData>("get", "/basicFileCat/getBasicFileCatList", { params });  
}
 

// @Tags BasicFileCat
// @Summary 导出excelBasicFileCat列表
export const excelList = (params?: object) => {
   return http.request<resData>("get", "/basicFileCat/excelList", { params });  
}
