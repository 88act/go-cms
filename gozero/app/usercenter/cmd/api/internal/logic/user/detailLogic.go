package user

import (
	"context"

	"go-cms/app/usercenter/cmd/api/internal/svc"
	"go-cms/app/usercenter/cmd/api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) DetailLogic {
	return DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(userId int64) (resp *types.UserInfoResp, err error) {

	if memUser, err := l.svcCtx.MemUserSev.Get(l.ctx, userId, ""); err == nil {
		var user types.MemUser
		_ = copier.Copy(&user, memUser)
		return &types.UserInfoResp{UserInfo: user}, nil
	} else {
		return nil, err
	}
	// rpcReq := usercenter.GetUserInfoReq{Id: userId}
	// if rpcResp, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &rpcReq); err == nil {
	// 	//logx.Error(rpcResp)
	// 	var user types.MemUser
	// 	_ = copier.Copy(&user, rpcResp.User)
	// 	return &types.UserInfoResp{UserInfo: user}, nil
	// } else {
	// 	return nil, errors.Wrapf(err, "req: %+v", userId)
	// }
}
