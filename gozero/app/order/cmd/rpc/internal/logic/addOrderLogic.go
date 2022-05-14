package logic

import (
	"context"
	"encoding/json"
	"time"

	"go-cms/app/mqueue/cmd/job/jobtype"
	"go-cms/app/order/cmd/rpc/internal/svc"
	"go-cms/app/order/cmd/rpc/pb"
	"go-cms/app/order/model"
	"go-cms/common/uniqueid"
	"go-cms/common/xerr"

	"github.com/hibiken/asynq"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

const CloseOrderTimeMinutes = 1 //延时关闭未支付订单/分钟
type AddOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderLogic {
	return &AddOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 下订单
func (l *AddOrderLogic) AddOrder(in *pb.AddOrderReq) (*pb.AddOrderResp, error) {

	order := model.OrderOrder{}
	_ = copier.Copy(&order, in)
	order.Sn = uniqueid.GenSn(uniqueid.SN_PREFIX_HOMESTAY_ORDER)
	id, err := l.svcCtx.OrderOrderSev.Create(l.ctx, order)
	if err != nil || id == 0 {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_InsertErr), "Order Database Exception order : %+v , err: %v", order, err)
	}

	//2、延时关闭 of order tasks.
	payload, err := json.Marshal(jobtype.DeferCloseHomestayOrderPayload{Sn: order.Sn})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("创建延时关闭订单 task json Marshal fail err :%+v , sn : %s", err, order.Sn)
	} else {
		info, err := l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(jobtype.DeferCloseHomestayOrder, payload), asynq.ProcessIn(CloseOrderTimeMinutes*time.Minute))
		if err != nil {
			logx.WithContext(l.ctx).Errorf("创建延时关闭订单 task insert queue fail err :%+v , sn : %s", err, order.Sn)
		}
		logx.WithContext(l.ctx).Errorf("创建延时关闭订单 成功 ,info =  :%+v ,", info)
	}

	return &pb.AddOrderResp{Sn: order.Sn}, nil
}
