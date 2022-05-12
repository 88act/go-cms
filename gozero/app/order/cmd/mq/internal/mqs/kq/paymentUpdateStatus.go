package kq

import (
	"context"
	"encoding/json"
	"go-cms/app/order/cmd/mq/internal/svc"
	"go-cms/app/order/cmd/rpc/order"
	"go-cms/app/order/model"
	paymentModel "go-cms/app/payment/model"
	"go-cms/common/kqueue"
	"go-cms/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

/**
Listening to the payment flow status change notification message queue
*/
type PaymentUpdateStatusMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentUpdateStatusMq(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentUpdateStatusMq {
	return &PaymentUpdateStatusMq{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaymentUpdateStatusMq) Consume(_, val string) error {

	var message kqueue.ThirdPaymentUpdatePayStatusNotifyMessage
	if err := json.Unmarshal([]byte(val), &message); err != nil {
		logx.WithContext(l.ctx).Error("PaymentUpdateStatusMq->Consume Unmarshal err : %v , val : %s", err, val)
		return err
	}

	if err := l.execService(message); err != nil {
		logx.WithContext(l.ctx).Error("PaymentUpdateStatusMq->execService  err : %v , val : %s , message:%+v", err, val, message)
		return err
	}

	return nil
}

func (l *PaymentUpdateStatusMq) execService(message kqueue.ThirdPaymentUpdatePayStatusNotifyMessage) error {

	orderTradeState := l.getOrderTradeStateByPaymentTradeState(message.PayStatus)
	if orderTradeState != -99 {
		//update homestay order state
		//_, err := l.svcCtx.OrderRpc.UpdateHomestayOrderTradeState(l.ctx, &order.UpdateHomestayOrderTradeStateReq{
		_, err := l.svcCtx.OrderRpc.UpdatePayStatus(l.ctx, &order.UpdatePayStatusReq{
			Sn:        message.OrderSn,
			StatusPay: 1,
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrMsg("update homestay order state fail"), "update homestay order state fail err : %v ,message:%+v", err, message)
		}
	}

	return nil
}

//Get order status based on payment status.
func (l *PaymentUpdateStatusMq) getOrderTradeStateByPaymentTradeState(thirdPaymentPayStatus int32) int32 {

	switch thirdPaymentPayStatus {
	case paymentModel.PayStatus_PayOk:
		return model.OrderStatus_WaitUse
	case paymentModel.PayStatus_refunded:
		return model.OrderStatus_Refund
	default:
		return -99
	}

}
