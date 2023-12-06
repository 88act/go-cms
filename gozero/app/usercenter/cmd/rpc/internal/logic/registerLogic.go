package logic

import (
	"context"
	"go-cms/app/usercenter/cmd/rpc/internal/svc"
	"go-cms/app/usercenter/cmd/rpc/usercenter"
	"go-cms/app/usercenter/model"
	"go-cms/common/tool"
	"go-cms/common/utils"
	"go-cms/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

var ErrUserAlreadyRegisterError = xerr.NewErrMsg("用户已存在")

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *usercenter.RegisterReq) (*usercenter.RegisterResp, error) {
	var userId int64
	var err error
	switch in.LoginType { //注册登录类型  1 用户名密码 2 邮箱密码 3 手机/验证码  4 微信 5 QQ 6 github 7 gmail   8 piko
	case 1:
		userId, err = l.regByUsername(in)
	case 2:
		userId, err = l.regByMobile(in)
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
	return &usercenter.RegisterResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}

func (l *RegisterLogic) regByMobile(in *usercenter.RegisterReq) (int64, error) {
	mapData := map[string]interface{}{"mobile": in.Mobile}
	if user, err := l.svcCtx.MemUserSev.GetByMap(l.ctx, mapData, ""); err == nil || err == gorm.ErrRecordNotFound {
		if user.Id > 0 {
			return 0, errors.Wrapf(ErrUserAlreadyRegisterError, "Register user exists mobile:%s,err:%v", in.Mobile, err)
		} else {
			var data = model.MemUser{}
			_ = copier.Copy(&data, in)
			//data.Username = in.Mobile
			data.Password = in.Password
			// 密码在客户端做了md5 处理
			// data.PasswordSlat = tool.Krand(8, tool.KC_RAND_KIND_ALL)
			// pwd := data.Password + data.PasswordSlat
			// data.Password = tool.Md5ByString(pwd)
			if utils.IsEmpty(data.Nickname) {
				data.Nickname = tool.Krand(8, tool.KC_RAND_KIND_ALL)
			}
			data.Status = 1
			if userId, err := l.svcCtx.MemUserSev.Create(l.ctx, data); err == nil {
				return userId, nil
			} else {
				return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "注册用户出错，db err:%v,user:%+v", err, data)
			}
		}
	} else {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "db err ,  err:%v", err.Error())
	}
}

func (l *RegisterLogic) regByUsername(in *usercenter.RegisterReq) (int64, error) {
	mapData := map[string]interface{}{"username": in.Username}
	if user, err := l.svcCtx.MemUserSev.GetByMap(l.ctx, mapData, ""); err == nil || err == gorm.ErrRecordNotFound {
		if user.Id > 0 {
			return 0, errors.Wrapf(ErrUserAlreadyRegisterError, "Register user exists mobile:%s,err:%v", in.Mobile, err)
		} else {
			var data = model.MemUser{}
			_ = copier.Copy(&data, in)
			//data.Username = in.Mobile
			data.Password = in.Password
			data.PasswordSlat = tool.Krand(8, tool.KC_RAND_KIND_ALL)
			pwd := data.Password + data.PasswordSlat
			data.Password = tool.Md5ByString(pwd)
			if utils.IsEmpty(data.Nickname) {
				data.Nickname = tool.Krand(8, tool.KC_RAND_KIND_ALL)
			}
			data.Status = 1
			if userId, err := l.svcCtx.MemUserSev.Create(l.ctx, data); err == nil {
				return userId, nil
			} else {
				return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "注册用户出错，db err:%v,user:%+v", err, data)
			}
		}
	} else {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "db err ,  err:%v", err.Error())
	}
}

//session := db.WithContext(ctx)

// if err := l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
// 	user := new(model.User)
// 	user.Mobile = in.Mobile
// 	if len(user.Nickname) == 0 {
// 		user.Nickname = tool.Krand(8, tool.KC_RAND_KIND_ALL)
// 	}
// 	if len(in.Password) > 0 {
// 		user.Password = tool.Md5ByString(in.Password)
// 	}
// 	logx.Errorf("Register == 444444 \r\n")
// 	//insertResult, err := l.svcCtx.UserModel.Insert(ctx, session, user)
// 	lastId, err := l.svcCtx.UserModel.AddOne(ctx, user)
// 	if err != nil {
// 		return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Register db user Insert err:%v,user:%+v", err, user)
// 	}
// 	//lastId, err := insertResult.LastInsertId()
// 	logx.Errorf("Register == 555555  ")
// 	logx.Errorf(string(lastId))
// 	if err != nil {
// 		return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Register db user insertResult.LastInsertId err:%v,user:%+v", err, user)
// 	}

// 	logx.Errorf("Register == 66666 \r\n")

// 	userId = lastId

// 	// userAuth := new(model.UserAuth)
// 	// userAuth.UserId = lastId
// 	// userAuth.AuthKey = in.AuthKey
// 	// userAuth.AuthType = in.AuthType
// 	// if _, err := l.svcCtx.UserAuthModel.Insert(ctx, session, userAuth); err != nil {
// 	// 	return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Register db user_auth Insert err:%v,userAuth:%v", err, userAuth)
// 	// }
// 	return nil
// }); err != nil {
// 	return nil, err
// }
