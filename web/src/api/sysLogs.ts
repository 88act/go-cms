import { http,resData } from "@/utils/http";



// @Tags SysLogs
// @Summary 创建SysLogs
export const createSysLogs = (data?: object) => {
  return http.request<resData>("post", "/sysLogs/createSysLogs", { data });
};

// @Tags SysLogs
// @Summary 删除SysLogs
export const deleteSysLogs = (data?: object) => {
  return http.request<resData>("post", "/sysLogs/deleteSysLogs", { data });
};


// @Tags SysLogs
// @Summary 删除SysLogs
export const deleteSysLogsByIds = (data?: object) => {
  return http.request<resData>("post", "/sysLogs/deleteSysLogsByIds", { data });
};


// @Tags SysLogs
// @Summary 更新SysLogs
export const updateSysLogs = (data?: object) => {
  return http.request<resData>("post", "/sysLogs/updateSysLogs", { data });
};


// @Tags SysLogs
// @Summary 用id查询SysLogs
export const findSysLogs =  (params?: object) => {
  return http.request<resData>("get", "/sysLogs/findSysLogs", { params });
};


// @Tags SysLogs
// @Summary 分页获取SysLogs列表
export const getSysLogsList =  (params?: object) => {
  return http.request<resData>("get", "/sysLogs/getSysLogsList", { params });
};




// @Tags SysLogs
// @Summary 快速编辑
export const quickEdit = (data?: object) => {
  return http.request<resData>("post", "/sysLogs/quickEdit", { data });
};


// @Tags SysLogs
// @Summary 导出excelSysLogs列表
export const excelList =(data?: object) => {
  return http.request<resData>("get", "/sysLogs/excelList", { data });
};




