import { http,resData } from "@/utils/http";
 
// @Tags CmsVisitor
// @Summary 创建CmsVisitor
export const createCmsVisitor = (data?: object) => {
   return http.request<resData>("post", "/cmsVisitor/createCmsVisitor", { data });  
}

// @Tags CmsVisitor
// @Summary 删除CmsVisitor
export const deleteCmsVisitor = (data?: object) => {
   return http.request<resData>("post", "/cmsVisitor/deleteCmsVisitor", { data });  
}
 

// @Tags CmsVisitor
// @Summary 删除ByIdsCmsVisitor
export const deleteCmsVisitorByIds = (data?: object) => {
   return http.request<resData>("post", "/cmsVisitor/deleteCmsVisitorByIds", { data });  
}
 

// @Tags CmsVisitor
// @Summary 更新CmsVisitor
export const updateCmsVisitor = (data?: object) => {
   return http.request<resData>("post", "/cmsVisitor/updateCmsVisitor", { data });  
}

// @Tags CmsVisitor
// @Summary 快速编辑
export const quickEdit = (data?: object) => {
   return http.request<resData>("post", "/cmsVisitor/quickEdit", { data });  
}
 

// @Tags CmsVisitor
// @Summary 用id查询CmsVisitor
export const findCmsVisitor = (params?: object) => {
   return http.request<resData>("get", "/cmsVisitor/findCmsVisitor", { params });  
}
 

// @Tags CmsVisitor
// @Summary 分页获取CmsVisitor列表
export const getCmsVisitorList = (params?: object) => {
   return http.request<resData>("get", "/cmsVisitor/getCmsVisitorList", { params });  
}
 

// @Tags CmsVisitor
// @Summary 导出excelCmsVisitor列表
export const excelList = (params?: object) => {
   return http.request<resData>("get", "/cmsVisitor/excelList", { params });  
}
