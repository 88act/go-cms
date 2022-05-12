package logic

import (
	"context"

	"go-cms/app/order/cmd/rpc/internal/svc"
	"go-cms/app/order/cmd/rpc/pb"
	"go-cms/app/order/model"
	"go-cms/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderLogic {
	return &UpdateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新订单状态
func (l *UpdateOrderLogic) UpdateOrder(in *pb.UpdateOrderReq) (*pb.UpdateOrderResp, error) {
	order := model.OrderOrder{}
	_ = copier.Copy(&order, in)
	//order.Sn = uniqueid.GenSn(uniqueid.SN_PREFIX_HOMESTAY_ORDER)
	mapWhere := make(map[string]interface{})
	mapWhere["sn"] = in.Sn
	mapWhere["user_id"] = in.UserId

	mapData := make(map[string]interface{})
	mapData["remark"] = in.Remark
	mapData["order_code"] = in.OrderCode
	mapData["status_order"] = in.StatusOrder
	mapData["status"] = in.Status
	//mapData["coupon_id"] = in.
	rows, err := l.svcCtx.OrderOrderSev.UpdateByMap(l.ctx, mapWhere, mapData)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_UPDATE_AFFECTED_ZERO_ERROR), "err:%v", err)
	}
	return &pb.UpdateOrderResp{Rows: rows}, nil
}
