package logic

import (
	"context"

	"go-cms/app/order/cmd/rpc/internal/svc"
	"go-cms/app/order/cmd/rpc/pb"
	"go-cms/common/xerr"

	"github.com/pkg/errors"
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

	mapWhere := make(map[string]interface{})
	mapWhere["sn"] = in.Sn
	//mapWhere["user_id"] = in.UserId
	mapData := make(map[string]interface{})
	mapData["status_order"] = in.StatusOrder
	//mapData["coupon_id"] = in.
	rows, err := l.svcCtx.OrderOrderSev.UpdateByMap(l.ctx, mapWhere, mapData)
	if err != nil || rows == 0 {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_UPDATE_AFFECTED_ZERO_ERROR), "err:%v", err)
	}
	return &pb.UpdateOrderStatusResp{Rows: rows}, nil

}
