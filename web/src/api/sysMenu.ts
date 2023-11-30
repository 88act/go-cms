import { http,resData } from "@/utils/http";
 

// @Tags SysMenu
// @Summary 创建SysMenu
export const createSysMenu = (data?: object) => {
  return http.request<resData>("post", "/sysMenu/createSysMenu", { data });
};

// @Tags SysMenu
// @Summary 删除SysMenu
export const deleteSysMenu = (data?: object) => {
  return http.request<resData>("post", "/sysMenu/deleteSysMenu", { data });
};


// @Tags SysMenu
// @Summary 删除SysMenu
export const deleteSysMenuByIds = (data?: object) => {
  return http.request<resData>("post", "/sysMenu/deleteSysMenuByIds", { data });
};


// @Tags SysMenu
// @Summary 更新SysMenu
export const updateSysMenu = (data?: object) => {
  return http.request<resData>("post", "/sysMenu/updateSysMenu", { data });
};


// @Tags SysMenu
// @Summary 用id查询SysMenu
export const findSysMenu =  (params?: object) => {
  return http.request<resData>("get", "/sysMenu/findSysMenu", { params });
};


// @Tags SysMenu
// @Summary 分页获取SysMenu列表
export const getSysMenuList =  (params?: object) => {
  return http.request<resData>("get", "/sysMenu/getSysMenuList", { params });
};




// @Tags SysMenu
// @Summary 快速编辑
export const quickEdit = (data?: object) => {
  return http.request<resData>("post", "/sysMenu/quickEdit", { data });
};


// @Tags SysMenu
// @Summary 导出excelSysMenu列表
export const excelList =(data?: object) => {
  return http.request<resData>("get", "/sysMenu/excelList", { data });
};

