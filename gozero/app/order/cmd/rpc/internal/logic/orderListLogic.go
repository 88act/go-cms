package logic

import (
	"context"

	"looklook/app/order/cmd/rpc/internal/svc"
	"looklook/app/order/cmd/rpc/order"
	"looklook/app/order/cmd/rpc/pb"
	"looklook/app/order/model"
	"looklook/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type OrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderListLogic {
	return &OrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 订单列表
func (l *OrderListLogic) OrderList(in *pb.OrderListReq) (*pb.OrderListResp, error) {
	// todo: add your logic here and delete this line
	seq := model.OrderOrderSearch{}
	seq.Page = int(in.Page)
	seq.PageSize = int(in.PageSize)
	seq.UserId = &in.UserId
	list, _, err := l.svcCtx.OrderOrderSev.GetList(l.ctx, seq, "")
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_RecordNotFound), "订单不存在", err)
	}
	var resp []*order.OrderInfo
	if len(list) > 0 {
		for _, actAct := range list {
			var obj order.OrderInfo
			_ = copier.Copy(&obj, actAct)
			//act.CreateTime = homestayOrder.CreateTime.Unix()
			//act.LiveStartDate = homestayOrder.LiveStartDate.Unix()
			//act.LiveEndDate = homestayOrder.LiveEndDate.Unix()
			resp = append(resp, &obj)
		}
	}
	return &pb.OrderListResp{List: resp}, nil
}
