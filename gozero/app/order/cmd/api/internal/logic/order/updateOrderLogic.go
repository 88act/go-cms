package order

import (
	"context"

	"looklook/app/order/cmd/api/internal/svc"
	"looklook/app/order/cmd/api/internal/types"
	"looklook/app/order/cmd/rpc/order"
	"looklook/common/ctxdata"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateOrderLogic {
	return UpdateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateOrderLogic) UpdateOrder(req *types.UpdateOrderReq) (resp *types.UpdateOrderResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	updateOrderReq := order.UpdateOrderReq{}
	_ = copier.Copy(&updateOrderReq, req)
	updateOrderReq.UserId = userId
	rpcResp, err := l.svcCtx.OrderRpc.UpdateOrder(l.ctx, &updateOrderReq)
	if err != nil {
		return nil, err
	}
	return &types.UpdateOrderResp{
		Rows: rpcResp.Rows,
	}, nil
}
