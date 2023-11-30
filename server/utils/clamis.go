package utils

import (
	"fmt"
	"go-cms/global"
	systemReq "go-cms/model/system/request"

	"github.com/gin-gonic/gin"
)

// 获取从jwt解析出来的用户ID
func GetUserID(c *gin.Context) int64 {
	if claims, exists := c.Get("claims"); !exists {
		global.LOG.Error("从Gin的Context中获取从jwt解析出来的用户ID失败, 请检查路由是否使用jwt中间件!")
		return 0
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.UserId
	}
}

// 获取从jwt解析出来的用户type
func GetUserType(c *gin.Context) int {
	if claims, exists := c.Get("claims"); !exists {
		global.LOG.Error("从Gin的Context中获取从jwt解析出来的用户ID失败, 请检查路由是否使用jwt中间件!")
		return 0
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.UserType
	}
}

// 获取从jwt解析出来的用户 GetCuId
func GetCuId(c *gin.Context) int64 {
	if claims, exists := c.Get("claims"); !exists {
		global.LOG.Error("从Gin的Context中获取从jwt解析 Guid 失败, 请检查路由是否使用jwt中间件!")
		return 0
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.CuId
	}
}

// 获取从jwt解析出来的用户 GetCuGuid
func GetCuGuid(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		global.LOG.Error("从Gin的Context中获取从jwt解析 Guid 失败, 请检查路由是否使用jwt中间件!")
		return ""
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.CuGuid
	}
}

// 获取从jwt解析出来的用户 Guid
func GetUserGuid(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		global.LOG.Error("从Gin的Context中获取从jwt解析 Guid 失败, 请检查路由是否使用jwt中间件!")
		return ""
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.Guid
	}
}

// 获取从jwt解析出来的用户角色id
// func GetUserRoleId(c *gin.Context) int64 {
// 	if claims, exists := c.Get("claims"); !exists {
// 		global.LOG.Error("从Gin的Context中获取从jwt解析出来用户角色id 失败, 请检查路由是否使用jwt中间件!")
// 		return 0
// 	} else {
// 		waitUse := claims.(*systemReq.CustomClaims)
// 		return waitUse.Role
// 	}
// }

// 获取从jwt解析出来的用户角色id
func GetUserInfo(c *gin.Context) *systemReq.CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		global.LOG.Error("从Gin的Context中获取从jwt解析出来的用户UserInfo失败, 请检查路由是否使用jwt中间件!")
		return nil
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse
	}
}

// 是否管理员 UserType == 9
func BePlatAdmin(c *gin.Context) bool {
	beAdmin := false
	claims := GetUserInfo(c)
	fmt.Println("claims", claims)
	if claims.UserType == 9 { // 1 普通员工 2 商户子帐号  3 商户管理员  9  平台管理员
		beAdmin = true
	}
	return beAdmin
}

// 是否商户管理员 UserType == 3
// func BeCuAdmin(c *gin.Context) bool {
// 	beAdmin := false
// 	claims := GetUserInfo(c)
// 	fmt.Println("claims", claims)
// 	if claims.UserType == 3 { // 1 普通员工 2 商户子帐号  3 商户管理员  9  平台管理员
// 		beAdmin = true
// 	}
// 	return beAdmin
// }
