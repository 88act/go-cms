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

type OrderDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderDetailLogic {
	return OrderDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderDetailLogic) OrderDetail(req *types.OrderDetailReq) (resp *types.OrderDetailResp, err error) {

	userId := ctxdata.GetUidFromCtx(l.ctx)
	rpcResp, err := l.svcCtx.OrderRpc.OrderDetail(l.ctx, &order.OrderDetailReq{
		Sn:     req.Sn,
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	var orderInfo types.OrderInfo
	_ = copier.Copy(&orderInfo, rpcResp.Info)

	return &types.OrderDetailResp{
		Info: orderInfo,
	}, nil
}
