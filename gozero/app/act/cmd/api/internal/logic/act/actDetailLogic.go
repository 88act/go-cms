package act

import (
	"context"

	"looklook/app/act/cmd/api/internal/svc"
	"looklook/app/act/cmd/api/internal/types"
	"looklook/app/act/cmd/rpc/act"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type ActDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) ActDetailLogic {
	return ActDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActDetailLogic) ActDetail(req types.ActDetailReq) (resp *types.ActDetailResp, err error) {
	///userId := ctxdata.GetUidFromCtx(l.ctx)
	rpcResp, err := l.svcCtx.ActRpc.ActDetail(l.ctx, &act.ActDetailReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	var actAct types.ActAct
	_ = copier.Copy(&actAct, rpcResp.Act)

	return &types.ActDetailResp{
		Act: actAct,
	}, nil
}
