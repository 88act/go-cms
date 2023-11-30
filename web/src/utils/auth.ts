import Cookies from "js-cookie";
import { storageLocal } from "@pureadmin/utils";
import { useUserStoreHook } from "@/store/modules/user";

export interface DataInfo<T> {
  /** token */
  accessToken : string;
  /** `accessToken`的过期时间（时间戳） */
  expires : T;
  /** 用于调用刷新accessToken的接口时所需的token */
  refreshToken : string;
  /** 用户id */
  userId ?: number;
  /** 用户名 */
  username ?: string;
  /** 当前登陆用户的角色 */
  roles ?: Array<string>;
  userType ?: number;
}

export const userKey = "user-info";
export const TokenKey = "authorized-token";
/**
 * 通过`multiple-tabs`是否在`cookie`中，判断用户是否已经登录系统，
 * 从而支持多标签页打开已经登录的系统后无需再登录。
 * 浏览器完全关闭后`multiple-tabs`将自动从`cookie`中销毁，
 * 再次打开浏览器需要重新登录系统
 * */
export const multipleTabsKey = "multiple-tabs";

/** 获取`token` */
export function getToken() : DataInfo<number> {
  // 此处与`TokenKey`相同，此写法解决初始化时`Cookies`中不存在`TokenKey`报错
  return Cookies.get(TokenKey)
    ? JSON.parse(Cookies.get(TokenKey))
    : storageLocal().getItem(userKey);
}

/**
 * @description 设置`token`以及一些必要信息并采用无感刷新`token`方案
 * 无感刷新：后端返回`accessToken`（访问接口使用的`token`）、`refreshToken`（用于调用刷新`accessToken`的接口时所需的`token`，`refreshToken`的过期时间（比如30天）应大于`accessToken`的过期时间（比如2小时））、`expires`（`accessToken`的过期时间）
 * 将`accessToken`、`expires`这两条信息放在key值为authorized-token的cookie里（过期自动销毁）
 * 将`username`、`roles`、`refreshToken`、`expires`这四条信息放在key值为`user-info`的localStorage里（利用`multipleTabsKey`当浏览器完全关闭后自动销毁）
 */
export function setToken(data : DataInfo<Date>) {
   const { isRemembered, loginDay } = useUserStoreHook();
  console.log("setToken===data=", data)
  let expires = data.expires;
  let accessToken = data.accessToken;
  let refreshToken = "refreshToken需后端返回";
  // let expires = 0;
  // const { accessToken, refreshToken } = data;
  // expires = new Date(data.expires).getTime(); // 如果后端直接设置时间戳，将此处代码改为expires = data.expires，然后把上面的DataInfo<Date>改成DataInfo<number>即可
  let cookieString = JSON.stringify({ accessToken, expires });
  expires > 0
    ? Cookies.set(TokenKey, cookieString, {
      expires: (expires - Date.now()) / 86400000
    })
    : Cookies.set(TokenKey, cookieString);

  Cookies.set(
    multipleTabsKey,
    "true",
    isRemembered
      ? {
        expires: loginDay
      }
      : {}
  );

  if (data.userId>0 && data.username!="") {
    console.log("密码登录后获得 data =",data)
    const { userId, username, roles,userType } = data;
    setUserKey(userId, username, roles,userType);
  } else {
    console.log("从cookie 获取信息登录 data =",data)
    const username = storageLocal().getItem<DataInfo<number>>(userKey)?.username ?? "";
    const roles = storageLocal().getItem<DataInfo<number>>(userKey)?.roles ?? [];
    let userIdStr : string = storageLocal().getItem<DataInfo<number>>(userKey)?.userId ?? "0";
    let userId : number = Number(userIdStr)
    let userTypeStr : string = storageLocal().getItem<DataInfo<number>>(userKey)?.userType ?? "0";
    let userType : number = Number(userTypeStr)
    console.log("从cookie 获取信息 登录", userId, username, roles)
    setUserKey(userId, username, roles,userType);
  }

  function setUserKey(userId : number,  username : string, roles : Array<string>,userType : number) {
    console.log("保存 cookie setUserKey ， userId ，username  roles ", userId, username, roles)
    useUserStoreHook().SET_USERNAME(username);
    useUserStoreHook().SET_USERID(userId);
    useUserStoreHook().SET_ROLES(roles);
    useUserStoreHook().SET_USERTYPE(userType);
    storageLocal().setItem(userKey, {
      refreshToken,
      expires,
      username,
      userId,
      roles,
      userType,
    });
  }

}

/** 删除`token`以及key值为`user-info`的localStorage信息 */
export function removeToken() {
  Cookies.remove(TokenKey);
  Cookies.remove(multipleTabsKey);
  storageLocal().removeItem(userKey);
}

/** 格式化token（jwt格式） */
export const formatToken = (token : string) : string => {
  // return "Bearer " + token;
  return token
};
