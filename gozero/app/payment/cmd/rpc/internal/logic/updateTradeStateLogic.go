package logic

import (
	"context"
	"encoding/json"
	"go-cms/common/kqueue"
	"time"

	"go-cms/app/payment/cmd/rpc/internal/svc"
	"go-cms/app/payment/cmd/rpc/pb"
	"go-cms/app/payment/model"
	"go-cms/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UpdateTradeStateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTradeStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTradeStateLogic {
	return &UpdateTradeStateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTradeStateLogic) UpdateTradeState(in *pb.UpdateTradeStateReq) (*pb.UpdateTradeStateResp, error) {

	seqMap := make(map[string]interface{})
	seqMap["sn"] = in.Sn
	thirdPayment, err := l.svcCtx.PayPaymentSev.GetByMap(l.ctx, seqMap, "")
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_RecordNotFound), "支付信息不存在 :%s,err:%v", in.Sn, err)
	}

	//2、Judgment Status
	if in.PayStatus == model.PayStatus_PayOk || in.PayStatus == model.PayStatus_Fail {
		//Want to modify as payment success, failure scenarios
		if *thirdPayment.PayStatus != int(model.PayStatus_WaitPay) {
			return &pb.UpdateTradeStateResp{}, nil
		}

	} else if in.PayStatus == model.PayStatus_refunded {
		//Want to change to refund success scenario

		if *thirdPayment.PayStatus != int(model.PayStatus_PayOk) {
			return nil, errors.Wrapf(xerr.NewErrMsg("Only orders with successful payment can be refunded"), "Only orders with successful payment can be refunded in : %+v", in)
		}
	} else {
		return nil, errors.Wrapf(xerr.NewErrMsg("This status is not currently supported"), "Modify payment flow status is not supported  in : %+v", in)
	}

	//3、update .
	thirdPayment.TradeState = in.TradeState
	thirdPayment.TransactionId = in.TransactionId
	thirdPayment.TradeType = in.TradeType
	thirdPayment.TradeStateDesc = in.TradeStateDesc
	*thirdPayment.PayStatus = int(in.PayStatus)
	*thirdPayment.PayTime = time.Unix(in.PayTime, 0)
	if err := l.svcCtx.PayPaymentSev.Update(l.ctx, *thirdPayment); err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_UpdateErr), " UpdateTradeState UpdateWithVersion db  err:%v ,thirdPayment : %+v , in : %+v", err, thirdPayment, in)
	}

	//4、通知订阅消息队列 notify  sub "payment-update-paystatus-topic"  services(order-mq ..), pub、sub use kq
	if err := l.pubKqPaySuccess(in.Sn, in.PayStatus); err != nil {
		logx.WithContext(l.ctx).Errorf("l.pubKqPaySuccess : %+v", err)
	}

	return &pb.UpdateTradeStateResp{}, nil
}

func (l *UpdateTradeStateLogic) pubKqPaySuccess(orderSn string, payStatus int32) error {

	m := kqueue.ThirdPaymentUpdatePayStatusNotifyMessage{
		OrderSn:   orderSn,
		PayStatus: payStatus,
	}

	body, err := json.Marshal(m)
	if err != nil {
		return errors.Wrapf(xerr.NewErrMsg(" 发布kafka 支付信息失败 kq UpdateTradeStateLogic pushKqPaySuccess task marshal error "), "kq UpdateTradeStateLogic pushKqPaySuccess task marshal error  , v : %+v", m)
	}
	// 发布kafka 消息 支付成功
	return l.svcCtx.KqueuePaymentUpdatePayStatusClient.Push(string(body))
}
