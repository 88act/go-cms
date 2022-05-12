package logic

import (
	"context"

	"go-cms/app/order/cmd/rpc/internal/svc"
	"go-cms/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderStatusLogic {
	return &UpdateOrderStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新订单状态
func (l *UpdateOrderStatusLogic) UpdateOrderStatus(in *pb.UpdateOrderStatusReq) (*pb.UpdateOrderStatusResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateOrderStatusResp{}, nil
}
