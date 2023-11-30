import { http,resData } from "@/utils/http";
 
// @Tags CmsTag
// @Summary 创建CmsTag
export const createCmsTag = (data?: object) => {
   return http.request<resData>("post", "/cmsTag/createCmsTag", { data });  
}

// @Tags CmsTag
// @Summary 删除CmsTag
export const deleteCmsTag = (data?: object) => {
   return http.request<resData>("post", "/cmsTag/deleteCmsTag", { data });  
}
 

// @Tags CmsTag
// @Summary 删除ByIdsCmsTag
export const deleteCmsTagByIds = (data?: object) => {
   return http.request<resData>("post", "/cmsTag/deleteCmsTagByIds", { data });  
}
 

// @Tags CmsTag
// @Summary 更新CmsTag
export const updateCmsTag = (data?: object) => {
   return http.request<resData>("post", "/cmsTag/updateCmsTag", { data });  
}

// @Tags CmsTag
// @Summary 快速编辑
export const quickEdit = (data?: object) => {
   return http.request<resData>("post", "/cmsTag/quickEdit", { data });  
}
 

// @Tags CmsTag
// @Summary 用id查询CmsTag
export const findCmsTag = (params?: object) => {
   return http.request<resData>("get", "/cmsTag/findCmsTag", { params });  
}
 

// @Tags CmsTag
// @Summary 分页获取CmsTag列表
export const getCmsTagList = (params?: object) => {
   return http.request<resData>("get", "/cmsTag/getCmsTagList", { params });  
}
 

// @Tags CmsTag
// @Summary 导出excelCmsTag列表
export const excelList = (params?: object) => {
   return http.request<resData>("get", "/cmsTag/excelList", { params });  
}
