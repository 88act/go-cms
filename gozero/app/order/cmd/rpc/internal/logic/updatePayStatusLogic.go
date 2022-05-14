package logic

import (
	"context"

	"go-cms/app/order/cmd/rpc/internal/svc"
	"go-cms/app/order/cmd/rpc/pb"
	"go-cms/common/xerr"

	"github.com/pkg/errors"
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
	mapWhere := make(map[string]interface{})
	mapWhere["sn"] = in.Sn
	//mapWhere["user_id"] = in.UserId
	mapData := make(map[string]interface{})
	mapData["status_pay"] = in.StatusPay
	//mapData["coupon_id"] = in.
	rows, err := l.svcCtx.OrderOrderSev.UpdateByMap(l.ctx, mapWhere, mapData)
	if err != nil || rows == 0 {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_UPDATE_AFFECTED_ZERO_ERROR), "err:%v", err)
	}
	return &pb.UpdatePayStatusResp{Rows: rows}, nil
}
