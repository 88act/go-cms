package logic

import (
	"context"

	"looklook/app/order/cmd/rpc/internal/svc"
	"looklook/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePayStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePayStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePayStatusLogic {
	return &UpdatePayStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新支付状态
func (l *UpdatePayStatusLogic) UpdatePayStatus(in *pb.UpdatePayStatusReq) (*pb.UpdatePayStatusResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdatePayStatusResp{}, nil
}
