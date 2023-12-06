package cms

import (
	"context"

	"go-cms/app/cms/cmd/api/internal/svc"
	"go-cms/app/cms/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DiscussListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDiscussListLogic(ctx context.Context, svcCtx *svc.ServiceContext) DiscussListLogic {
	return DiscussListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DiscussListLogic) DiscussList(req *types.PageInfoReq) (resp *types.DiscussListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
