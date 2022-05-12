package order

import (
	"context"

	"go-cms/app/order/cmd/api/internal/svc"
	"go-cms/app/order/cmd/api/internal/types"
	"go-cms/app/order/cmd/rpc/order"
	"go-cms/common/ctxdata"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddOrderLogic {
	return AddOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddOrderLogic) AddOrder(req *types.AddOrderReq) (resp *types.AddOrderResp, err error) {
	// todo: add your logic here and delete this line
	userId := ctxdata.GetUidFromCtx(l.ctx)
	addOrderReq := order.AddOrderReq{}
	_ = copier.Copy(&addOrderReq, req)
	addOrderReq.UserId = userId
	rpcResp, err := l.svcCtx.OrderRpc.AddOrder(l.ctx, &addOrderReq)
	if err != nil {
		return nil, err
	}
	return &types.AddOrderResp{
		Sn: rpcResp.Sn,
	}, nil
}
