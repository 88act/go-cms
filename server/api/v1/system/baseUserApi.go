package system

import (
	"go-cms/global"
	"go-cms/model/business"

	systemReq "go-cms/model/system/request"
	systemRes "go-cms/model/system/response"
	bizSev "go-cms/service/business"
	commSev "go-cms/service/common"
	"go-cms/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jaevor/go-nanoid"
	"github.com/jinzhu/copier"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

type BaseUserApi struct {
	commSev.BaseApi
}

// @Tags Base
// @Summary 用户登录
func (m *BaseUserApi) Login(c *gin.Context) {
	var loginUser business.Login
	_ = c.ShouldBindJSON(&loginUser)
	if err := utils.Verify(loginUser, utils.LoginVerify); err != nil {
		m.FailWithMessage(err.Error(), c)
		return
	}
	if store.Verify(loginUser.CaptchaId, loginUser.Captcha, true) {
		if err, userMini := userService.Login(c, loginUser); err == nil {
			m.tokenNext(c, userMini)
		} else {
			global.LOG.Error("登陆失败,"+err.Error(), zap.Any("err", err))
			m.FailWithMessage("登陆失败,"+err.Error(), c)
		}
	} else {
		m.FailWithMessage("验证码错误", c)
	}
}

// 登录以后签发jwt
func (m *BaseUserApi) tokenNext(c *gin.Context, user business.MemUserMini) {
	j := &utils.JWT{SigningKey: []byte(global.CONFIG.JWT.SigningKey)} // 唯一签名

	claims := systemReq.CustomClaims{
		UserId:   user.Id,
		CuGuid:   user.CuGuid,
		CuId:     user.CuId,
		Username: user.Username,
		Guid:     user.Guid,
		//Role:       user.Role,
		//RoleList:   user.RoleList,
		UserType:   user.UserType,
		BufferTime: global.CONFIG.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                          // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.CONFIG.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "go-cms",                                          // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.LOG.Error("获取token失败!", zap.Any("err", err))
		m.FailWithMessage("获取token失败", c)
		return
	}
	if !global.CONFIG.System.UseMultipoint {
		m.OkWithDetailed(business.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}

	// if err, jwtStr := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
	// 	if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
	// 		global.LOG.Error("设置登录状态失败!", zap.Any("err", err))
	// 		m.FailWithMessage("设置登录状态失败", c)
	// 		return
	// 	}
	// 	m.OkWithDetailed(business.LoginResponse{
	// 		User:      user,
	// 		Token:     token,
	// 		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	// 	}, "登录成功", c)
	// } else if err != nil {
	// 	global.LOG.Error("设置登录状态失败!", zap.Any("err", err))
	// 	m.FailWithMessage("设置登录状态失败", c)
	// } else {
	// 	var blackJWT system.JwtBlacklist
	// 	blackJWT.Jwt = jwtStr
	// 	if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
	// 		m.FailWithMessage("jwt作废失败", c)
	// 		return
	// 	}
	// 	if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
	// 		m.FailWithMessage("设置登录状态失败", c)
	// 		return
	// 	}
	// 	m.OkWithDetailed(business.LoginResponse{
	// 		User:      user,
	// 		Token:     token,
	// 		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	// 	}, "登录成功", c)
	// }
}

// @Tags SysUser
// @Summary 获取用户信息
func (m *BaseUserApi) GetUserInfo(c *gin.Context) {
	userId := utils.GetUserID(c)
	userInfo, err := bizSev.GetMemUserSev().Get(c, userId, "")
	if err == nil {
		if userInfo.Status != 1 {
			m.FailWithMessage("获取失败,用户状态异常", c)
		}
		var userMini business.MemUserMini
		_ = copier.Copy(&userMini, userInfo)
		userMini.CuGuid = utils.GetCuGuid(c)
		m.OkWithDetailed(gin.H{"userInfo": userMini}, "获取成功", c)
	} else {
		m.FailWithMessage("获取失败"+err.Error(), c)
	}
}

// @Summary 用户修改密码
func (m *BaseUserApi) ChangePwd(c *gin.Context) {
	var userCPW business.ChangePasswordStruct
	_ = c.ShouldBindJSON(&userCPW)
	if len(userCPW.NewPassword) < 6 {
		m.FailWithMessage("新密码长度不对", c)
		c.Abort()
		return
	}
	//如果是修改其他人， 必须是 admin id = 1
	bePlatAdmin := utils.BePlatAdmin(c)
	userId := utils.GetUserID(c)
	if userCPW.UserId > 0 && userCPW.UserId != userId {
		if !bePlatAdmin {
			m.FailWithMessage("修改其他人密码，必须为管理员用户", c)
			c.Abort()
			return
		}
	} else {
		// 改自己的
		userCPW.UserId = userId
	}
	if err := bizSev.GetMemUserSev().ChangePassword(c, userCPW, bePlatAdmin); err == nil {
		m.OkWithMessage("修改成功", c)
	} else {
		m.FailWithMessage(err.Error(), c)
	}
}

// @Summary 用户修改密码
func (m *BaseUserApi) ChangePwdMy(c *gin.Context) {
	var userCPW business.ChangePasswordStruct
	_ = c.ShouldBindJSON(&userCPW)
	if len(userCPW.NewPassword) < 6 {
		m.FailWithMessage("新密码长度不对", c)
		c.Abort()
		return
	}
	userCPW.UserId = utils.GetUserID(c)
	if err := bizSev.GetMemUserSev().ChangePassword(c, userCPW, false); err == nil {
		m.OkWithMessage("修改成功", c)
	} else {
		m.FailWithMessage(err.Error(), c)
	}
}

func (m *BaseUserApi) ShowQrcode(c *gin.Context) {
	var mUser business.MemUser
	_ = c.ShouldBindJSON(&mUser)
	if mUser.Id == 0 {
		m.FailWithMessage("参数错误", c)
	}

	memUser, err := bizSev.GetMemUserSev().Get(c, mUser.Id, "")

	cuGuid := utils.GetCuGuid(c)
	if err != nil {
		m.FailWithMessage("用户不存在", c)
	} else {
		decenaryID, err := nanoid.CustomASCII("0123456789", 12)
		if err != nil {
			global.LOG.Error(err.Error())
		}
		scan := decenaryID()
		mapData := make(map[string]interface{})
		mapData["scan"] = scan
		mapData["scan_time"] = time.Now()

		if err := bizSev.GetMemUserSev().UpdateByMap(c, memUser, mapData); err != nil {
			global.LOG.Error("生成二维码失败!", zap.Any("err", err))
			m.FailWithMessage("生成二维码失败", c)
		} else {
			//格式 scan,域名,rtcType,cuGuid
			scan = scan + "," + global.CONFIG.System.ApiUrl + "," + global.CONFIG.License.RtcType + "," + cuGuid
			m.OkWithData(scan, c)
		}
	}

}

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
// var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

// @Tags Base
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证码获取成功"}"
// @Router /base/captcha [post]
func (m *BaseUserApi) Captcha(c *gin.Context) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.CONFIG.Captcha.ImgHeight, global.CONFIG.Captcha.ImgWidth, global.CONFIG.Captcha.KeyLong, 0.7, 80)
	//cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		global.LOG.Error("验证码获取失败!", zap.Any("err", err))
		m.FailWithMessage("验证码获取失败", c)
	} else {
		m.OkWithDetailed(systemRes.SysCaptchaResponse{
			CaptchaId: id,
			PicPath:   b64s,
		}, "验证码获取成功", c)
	}
}
