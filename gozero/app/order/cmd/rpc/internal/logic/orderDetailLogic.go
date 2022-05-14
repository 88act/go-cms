package logic

import (
	"context"

	"go-cms/app/order/cmd/rpc/internal/svc"
	"go-cms/app/order/cmd/rpc/order"
	"go-cms/app/order/cmd/rpc/pb"
	"go-cms/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type OrderDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDetailLogic {
	return &OrderDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 订单详情
func (l *OrderDetailLogic) OrderDetail(in *pb.OrderDetailReq) (*pb.OrderDetailResp, error) {
	// todo: add your logic here and delete this line
	seqMap := make(map[string]interface{})
	seqMap["sn"] = in.Sn
	if in.UserId > 0 {
		seqMap["user_id"] = in.UserId
	}
	data, err := l.svcCtx.OrderOrderSev.GetByMap(l.ctx, seqMap, "")
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_RecordNotFound), "活动不存在 :%s,err:%v", in.Sn, err)
	}
	var resp order.OrderInfo
	_ = copier.Copy(&resp, data)
	//logx.Errorv(respAct)
	return &pb.OrderDetailResp{Info: &resp}, nil
}
