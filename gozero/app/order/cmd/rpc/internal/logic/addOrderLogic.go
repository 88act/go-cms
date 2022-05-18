package logic

import (
	"context"
	"encoding/json"
	"time"

	"go-cms/app/mqueue/cmd/job/jobtype"
	"go-cms/app/order/cmd/rpc/internal/svc"
	"go-cms/app/order/cmd/rpc/pb"
	"go-cms/app/order/model"
	"go-cms/app/usercenter/cmd/rpc/usercenter"
	"go-cms/common/kqueue"
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

	userInfo, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{
		Id: in.UserId,
	})
	if err != nil || userInfo == nil || userInfo.User.Status != 1 {
		return nil, errors.Wrapf(xerr.NewErrMsg("用户不存在或已失效"), "userid=%d ,order objId=%s, err: %v", in.UserId, in.ObjId, err)
	}

	//1 add order
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
		logx.WithContext(l.ctx).Errorf(" 延时关闭订单任务 json Marshal fail err :%+v , sn : %s", err, order.Sn)
	} else {
		info, err := l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(jobtype.DeferCloseHomestayOrder, payload), asynq.ProcessIn(CloseOrderTimeMinutes*time.Minute))
		if err != nil {
			logx.WithContext(l.ctx).Errorf(" 延时关闭订单任务 insert queue fail err :%+v , sn : %s", err, order.Sn)
		}
		logx.WithContext(l.ctx).Errorf(" 延时关闭订单任务成功 ,info =  :%+v ,", info)
	}

	//3、消息队列 发送邮件通知.
	emailObj := kqueue.EmailMessage{Email: "10512203@qq.com", TmpId: "1", Title: "标题"}
	err = l.sendEmail(emailObj)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("消息队列 发送 邮件通知 失败"), "userid=%d ,order objId=%s, err: %v", in.UserId, in.ObjId, err)
	}
	//4、消息队列 发送订单短信通知.
	smsObj := kqueue.SmsMessage{Mobile: "130130130130", TmpId: "1", Title: "sms标题"}
	err = l.sendSms(smsObj)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("消息队列 发送 短信通知 失败"), "userid=%d ,order objId=%s, err: %v", in.UserId, in.ObjId, err)
	}

	//4、通知订阅消息队列 notify  sub "payment-update-paystatus-topic"  services(order-mq ..), pub、sub use kq

	payObj := kqueue.ThirdPaymentUpdatePayStatusNotifyMessage{
		OrderSn:   order.Sn,
		PayStatus: 2,
	}

	if err := l.pubKqPaySuccess(payObj); err != nil {
		logx.WithContext(l.ctx).Errorf("通知订阅消息队 l.pubKqPaySuccess : %+v", err)
	}

	return &pb.AddOrderResp{Sn: order.Sn}, nil
}

func (l *AddOrderLogic) sendEmail(emailObj kqueue.EmailMessage) error {

	body, err := json.Marshal(emailObj)
	if err != nil {
		return errors.Wrapf(xerr.NewErrMsg(" pub 邮件消息出错"), "  error : %+v", err)
	}
	logx.Errorf("发布kafka 消息 sendEmail body =%v", body)
	// 发布kafka 消息
	return l.svcCtx.KqSendEmailClient.Push(string(body))
}

func (l *AddOrderLogic) sendSms(smsObj kqueue.SmsMessage) error {

	body, err := json.Marshal(smsObj)
	if err != nil {
		return errors.Wrapf(xerr.NewErrMsg(" pub 短信消息出错"), "  error : %+v", err)
	}
	logx.Errorf("发布kafka 消息 sendSms body =%v", body)
	// 发布kafka 消息
	return l.svcCtx.KqSendSmsClient.Push(string(body))
}

func (l *AddOrderLogic) pubKqPaySuccess(obj kqueue.ThirdPaymentUpdatePayStatusNotifyMessage) error {
	body, err := json.Marshal(obj)
	if err != nil {
		return errors.Wrapf(xerr.NewErrMsg(" 发布kafka 支付信息失败   error "), "k sk marshal error  , v : %+v", err.Error())
	}
	// 发布kafka 消息 支付成功
	return l.svcCtx.KqUpdatePayStatusClient.Push(string(body))
}
