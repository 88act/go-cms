package logic

import (
	"context"
	"looklook/app/payment/cmd/rpc/internal/svc"
	"looklook/app/payment/cmd/rpc/payment"
	"looklook/app/payment/cmd/rpc/pb"
	"looklook/app/payment/model"
	"looklook/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type GetPaymentSuccessRefundByOrderSnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPaymentSuccessRefundByOrderSnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPaymentSuccessRefundByOrderSnLogic {
	return &GetPaymentSuccessRefundByOrderSnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPaymentSuccessRefundByOrderSnLogic) GetPaymentSuccessRefundByOrderSn(in *pb.GetPaymentSuccessRefundByOrderSnReq) (*pb.GetPaymentSuccessRefundByOrderSnResp, error) {

	seqMap := make(map[string]interface{})
	seqMap["trade_state"] = model.PayStatus_PayOk
	seqMap["order_sn"] = in.OrderSn
	seqMap["trade_state"] = model.PayStatus_refunded
	data, err := l.svcCtx.PayPaymentSev.GetByMap(l.ctx, seqMap, "")
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_RecordNotFound), "支付信息不存在  err : %v , orderSn:%s", err, in.OrderSn)
	}
	var resp payment.PaymentDetail
	if data != nil {
		_ = copier.Copy(&resp, data)
		resp.CreatedAt = data.CreatedAt.Unix()
		resp.UpdatedAt = data.UpdatedAt.Unix()
		resp.PayTime = data.PayTime.Unix()
	}
	return &pb.GetPaymentSuccessRefundByOrderSnResp{PaymentDetail: &resp}, nil
}
