import { http,resData } from "@/utils/http";



// @Tags SysApis
// @Summary 创建SysApis
export const createSysApis = (data?: object) => {
  return http.request<resData>("post", "/sysApis/createSysApis", { data });
};

// @Tags SysApis
// @Summary 删除SysApis
export const deleteSysApis = (data?: object) => {
  return http.request<resData>("post", "/sysApis/deleteSysApis", { data });
};


// @Tags SysApis
// @Summary 删除SysApis
export const deleteSysApisByIds = (data?: object) => {
  return http.request<resData>("post", "/sysApis/deleteSysApisByIds", { data });
};


// @Tags SysApis
// @Summary 更新SysApis
export const updateSysApis = (data?: object) => {
  return http.request<resData>("post", "/sysApis/updateSysApis", { data });
};


// @Tags SysApis
// @Summary 用id查询SysApis
export const findSysApis =  (params?: object) => {
  return http.request<resData>("get", "/sysApis/findSysApis", { params });
};


// @Tags SysApis
// @Summary 分页获取SysApis列表
export const getSysApisList =  (params?: object) => {
  return http.request<resData>("get", "/sysApis/getSysApisList", { params });
};




// @Tags SysApis
// @Summary 快速编辑
export const quickEdit = (data?: object) => {
  return http.request<resData>("post", "/sysApis/quickEdit", { data });
};


// @Tags SysApis
// @Summary 导出excelSysApis列表
export const excelList =(data?: object) => {
  return http.request<resData>("get", "/sysApis/excelList", { data });
};




