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

type OrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderListLogic {
	return OrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderListLogic) OrderList(req *types.OrderListReq) (resp *types.OrderListResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	rpcResp, err := l.svcCtx.OrderRpc.OrderList(l.ctx, &order.OrderListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		UserId:   userId,
		Search:   "",
	})
	if err != nil {
		return nil, err
	}
	var list []types.OrderInfo
	if len(rpcResp.List) > 0 {
		for _, item := range rpcResp.List {
			var obj types.OrderInfo
			_ = copier.Copy(&obj, item)
			list = append(list, obj)
		}
	}
	return &types.OrderListResp{
		List: list,
	}, nil
}
