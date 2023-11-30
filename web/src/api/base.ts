import { http,resData } from "@/utils/http";

// @Tags base
// 登录
export const getLogin = (data?: object) => {
  return http.request<resData>("post", "/base/login", { data });
};

// @Tags base
// captcha
export const captcha = (data?: object) => {
  return http.request<resData>("post", "/base/captcha", { data });
};

// @Tags base
// 刷新token
export const refreshTokenApi = (data?: object) => {
  return http.request<resData>("post", "/base/refreshToken", { data });
};

// @Tags base
//用户资料
export const getUserInfo = () => {
  return http.request<resData>("get", "/base/getUserInfo");
};

// @Tags base
//  获取服务器运行状态
export const getSystemState = () => {
  return http.request<resData>("get", "/base/getServerInfo");
 };

// @Tags base
// 菜单
export const getMenu = () => {
  // return http.request<any>("get", "/getAsyncRoutes");
   return http.request<resData>("get", "/base/getMenu");
};



// @Tags base
// @Summary 修改密码
export const changePwd = (data?: object) => {
   return http.request<resData>("post", "/base/changePwd", { data });
}

// @Tags base
// @Summary 修改自己密码
export const changePwdMy = (data?: object) => {
   return http.request<resData>("post", "/base/changePwdMy", { data });
}


// @Tags base
// @Summary 登录二维码
export const showQrcode = (data?: object) => {
   return http.request<resData>("post", "/base/showQrcode", { data });
}
