import { http,resData } from "@/utils/http";

 
export const pidData = (data?: object) => {
  return http.request<resData>("post", "/commDb/pidData", { data });
}; 


export const pidTreeData = (data?: object) => {
  return http.request<resData>("post", "/commDb/pidTreeData", { data });
};
  

export const getDictData = (data?: object) => {
  return http.request<resData>("post", "/commDb/getDict", { data });
};
 
 

export const getDictData2 = (data?: object) => {
  return http.request<resData>("post", "/commDb/getDict2", { data });
};
 
 


export const getDictTreeData = (data?: object) => {
  return http.request<resData>("post", "/commDb/getDictTree", { data });
};
  

export const updatePidSort = (data?: object) => {
  return http.request<resData>("post", "/commDb/updatePidSort", { data });
};
 