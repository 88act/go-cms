package logic

import (
	"context"
	"encoding/json"
	"go-cms/app/mqueue/cmd/job/internal/svc"
	"go-cms/app/mqueue/cmd/job/jobtype"
	"go-cms/app/order/cmd/rpc/order"
	"go-cms/app/order/model"
	"go-cms/common/xerr"

	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var ErrCloseOrderFal = xerr.NewErrMsg("延时关闭订单失败ErrCloseOrderFal")

// CloseHomestayOrderHandler close no pay homestayOrder
type CloseHomestayOrderHandler struct {
	svcCtx *svc.ServiceContext
}

func NewCloseHomestayOrderHandler(svcCtx *svc.ServiceContext) *CloseHomestayOrderHandler {
	return &CloseHomestayOrderHandler{
		svcCtx: svcCtx,
	}
}

// 延时关闭没支付订单 : if return err != nil , asynq will retry
func (l *CloseHomestayOrderHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	logx.WithContext(ctx).Errorf("开始 延时关闭订单 asynq.Task =  :%+v ,", t)
	var p jobtype.DeferCloseHomestayOrderPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return errors.Wrapf(ErrCloseOrderFal, "closeOrderStateMqHandler payload err:%v, payLoad:%+v", err, t.Payload())
	}

	resp, err := l.svcCtx.OrderRpc.OrderDetail(ctx, &order.OrderDetailReq{
		Sn: p.Sn,
	})
	if err != nil || resp.Info == nil {
		return errors.Wrapf(ErrCloseOrderFal, "closeOrderStateMqHandler  订单不存在 err:%v, sn:%s ,order: %+v", err, p.Sn, resp.Info)
	}
	logx.WithContext(ctx).Errorf("延时关闭  StatusPay:%s ,sn: %s ,", resp.Info.StatusPay, p.Sn)

	if resp.Info.StatusPay == model.OrderStatus_WaitPay {
		logx.WithContext(ctx).Errorf("延时关闭 2  StatusPay  :%+v ", model.OrderStatus_WaitPay)
		_, err := l.svcCtx.OrderRpc.UpdateOrderStatus(ctx, &order.UpdateOrderStatusReq{
			Sn:          p.Sn,
			StatusOrder: model.OrderStatus_Cancel,
		})
		if err != nil {
			return errors.Wrapf(ErrCloseOrderFal, "closeOrderStateMqHandler close order fail  err:%v, sn:%s ", err, p.Sn)
		}
	}
	logx.WithContext(ctx).Errorf("延时关闭没支付订单 成功 ,asynq.Task =  :%+v ,", t)
	return nil
}
