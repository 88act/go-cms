package request

import (
	"github.com/dgrijalva/jwt-go"
)

// Custom claims structure
type CustomClaims struct {
	UserId   int64  `json:"userId"  form:"userId" `
	CuId     int64  `json:"cuId"  form:"cuId" `
	CuGuid   string `json:"cuGuid"  form:"cuGuid" `
	Guid     string `json:"guid"   form:"guid" `
	Username string `json:"username"   form:"username" `
	//RoleList   string `json:"roleList"   form:"roleList" `
	//Role       int64  `json:"role"  form:"role" `
	UserType   int   `json:"userType"   form:"userType" `
	BufferTime int64 // 刷新时间
	jwt.StandardClaims
}
