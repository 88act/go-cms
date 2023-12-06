package user

import (
	"context"

	"go-cms/app/usercenter/cmd/api/internal/svc"
	"go-cms/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TokenVerifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTokenVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) TokenVerifyLogic {
	return TokenVerifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TokenVerifyLogic) TokenVerify() (resp *types.OkResp, err error) {
	// todo: add your logic here and delete this line

	return
}
