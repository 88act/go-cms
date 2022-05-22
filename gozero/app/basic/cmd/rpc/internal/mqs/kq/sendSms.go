package kq

import (
	"context"
	"encoding/json"
	"go-cms/app/basic/cmd/rpc/internal/svc"
	"go-cms/app/basic/model"
	"go-cms/common/kqueue"

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

	logx.WithContext(l.ctx).Error("消耗/订阅 消息 ,开始执行....SendSmsMq  ")
	var message kqueue.SmsMessage
	if err := json.Unmarshal([]byte(val), &message); err != nil {
		logx.WithContext(l.ctx).Error("消息队列kqueue.SmsMessage 解析json出错   err : %v , val : %s", err, val)
		return err
	}

	if err := l.execService(message); err != nil {
		logx.WithContext(l.ctx).Error("消息队列kqueue.SmsMessage 执行出错  err : %v , val : %s , message:%+v", err, val, message)
		return err
	}

	return nil
}

func (l *SendSmsMq) execService(msg kqueue.SmsMessage) error {
	data := model.BasicSms{Phone: msg.Mobile, Title: msg.Title, TempletId: msg.TmpId}
	id, err := l.svcCtx.BasicSmsSev.Create(l.ctx, data)
	if err != nil {
		return err
	}
	logx.WithContext(l.ctx).Infof("\n 创建BasicSms记录 id=%d", id)
	return nil
}
