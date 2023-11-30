import { http,resData } from "@/utils/http";
 
// @Tags CmsDetail
// @Summary 创建CmsDetail
export const createCmsDetail = (data?: object) => {
   return http.request<resData>("post", "/cmsDetail/createCmsDetail", { data });  
}

// @Tags CmsDetail
// @Summary 删除CmsDetail
export const deleteCmsDetail = (data?: object) => {
   return http.request<resData>("post", "/cmsDetail/deleteCmsDetail", { data });  
}
 

// @Tags CmsDetail
// @Summary 删除ByIdsCmsDetail
export const deleteCmsDetailByIds = (data?: object) => {
   return http.request<resData>("post", "/cmsDetail/deleteCmsDetailByIds", { data });  
}
 

// @Tags CmsDetail
// @Summary 更新CmsDetail
export const updateCmsDetail = (data?: object) => {
   return http.request<resData>("post", "/cmsDetail/updateCmsDetail", { data });  
}

// @Tags CmsDetail
// @Summary 快速编辑
export const quickEdit = (data?: object) => {
   return http.request<resData>("post", "/cmsDetail/quickEdit", { data });  
}
 

// @Tags CmsDetail
// @Summary 用id查询CmsDetail
export const findCmsDetail = (params?: object) => {
   return http.request<resData>("get", "/cmsDetail/findCmsDetail", { params });  
}
 

// @Tags CmsDetail
// @Summary 分页获取CmsDetail列表
export const getCmsDetailList = (params?: object) => {
   return http.request<resData>("get", "/cmsDetail/getCmsDetailList", { params });  
}
 

// @Tags CmsDetail
// @Summary 导出excelCmsDetail列表
export const excelList = (params?: object) => {
   return http.request<resData>("get", "/cmsDetail/excelList", { params });  
}
