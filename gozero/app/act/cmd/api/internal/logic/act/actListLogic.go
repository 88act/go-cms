package act

import (
	"context"

	"looklook/app/act/cmd/api/internal/svc"
	"looklook/app/act/cmd/api/internal/types"
	"looklook/app/act/cmd/rpc/act"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type ActListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ActListLogic {
	return ActListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActListLogic) ActList(req types.ActListReq) (resp *types.ActListResp, err error) {
	///userId := ctxdata.GetUidFromCtx(l.ctx)
	rpcResp, err := l.svcCtx.ActRpc.ActList(l.ctx, &act.ActListReq{
		Page: req.Page, PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	var actList []types.ActAct
	if len(rpcResp.List) > 0 {
		for _, item := range rpcResp.List {
			var obj types.ActAct
			_ = copier.Copy(&obj, item)
			actList = append(actList, obj)
		}
	}
	return &types.ActListResp{
		List: actList,
	}, nil
}
