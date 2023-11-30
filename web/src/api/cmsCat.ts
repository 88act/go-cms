import { http,resData } from "@/utils/http";
 
// @Tags CmsCat
// @Summary 创建CmsCat
export const createCmsCat = (data?: object) => {
   return http.request<resData>("post", "/cmsCat/createCmsCat", { data });  
}

// @Tags CmsCat
// @Summary 删除CmsCat
export const deleteCmsCat = (data?: object) => {
   return http.request<resData>("post", "/cmsCat/deleteCmsCat", { data });  
}
 

// @Tags CmsCat
// @Summary 删除ByIdsCmsCat
export const deleteCmsCatByIds = (data?: object) => {
   return http.request<resData>("post", "/cmsCat/deleteCmsCatByIds", { data });  
}
 

// @Tags CmsCat
// @Summary 更新CmsCat
export const updateCmsCat = (data?: object) => {
   return http.request<resData>("post", "/cmsCat/updateCmsCat", { data });  
}

// @Tags CmsCat
// @Summary 快速编辑
export const quickEdit = (data?: object) => {
   return http.request<resData>("post", "/cmsCat/quickEdit", { data });  
}
 

// @Tags CmsCat
// @Summary 用id查询CmsCat
export const findCmsCat = (params?: object) => {
   return http.request<resData>("get", "/cmsCat/findCmsCat", { params });  
}
 

// @Tags CmsCat
// @Summary 分页获取CmsCat列表
export const getCmsCatList = (params?: object) => {
   return http.request<resData>("get", "/cmsCat/getCmsCatList", { params });  
}
 

// @Tags CmsCat
// @Summary 导出excelCmsCat列表
export const excelList = (params?: object) => {
   return http.request<resData>("get", "/cmsCat/excelList", { params });  
}
