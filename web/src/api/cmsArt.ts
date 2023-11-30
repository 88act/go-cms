import { http,resData } from "@/utils/http";
 
// @Tags CmsArt
// @Summary 创建CmsArt
export const createCmsArt = (data?: object) => {
   return http.request<resData>("post", "/cmsArt/createCmsArt", { data });  
}

// @Tags CmsArt
// @Summary 删除CmsArt
export const deleteCmsArt = (data?: object) => {
   return http.request<resData>("post", "/cmsArt/deleteCmsArt", { data });  
}
 

// @Tags CmsArt
// @Summary 删除ByIdsCmsArt
export const deleteCmsArtByIds = (data?: object) => {
   return http.request<resData>("post", "/cmsArt/deleteCmsArtByIds", { data });  
}
 

// @Tags CmsArt
// @Summary 更新CmsArt
export const updateCmsArt = (data?: object) => {
   return http.request<resData>("post", "/cmsArt/updateCmsArt", { data });  
}

// @Tags CmsArt
// @Summary 快速编辑
export const quickEdit = (data?: object) => {
   return http.request<resData>("post", "/cmsArt/quickEdit", { data });  
}
 

// @Tags CmsArt
// @Summary 用id查询CmsArt
export const findCmsArt = (params?: object) => {
   return http.request<resData>("get", "/cmsArt/findCmsArt", { params });  
}
 

// @Tags CmsArt
// @Summary 分页获取CmsArt列表
export const getCmsArtList = (params?: object) => {
   return http.request<resData>("get", "/cmsArt/getCmsArtList", { params });  
}
 

// @Tags CmsArt
// @Summary 导出excelCmsArt列表
export const excelList = (params?: object) => {
   return http.request<resData>("get", "/cmsArt/excelList", { params });  
}
