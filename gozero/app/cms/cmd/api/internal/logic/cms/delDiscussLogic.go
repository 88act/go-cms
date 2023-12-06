package cms

import (
	"context"

	"go-cms/app/cms/cmd/api/internal/svc"
	"go-cms/app/cms/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelDiscussLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelDiscussLogic(ctx context.Context, svcCtx *svc.ServiceContext) DelDiscussLogic {
	return DelDiscussLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelDiscussLogic) DelDiscuss(req *types.IdReq) (resp *types.OkResp, err error) {
	// todo: add your logic here and delete this line

	return
}
