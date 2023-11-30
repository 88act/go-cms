import { http,resData } from "@/utils/http";
 
// @Tags CmsGroup
// @Summary 创建CmsGroup
export const createCmsGroup = (data?: object) => {
   return http.request<resData>("post", "/cmsGroup/createCmsGroup", { data });  
}

// @Tags CmsGroup
// @Summary 删除CmsGroup
export const deleteCmsGroup = (data?: object) => {
   return http.request<resData>("post", "/cmsGroup/deleteCmsGroup", { data });  
}
 

// @Tags CmsGroup
// @Summary 删除ByIdsCmsGroup
export const deleteCmsGroupByIds = (data?: object) => {
   return http.request<resData>("post", "/cmsGroup/deleteCmsGroupByIds", { data });  
}
 

// @Tags CmsGroup
// @Summary 更新CmsGroup
export const updateCmsGroup = (data?: object) => {
   return http.request<resData>("post", "/cmsGroup/updateCmsGroup", { data });  
}

// @Tags CmsGroup
// @Summary 快速编辑
export const quickEdit = (data?: object) => {
   return http.request<resData>("post", "/cmsGroup/quickEdit", { data });  
}
 

// @Tags CmsGroup
// @Summary 用id查询CmsGroup
export const findCmsGroup = (params?: object) => {
   return http.request<resData>("get", "/cmsGroup/findCmsGroup", { params });  
}
 

// @Tags CmsGroup
// @Summary 分页获取CmsGroup列表
export const getCmsGroupList = (params?: object) => {
   return http.request<resData>("get", "/cmsGroup/getCmsGroupList", { params });  
}
 

// @Tags CmsGroup
// @Summary 导出excelCmsGroup列表
export const excelList = (params?: object) => {
   return http.request<resData>("get", "/cmsGroup/excelList", { params });  
}
