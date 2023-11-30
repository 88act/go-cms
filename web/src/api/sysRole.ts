import { http,resData } from "@/utils/http";


// @Tags SysRole
// @Summary 创建SysRole
export const createSysRole = (data?: object) => {
  return http.request<resData>("post", "/sysRole/createSysRole", { data });
};

// @Tags SysRole
// @Summary 删除SysRole
export const deleteSysRole = (data?: object) => {
  return http.request<resData>("post", "/sysRole/deleteSysRole", { data });
};


// @Tags SysRole
// @Summary 删除SysRole
export const deleteSysRoleByIds = (data?: object) => {
  return http.request<resData>("post", "/sysRole/deleteSysRoleByIds", { data });
};


// @Tags SysRole
// @Summary 更新SysRole
export const updateSysRole = (data?: object) => {
  return http.request<resData>("post", "/sysRole/updateSysRole", { data });
};


// @Tags SysRole
// @Summary 用id查询SysRole
export const findSysRole =  (params?: object) => {
  return http.request<resData>("get", "/sysRole/findSysRole", { params });
};


// @Tags SysRole
// @Summary 分页获取SysRole列表
export const getSysRoleList =  (params?: object) => {
  return http.request<resData>("get", "/sysRole/getSysRoleList", { params });
};




// @Tags SysRole
// @Summary 快速编辑
export const quickEdit = (data?: object) => {
  return http.request<resData>("post", "/sysRole/quickEdit", { data });
};


// @Tags SysRole
// @Summary 导出excelSysRole列表
export const excelList =(data?: object) => {
  return http.request<resData>("get", "/sysRole/excelList", { data });
};

