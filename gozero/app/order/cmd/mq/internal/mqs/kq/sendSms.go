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
type SendSmsMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendSmsMq(ctx context.Context, svcCtx *svc.ServiceContext) *SendSmsMq {
	return &SendSmsMq{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 消耗/订阅 消息
func (l *SendSmsMq) Consume(_, val string) error {

	logx.WithContext(l.ctx).Error("消耗/订阅 消息 ,开始执行.... ")
	var message kqueue.ThirdPaymentUpdatePayStatusNotifyMessage
	if err := json.Unmarshal([]byte(val), &message); err != nil {
		logx.WithContext(l.ctx).Error("第三方支付回调更改支付状态通知 SendSmsMq->Consume Unmarshal err : %v , val : %s", err, val)
		return err
	}

	if err := l.execService(message); err != nil {
		logx.WithContext(l.ctx).Error("第三方支付回调更改支付状态通知 SendSmsMq->execService  err : %v , val : %s , message:%+v", err, val, message)
		return err
	}

	return nil
}

func (l *SendSmsMq) execService(message kqueue.ThirdPaymentUpdatePayStatusNotifyMessage) error {

	orderPayState := l.getOrderPayStateByTrade(message.PayStatus)
	if orderPayState != -99 {
		//update homestay order state
		//_, err := l.svcCtx.OrderRpc.UpdateHomestayOrderTradeState(l.ctx, &order.UpdateHomestayOrderTradeStateReq{
		_, err := l.svcCtx.OrderRpc.UpdatePayStatus(l.ctx, &order.UpdatePayStatusReq{
			Sn:        message.OrderSn,
			StatusPay: orderPayState,
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrMsg("更新订单支付状态失败 "), "更新订单支付状态失败 err : %v ,message:%+v", err, message)
		}
	}

	return nil
}

//获取订单支付状态 By第三方交易状态 .
func (l *SendSmsMq) getOrderPayStateByTrade(thirdPaymentPayStatus int32) int32 {

	switch thirdPaymentPayStatus {
	case paymentModel.PayStatus_PayOk:
		return model.OrderStatus_Payed
	case paymentModel.PayStatus_refunded:
		return model.OrderStatus_Refund
	default:
		return -99
	}

}
