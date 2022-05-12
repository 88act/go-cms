package logic

import (
	"context"

	"go-cms/app/payment/cmd/rpc/internal/svc"
	"go-cms/app/payment/cmd/rpc/payment"
	"go-cms/app/payment/cmd/rpc/pb"
	"go-cms/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type GetPaymentBySnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPaymentBySnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPaymentBySnLogic {
	return &GetPaymentBySnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPaymentBySnLogic) GetPaymentBySn(in *pb.GetPaymentBySnReq) (*pb.GetPaymentBySnResp, error) {

	seqMap := make(map[string]interface{})
	seqMap["sn"] = in.Sn
	data, err := l.svcCtx.PayPaymentSev.GetByMap(l.ctx, seqMap, "")
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_RecordNotFound), "支付信息不存在 :%s,err:%v", in.Sn, err)
	}
	var resp payment.PaymentDetail
	if data != nil {
		_ = copier.Copy(&resp, data)
		resp.CreatedAt = data.CreatedAt.Unix()
		resp.UpdatedAt = data.UpdatedAt.Unix()
		resp.PayTime = data.PayTime.Unix()
	}
	return &pb.GetPaymentBySnResp{PaymentDetail: &resp}, nil
}
