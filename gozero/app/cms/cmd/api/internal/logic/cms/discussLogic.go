package cms

import (
	"context"

	"go-cms/app/cms/cmd/api/internal/svc"
	"go-cms/app/cms/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DiscussLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDiscussLogic(ctx context.Context, svcCtx *svc.ServiceContext) DiscussLogic {
	return DiscussLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DiscussLogic) Discuss(req *types.DiscussReq) (resp *types.OkResp, err error) {
	// todo: add your logic here and delete this line

	return
}
