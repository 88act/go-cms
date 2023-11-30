import { http,resData } from "@/utils/http";
 
// @Tags CmsDiscuss
// @Summary 创建CmsDiscuss
export const createCmsDiscuss = (data?: object) => {
   return http.request<resData>("post", "/cmsDiscuss/createCmsDiscuss", { data });  
}

// @Tags CmsDiscuss
// @Summary 删除CmsDiscuss
export const deleteCmsDiscuss = (data?: object) => {
   return http.request<resData>("post", "/cmsDiscuss/deleteCmsDiscuss", { data });  
}
 

// @Tags CmsDiscuss
// @Summary 删除ByIdsCmsDiscuss
export const deleteCmsDiscussByIds = (data?: object) => {
   return http.request<resData>("post", "/cmsDiscuss/deleteCmsDiscussByIds", { data });  
}
 

// @Tags CmsDiscuss
// @Summary 更新CmsDiscuss
export const updateCmsDiscuss = (data?: object) => {
   return http.request<resData>("post", "/cmsDiscuss/updateCmsDiscuss", { data });  
}

// @Tags CmsDiscuss
// @Summary 快速编辑
export const quickEdit = (data?: object) => {
   return http.request<resData>("post", "/cmsDiscuss/quickEdit", { data });  
}
 

// @Tags CmsDiscuss
// @Summary 用id查询CmsDiscuss
export const findCmsDiscuss = (params?: object) => {
   return http.request<resData>("get", "/cmsDiscuss/findCmsDiscuss", { params });  
}
 

// @Tags CmsDiscuss
// @Summary 分页获取CmsDiscuss列表
export const getCmsDiscussList = (params?: object) => {
   return http.request<resData>("get", "/cmsDiscuss/getCmsDiscussList", { params });  
}
 

// @Tags CmsDiscuss
// @Summary 导出excelCmsDiscuss列表
export const excelList = (params?: object) => {
   return http.request<resData>("get", "/cmsDiscuss/excelList", { params });  
}
