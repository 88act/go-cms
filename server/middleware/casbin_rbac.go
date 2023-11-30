package middleware

import (
	"github.com/gin-gonic/gin"
)

//var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

// 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 低安全， 只限制菜单 ，不检查 api
		c.Next()
		/*
			if global.CONFIG.System.SuperAdmin {
				c.Next()
			} else {
				claims, _ := c.Get("claims")
				waitUse := claims.(*request.CustomClaims)
				// 获取请求的URI
				obj := c.Request.URL.RequestURI()
				// 获取请求方法
				act := c.Request.Method
				// 获取用户的角色
				sub := waitUse.AuthorityId
				e := casbinService.Casbin()
				// 判断策略中是否存在

				//fmt.Println("判断策略中是否存在")
				//fmt.Println(obj)
				if obj == "/menu/getMenu" {
					c.Next()
				} else {
					success, _ := e.Enforce(sub, obj, act)
					if success {
						c.Next()
					} else {
						response.FailWithDetailed(gin.H{}, "权限不足", c)
						c.Abort()
						return
					}
				}

			}
		*/
	}
}
