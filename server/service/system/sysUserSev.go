package system

import (
	"errors"
	"fmt"

	"go-cms/model/business"
	"go-cms/utils"

	bizSev "go-cms/service/business"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

//@author:  [linjd] 10512203@qq.com
//@function: Register
//@description: 用户注册
//@param: u model.SysUser
//@return: err error, userInter model.SysUser

var PlatCuGuid = "go-cms000"

type UserService struct {
}

// @author:  [linjd] 10512203@qq.com
// @function: Login
// @description: 用户登录
func (userService *UserService) Login(c *gin.Context, req business.Login) (err error, user business.MemUserMini) {

	memUser, err := bizSev.GetMemUserSev().GetByUsername(c, req.Username, "")
	if err == nil {
		fmt.Println(memUser)
		pwd := req.Password + memUser.PasswordSlat
		if memUser.Password != utils.Md5ByString(pwd) {
			return errors.New("密码错误"), user
		}
		if memUser.Status != 1 {
			return errors.New("用户状态异常"), user
		}
		_ = copier.Copy(&user, memUser)
		user.CuGuid = req.CuGuid
		return nil, user
	}
	return errors.New("用户不存在"), user
}
