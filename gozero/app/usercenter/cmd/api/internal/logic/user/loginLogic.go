package user

import (
	"context"
	"errors"

	"go-cms/app/usercenter/cmd/api/internal/svc"
	"go-cms/app/usercenter/cmd/api/internal/types"
	"go-cms/common/tool"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {

	// customer, err := l.svcCtx.JqCustomerSev.GetByGuid(l.ctx, req.CuGuid, "")
	// if err != nil {
	// 	return nil, errors.New("不存在的客户")
	// }
	// if customer.Status != 1 {
	// 	return nil, errors.New("客户已禁用")
	// }

	if user, err := l.svcCtx.MemUserSev.GetByUsername(l.ctx, 0, req.Username); err == nil {
		newPW := req.Password + user.PasswordSlat
		newPW2 := tool.Md5ByString(newPW)

		//logx.Errorf(newPW + " | " + newPW2 + " | " + user.Password)
		if newPW2 != user.Password {
			return nil, errors.New("密码出错")
		}
		if user.Status != 1 {
			return nil, errors.New("用户状态异常")
		}
		myResp, err := GenerateToken(l.svcCtx.Config.JwtAuth.AccessExpire, l.svcCtx.Config.JwtAuth.AccessSecret, user)
		if err != nil {
			return nil, errors.New("生成token失败")
		}
		return myResp, nil
	} else {
		return nil, errors.New("用户不存在")
	}

	// if rpcResp, err := l.svcCtx.UsercenterRpc.Login(l.ctx, &rpcReq); err == nil {

	// 	var myResp types.LoginResp
	// 	_ = copier.Copy(&myResp, rpcResp)
	// 	logx.Errorf("========3====== ")
	// 	return &myResp, nil
	// } else {
	// 	logx.Errorf("========4======%v", err.Error())
	// 	return nil, errors.Wrapf(err, "req: %+v", req)
	// }
}
