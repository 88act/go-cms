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
type ImageZipMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImageZipMq(ctx context.Context, svcCtx *svc.ServiceContext) *ImageZipMq {
	return &ImageZipMq{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 消耗/订阅 消息
func (l *ImageZipMq) Consume(_, val string) error {

	logx.WithContext(l.ctx).Error("消耗/订阅 消息 ,开始执行....ImageZipMq ")
	var message kqueue.ImageZipMessage
	if err := json.Unmarshal([]byte(val), &message); err != nil {
		logx.WithContext(l.ctx).Error("第三方支付回调更改支付状态通知 ImageZipMq->Consume Unmarshal err : %v , val : %s", err, val)
		return err
	}

	if err := l.execService(message); err != nil {
		logx.WithContext(l.ctx).Error("第三方支付回调更改支付状态通知 ImageZipMq->execService  err : %v , val : %s , message:%+v", err, val, message)
		return err
	}

	return nil
}

func (l *ImageZipMq) execService(message kqueue.ImageZipMessage) error {

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
	// 	logx.WithContext(l.ctx).Info("第三方支付回调更改支付状态通知 成功 sn=%s , message:%+v", message.OrderSn, message)
	// }
	logx.Errorf("图片压缩 .... ")
	return nil
}
