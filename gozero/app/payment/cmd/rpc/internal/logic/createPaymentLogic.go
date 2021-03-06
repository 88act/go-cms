package logic

import (
	"context"
	"go-cms/app/payment/cmd/rpc/internal/svc"
	"go-cms/app/payment/cmd/rpc/pb"
	"go-cms/app/payment/model"
	"go-cms/common/uniqueid"
	"go-cms/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePaymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePaymentLogic {
	return &CreatePaymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreatePayment create wechat pay prepayorder.
func (l *CreatePaymentLogic) CreatePayment(in *pb.CreatePaymentReq) (*pb.CreatePaymentResp, error) {

	data := model.PayPayment{}
	data.Sn = uniqueid.GenSn(uniqueid.SN_PREFIX_THIRD_PAYMENT)
	data.UserId = &in.UserId
	data.PayMode = in.PayModel
	data.PayTotal = &in.PayTotal
	data.OrderSn = in.OrderSn
	data.ServiceType = model.ThirdPaymentServiceTypeHomestayOrder

	_, err := l.svcCtx.PayPaymentSev.Create(l.ctx, data)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "create wechat pay prepayorder db insert fail , err:%v ,data : %+v  ", err, data)
	}

	return &pb.CreatePaymentResp{
		Sn: data.Sn,
	}, nil
}
