package kq

import (
	"context"
	"encoding/json"
	"go-cms/app/order/cmd/rpc/internal/svc"
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

// 消耗/订阅 消息
func (l *PaymentUpdateStatusMq) Consume(_, val string) error {

	logx.WithContext(l.ctx).Error("order消耗/订阅  消息 ,开始执行.... ")
	var message kqueue.ThirdPaymentUpdatePayStatusNotifyMessage
	if err := json.Unmarshal([]byte(val), &message); err != nil {
		logx.WithContext(l.ctx).Error("第三方支付回调更改支付状态通知 PaymentUpdateStatusMq->Consume Unmarshal err : %v , val : %s", err, val)
		return err
	}

	if err := l.execService(message); err != nil {
		logx.WithContext(l.ctx).Error("第三方支付回调更改支付状态通知 PaymentUpdateStatusMq->execService  err : %v , val : %s , message:%+v", err, val, message)
		return err
	}

	return nil
}

func (l *PaymentUpdateStatusMq) execService(msg kqueue.ThirdPaymentUpdatePayStatusNotifyMessage) error {

	logx.WithContext(l.ctx).Errorf("order消息队列执行 =  message =%v", msg)
	orderPayState := l.getOrderPayStateByTrade(msg.PayStatus)
	if orderPayState != -99 {

		mapWhere := make(map[string]interface{})
		mapWhere["sn"] = msg.OrderSn
		//mapWhere["user_id"] = in.UserId
		mapData := make(map[string]interface{})
		mapData["status_pay"] = orderPayState
		//mapData["coupon_id"] = in.
		rows, err := l.svcCtx.OrderOrderSev.UpdateByMap(l.ctx, mapWhere, mapData)
		if err != nil || rows == 0 {
			return errors.Wrapf(xerr.NewErrMsg("更新订单支付状态失败 "), "更新订单支付状态失败 err : %v ,message:%+v", err, msg)
		}
		logx.WithContext(l.ctx).Errorf("消息队列 第三方支付回调更改支付状态通知 成功 sn=%s , message:%+v", msg.OrderSn, msg)
	}

	return nil
}

//获取订单支付状态 By第三方交易状态 .
func (l *PaymentUpdateStatusMq) getOrderPayStateByTrade(thirdPaymentPayStatus int32) int32 {

	switch thirdPaymentPayStatus {
	case paymentModel.PayStatus_PayOk:
		return model.OrderStatus_Payed
	case paymentModel.PayStatus_refunded:
		return model.OrderStatus_Refund
	default:
		return -99
	}

}
