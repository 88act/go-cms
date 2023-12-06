package logic

import (
	"context"

	"go-cms/app/usercenter/cmd/rpc/internal/svc"
	"go-cms/app/usercenter/cmd/rpc/usercenter"
	"go-cms/common/tool"
	"go-cms/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrGenerateTokenError = xerr.NewErrMsg("生成token失败")
var ErrUsernamePwdError = xerr.NewErrMsg("账号或密码不正确")

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *usercenter.LoginReq) (*usercenter.LoginResp, error) {
	var userId int64
	var err error
	switch in.LoginType { //注册登录类型  1 用户名密码 2 邮箱密码 3 手机/验证码  4 微信 5 QQ 6 github 7 gmail   8 piko
	case 1:
		userId, err = l.loginByUsername(in.Username, in.Password)
	case 2:
		userId, err = l.loginByMobile(in.Mobile, in.Password)
	default:
		return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
	}
	if err != nil {
		return nil, err
	}

	//2、Generate the token, so that the service doesn't call rpc internally
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&usercenter.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId : %d", userId)
	}
	return &usercenter.LoginResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}

func (l *LoginLogic) loginByMobile(mobile, password string) (int64, error) {
	logx.Errorf("loginByMobile begin....")
	mapData := map[string]interface{}{"mobile": mobile}
	if user, err := l.svcCtx.MemUserSev.GetByMap(l.ctx, mapData, ""); err == nil {
		if !(tool.Md5ByString(password+user.PasswordSlat) == user.Password) {
			return 0, errors.Wrap(ErrUsernamePwdError, "密码匹配出错")
		}
		if user.Status != 1 {
			return 0, errors.Wrap(xerr.NewErrCode(xerr.Fail), "用户状态异常, user.Status")
		}
		return user.Id, nil
	} else {
		if err == gorm.ErrRecordNotFound {
			return 0, errors.Wrapf(ErrUserNoExistsError, " mobile=%s", mobile)
		} else {
			return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "loginByMobile db err ,  err:%v", err.Error())
		}
	}
}

func (l *LoginLogic) loginByUsername(username, password string) (int64, error) {
	mapData := map[string]interface{}{"username": username}
	if user, err := l.svcCtx.MemUserSev.GetByMap(l.ctx, mapData, ""); err == nil {
		//密码在客户端做md5
		//newPW := password + user.PasswordSlat
		//newPW2 := tool.Md5ByString(newPW)
		newPW2 := password
		//logx.Errorf(newPW + " | " + newPW2 + " | " + user.Password)
		if !(newPW2 == user.Password) {
			return 0, errors.Wrap(ErrUsernamePwdError, "密码匹配出错")
		}
		if user.Status != 1 {
			return 0, errors.Wrap(xerr.NewErrCode(xerr.Fail), "用户状态异常, user.Status")
		}
		return user.Id, nil
	} else {
		if err == gorm.ErrRecordNotFound {
			return 0, errors.Wrapf(ErrUserNoExistsError, " username=%s", username)
		} else {
			return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "loginByUsername db err ,  err:%v", err.Error())
		}
	}
}

func (l *LoginLogic) loginBySmallWx() error {
	return nil
}
