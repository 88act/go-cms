import { http,resData } from "@/utils/http";
 
// @Tags CmsBlog
// @Summary 创建CmsBlog
export const createCmsBlog = (data?: object) => {
   return http.request<resData>("post", "/cmsBlog/createCmsBlog", { data });  
}

// @Tags CmsBlog
// @Summary 删除CmsBlog
export const deleteCmsBlog = (data?: object) => {
   return http.request<resData>("post", "/cmsBlog/deleteCmsBlog", { data });  
}
 

// @Tags CmsBlog
// @Summary 删除ByIdsCmsBlog
export const deleteCmsBlogByIds = (data?: object) => {
   return http.request<resData>("post", "/cmsBlog/deleteCmsBlogByIds", { data });  
}
 

// @Tags CmsBlog
// @Summary 更新CmsBlog
export const updateCmsBlog = (data?: object) => {
   return http.request<resData>("post", "/cmsBlog/updateCmsBlog", { data });  
}

// @Tags CmsBlog
// @Summary 快速编辑
export const quickEdit = (data?: object) => {
   return http.request<resData>("post", "/cmsBlog/quickEdit", { data });  
}
 

// @Tags CmsBlog
// @Summary 用id查询CmsBlog
export const findCmsBlog = (params?: object) => {
   return http.request<resData>("get", "/cmsBlog/findCmsBlog", { params });  
}
 

// @Tags CmsBlog
// @Summary 分页获取CmsBlog列表
export const getCmsBlogList = (params?: object) => {
   return http.request<resData>("get", "/cmsBlog/getCmsBlogList", { params });  
}
 

// @Tags CmsBlog
// @Summary 导出excelCmsBlog列表
export const excelList = (params?: object) => {
   return http.request<resData>("get", "/cmsBlog/excelList", { params });  
}
