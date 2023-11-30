import { http,resData } from "@/utils/http";
 
// @Tags CmsCatArt
// @Summary 创建CmsCatArt
export const createCmsCatArt = (data?: object) => {
   return http.request<resData>("post", "/cmsCatArt/createCmsCatArt", { data });  
}

// @Tags CmsCatArt
// @Summary 删除CmsCatArt
export const deleteCmsCatArt = (data?: object) => {
   return http.request<resData>("post", "/cmsCatArt/deleteCmsCatArt", { data });  
}
 

// @Tags CmsCatArt
// @Summary 删除ByIdsCmsCatArt
export const deleteCmsCatArtByIds = (data?: object) => {
   return http.request<resData>("post", "/cmsCatArt/deleteCmsCatArtByIds", { data });  
}
 

// @Tags CmsCatArt
// @Summary 更新CmsCatArt
export const updateCmsCatArt = (data?: object) => {
   return http.request<resData>("post", "/cmsCatArt/updateCmsCatArt", { data });  
}

// @Tags CmsCatArt
// @Summary 快速编辑
export const quickEdit = (data?: object) => {
   return http.request<resData>("post", "/cmsCatArt/quickEdit", { data });  
}
 

// @Tags CmsCatArt
// @Summary 用id查询CmsCatArt
export const findCmsCatArt = (params?: object) => {
   return http.request<resData>("get", "/cmsCatArt/findCmsCatArt", { params });  
}
 

// @Tags CmsCatArt
// @Summary 分页获取CmsCatArt列表
export const getCmsCatArtList = (params?: object) => {
   return http.request<resData>("get", "/cmsCatArt/getCmsCatArtList", { params });  
}
 

// @Tags CmsCatArt
// @Summary 导出excelCmsCatArt列表
export const excelList = (params?: object) => {
   return http.request<resData>("get", "/cmsCatArt/excelList", { params });  
}
