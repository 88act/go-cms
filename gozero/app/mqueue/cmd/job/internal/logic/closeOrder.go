package logic

import (
	"context"
	"encoding/json"
	"looklook/app/mqueue/cmd/job/internal/svc"
	"looklook/app/mqueue/cmd/job/jobtype"
	"looklook/app/order/cmd/rpc/order"
	"looklook/app/order/model"
	"looklook/common/xerr"

	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
)

var ErrCloseOrderFal = xerr.NewErrMsg("close order fail")

// CloseHomestayOrderHandler close no pay homestayOrder
type CloseHomestayOrderHandler struct {
	svcCtx *svc.ServiceContext
}

func NewCloseHomestayOrderHandler(svcCtx *svc.ServiceContext) *CloseHomestayOrderHandler {
	return &CloseHomestayOrderHandler{
		svcCtx: svcCtx,
	}
}

// defer  close no pay homestayOrder  : if return err != nil , asynq will retry
func (l *CloseHomestayOrderHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {

	var p jobtype.DeferCloseHomestayOrderPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return errors.Wrapf(ErrCloseOrderFal, "closeHomestayOrderStateMqHandler payload err:%v, payLoad:%+v", err, t.Payload())
	}

	resp, err := l.svcCtx.OrderRpc.OrderDetail(ctx, &order.OrderDetailReq{
		Sn: p.Sn,
	})
	if err != nil || resp.Info == nil {
		return errors.Wrapf(ErrCloseOrderFal, "closeHomestayOrderStateMqHandler  get order fail or order no exists err:%v, sn:%s ,HomestayOrder : %+v", err, p.Sn, resp.Info)
	}

	if resp.Info.StatusPay == model.OrderStatus_WaitPay {
		_, err := l.svcCtx.OrderRpc.UpdateOrderStatus(ctx, &order.UpdateOrderStatusReq{
			Sn:          p.Sn,
			StatusOrder: model.OrderStatus_Cancel,
		})
		if err != nil {
			return errors.Wrapf(ErrCloseOrderFal, "CloseHomestayOrderHandler close order fail  err:%v, sn:%s ", err, p.Sn)
		}
	}

	return nil
}
