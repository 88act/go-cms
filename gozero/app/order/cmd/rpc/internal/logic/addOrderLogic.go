package logic

import (
	"context"

	"looklook/app/order/cmd/rpc/internal/svc"
	"looklook/app/order/cmd/rpc/pb"
	"looklook/app/order/model"
	"looklook/common/uniqueid"
	"looklook/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderLogic {
	return &AddOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 下订单
func (l *AddOrderLogic) AddOrder(in *pb.AddOrderReq) (*pb.AddOrderResp, error) {

	order := model.OrderOrder{}
	_ = copier.Copy(&order, in)
	order.Sn = uniqueid.GenSn(uniqueid.SN_PREFIX_HOMESTAY_ORDER)
	id, err := l.svcCtx.OrderOrderSev.Create(l.ctx, order)
	if err != nil || id == 0 {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_InsertErr), "err:%v", err)
	}
	return &pb.AddOrderResp{Sn: order.Sn}, nil
}
