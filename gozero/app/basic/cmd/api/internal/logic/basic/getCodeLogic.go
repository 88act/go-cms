package basic

import (
	"context"

	"go-cms/app/basic/cmd/api/internal/svc"
	"go-cms/app/basic/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetCodeLogic {
	return GetCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCodeLogic) GetCode(req *types.GetCodeReq) (resp *types.GetCodeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
