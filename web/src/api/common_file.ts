import { httpUp } from "@/utils/httpupload";
import { http,resData } from "@/utils/http";

export const getFileByKey = (params) => {
  return http.request<resData>("get", "/commFile/getFileByKey", { params });
}


export const uploadFile = (data) => {
   return httpUp.request<resData>("post", "/commFile/upload", { data });
}

 /**  myUpload */
 export const myUpload = (data?: object) => {
   return httpUp.request<resData>("post", "/commFile/upload", { data });
};

