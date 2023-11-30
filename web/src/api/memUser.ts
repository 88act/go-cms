import { http,resData } from "@/utils/http";

// @Tags MemUser
// @Summary 创建MemUser
export const createMemUser = (data?: object) => {
   return http.request<resData>("post", "/memUser/createMemUser", { data });
}

// @Tags MemUser
// @Summary 删除MemUser
export const deleteMemUser = (data?: object) => {
   return http.request<resData>("post", "/memUser/deleteMemUser", { data });
}


// @Tags MemUser
// @Summary 删除ByIdsMemUser
export const deleteMemUserByIds = (data?: object) => {
   return http.request<resData>("post", "/memUser/deleteMemUserByIds", { data });
}


// @Tags MemUser
// @Summary 更新MemUser
export const updateMemUser = (data?: object) => {
   return http.request<resData>("post", "/memUser/updateMemUser", { data });
}

// @Tags MemUser
// @Summary 快速编辑
export const quickEdit = (data?: object) => {
   return http.request<resData>("post", "/memUser/quickEdit", { data });
}


// @Tags MemUser
// @Summary 用id查询MemUser
export const findMemUser = (params?: object) => {
   return http.request<resData>("get", "/memUser/findMemUser", { params });
}


// @Tags MemUser
// @Summary 分页获取MemUser列表
export const getMemUserList = (params?: object) => {
   return http.request<resData>("get", "/memUser/getMemUserList", { params });
}


// @Tags MemUser
// @Summary 导出excelMemUser列表
export const excelList = (params?: object) => {
   return http.request<resData>("get", "/memUser/excelList", { params });
}


// @Tags jqCustomerPwd
// @Summary 修改密码
export const memChangePwd = (data?: object) => {
   return http.request<resData>("post", "/memUser/memChangePwd", { data });
}


