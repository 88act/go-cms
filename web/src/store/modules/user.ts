import { defineStore } from "pinia";
import { store } from "@/store";
import { userType } from "./types";
import { routerArrays } from "@/layout/types";
import { router, resetRouter } from "@/router";
import { storageLocal } from "@pureadmin/utils";
import { getConfig } from "@/config";
import type { App } from "vue";
import { getLogin, getUserInfo,refreshTokenApi } from "@/api/base";
import { resData } from "@/utils/http";
//import { UserResult, RefreshTokenResult } from "@/api/base";
import { useMultiTagsStoreHook } from "@/store/modules/multiTags";
import { type DataInfo, setToken, removeToken, userKey } from "@/utils/auth";

export const useUserStore = defineStore({
  id: "pure-user",
  state: (): userType => ({

    userId: storageLocal().getItem<DataInfo<number>>(userKey)?.userId ?? 0,
    cuId: 0,
    cuGuid: "",
    userType: 0,
    userInfo:null,
    // 用户名
    username: storageLocal().getItem<DataInfo<number>>(userKey)?.username ?? "",
    // 页面级别权限
    roles: storageLocal().getItem<DataInfo<number>>(userKey)?.roles ?? [],
    // 前端生成的验证码（按实际需求替换）
    verifyCode: "",
    // 判断登录页面显示哪个组件（0：登录（默认）、1：手机登录、2：二维码登录、3：注册、4：忘记密码）
    currentPage: 0,
    // 是否勾选了登录页的免登录
    isRemembered: false,
    // 登录页的免登录存储几天，默认7天
    loginDay: 7,

  }),
  actions: {
     GetUserType() {
       let userType = this.userType
       if(!userType || userType ==0 )  {
          userType = storageLocal().getItem<DataInfo<number>>(userKey)?.userType
          this.userType = userType
          console.log("从cookie 读取userType= ",userType)
       }
       return  userType
     },
     SET_USERID(userId: number) {
       this.userId = userId;
     },
     SET_USERINFO(info:any) {
       this.userInfo = info;
     },

    SET_USERTYPE(userType: number) {
      this.userType = userType;
    },

    /** 存储用户名 */
    SET_USERNAME(username: string) {
      this.username = username;
    },
    /** 存储角色 */
    SET_ROLES(roles: Array<string>) {
      this.roles = roles;
    },
    /** 存储前端生成的验证码 */
    SET_VERIFYCODE(verifyCode: string) {
      this.verifyCode = verifyCode;
    },
    /** 存储登录页面显示哪个组件 */
    SET_CURRENTPAGE(value: number) {
      this.currentPage = value;
    },
    /** 存储是否勾选了登录页的免登录 */
    SET_ISREMEMBERED(bool: boolean) {
      this.isRemembered = bool;
    },
    /** 设置登录页的免登录存储几天 */
    SET_LOGINDAY(value: number) {
      this.loginDay = Number(value);
    },
    // /** 登入 */
    // async loginByUsername(data) {
    //   return new Promise<UserResult>((resolve, reject) => {
    //     getLogin(data)
    //       .then(data => {
    //         if (data) {
    //           setToken(data.data);
    //           resolve(data);
    //         }
    //       })
    //       .catch(error => {
    //         reject(error);
    //       });
    //   });
    // },


    /** 登入 */
      async loginByUsername(data) {
        return new Promise<resData>((resolve, reject) => {
          getLogin(data)
            .then(data => {
             //  console.log("========data====data==", data);
              if (data.code==200) {
                const resData = data.data;
                console.log("============resData==", resData);
                this.SET_USERINFO(resData.user);

                const newData:any = new Object();
                newData.accessToken = resData.token;
                newData.expires = resData.expiresAt;
                newData.username = resData.user.username;
                newData.userId = resData.user.id;
                // newData.wssurl = resData.user.sideMode;
                newData.roles = resData.user.roleList
                newData.userType = resData.user.userType

                // /** token */
                // accessToken: string;
                // /** `accessToken`的过期时间（时间戳） */
                // expires: T;
                // /** 用于调用刷新accessToken的接口时所需的token */
                // refreshToken: string;
                // /** 用户名 */
                // username?: string;
                // /** 当前登陆用户的角色 */
                // roles?: Array<string>;
                setToken(newData);
                console.log("====2============");
                resolve(data);
              }else {
                 resolve(data);
              }
            })
            .catch(error => {
              reject(error);
            });
        });
      },

    async loginByToken(token) {
      return new Promise<resData>(async (resolve, reject) => {
              const newData = new Object();
              newData.accessToken = token.token;
              let exp = Number(token.expiresAt)
              if (exp <=0 ){
                 // 30天
                 exp = Date.now()+ 2592000
              }
              newData.expires = exp
              newData.username = token.username;
              newData.userId = token.userId;
              newData.roles = ["3"];
              // /** token */
              // accessToken: string;
              // /** `accessToken`的过期时间（时间戳） */
              // expires: T;
              // /** 用于调用刷新accessToken的接口时所需的token */
              // refreshToken: string;
              // /** 用户名 */
              // username?: string;
              // /** 当前登陆用户的角色 */
              // roles?: Array<string>;
              setToken(newData);
              await this.GetUserInfo()
              resolve(token);
      });
    },
  /* 获取用户信息*/
  async GetUserInfo() {
    const res = await getUserInfo()
    if (res.code == 200) {

       let userInfo = res.data.userInfo
       // if (getConfig().Plat== "jq" && userInfo.userType != 9 ){
       //      console.log("平台用户登录 userInfo=",userInfo)
       //      router.push("/login");
       // }else{
          this.SET_USERINFO(userInfo);
          this.SET_USERNAME(userInfo.username);
          this.SET_USERID(userInfo.id);
          this.SET_USERTYPE(userInfo.userType);
         // this().SET_ROLES(userInfo.roles);
       //}
    }
    return res
  },
    /** 前端登出（不调用接口） */
    logOut() {
      this.username = "";
      this.roles = [];
      removeToken();
      useMultiTagsStoreHook().handleTags("equal", [...routerArrays]);
      resetRouter();
      router.push("/login");
    },
    /** 刷新`token` */
    async handRefreshToken(data) {
      return new Promise<resData>((resolve, reject) => {
        refreshTokenApi(data)
          .then(data => {
            if (data) {
              setToken(data.data);
              resolve(data);
            }
          })
          .catch(error => {
            reject(error);
          });
      });
    }
  }
});

export function useUserStoreHook() {
  return useUserStore(store);
}
