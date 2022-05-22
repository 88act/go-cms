package kq

import (
	"context"
	"encoding/json"
	"go-cms/app/basic/cmd/rpc/internal/svc"
	"go-cms/common/kqueue"

	"github.com/zeromicro/go-zero/core/logx"
)

/**
Listening to the payment flow status change notification message queue
*/
type SendEmailMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendEmailMq(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailMq {
	return &SendEmailMq{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 消耗/订阅 消息
func (l *SendEmailMq) Consume(_, val string) error {

	logx.WithContext(l.ctx).Error("消耗/订阅 消息 ,开始执行.... SendEmailMq ")
	var message kqueue.EmailMessage
	if err := json.Unmarshal([]byte(val), &message); err != nil {
		logx.WithContext(l.ctx).Error("第三方支付回调更改支付状态通知 SendEmailMq->Consume Unmarshal err : %v , val : %s", err, val)
		return err
	}

	if err := l.execService(message); err != nil {
		logx.WithContext(l.ctx).Error("第三方支付回调更改支付状态通知 SendEmailMq->execService  err : %v , val : %s , message:%+v", err, val, message)
		return err
	}

	return nil
}

func (l *SendEmailMq) execService(message kqueue.EmailMessage) error {

	// orderPayState := l.getOrderPayStateByTrade(message.PayStatus)
	// if orderPayState != -99 {
	// 	//update homestay order state
	// 	//_, err := l.svcCtx.OrderRpc.UpdateHomestayOrderTradeState(l.ctx, &order.UpdateHomestayOrderTradeStateReq{
	// 	_, err := l.svcCtx.OrderRpc.UpdatePayStatus(l.ctx, &order.UpdatePayStatusReq{
	// 		Sn:        message.OrderSn,
	// 		StatusPay: orderPayState,
	// 	})
	// 	if err != nil {
	// 		return errors.Wrapf(xerr.NewErrMsg("更新订单支付状态失败 "), "更新订单支付状态失败 err : %v ,message:%+v", err, message)
	// 	}
	// }

	logx.Errorf("收到订阅 send email")

	return nil
}
