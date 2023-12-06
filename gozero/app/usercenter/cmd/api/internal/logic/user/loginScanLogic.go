package user

import (
	"context"
	"errors"
	"time"

	"go-cms/app/usercenter/cmd/api/internal/svc"
	"go-cms/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginScanLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginScanLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginScanLogic {
	return LoginScanLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginScanLogic) LoginScan(req *types.LoginScanReq) (resp *types.LoginResp, err error) {
	//userId := ctxdata.GetUidFromCtx(l.ctx)
	if req.Value == "" {
		return nil, errors.New("参数错误")
	}
	// customer, err := l.svcCtx.JqCustomerSev.GetByGuid(l.ctx, req.CuGuid, "")
	// if err != nil {
	// 	return nil, errors.New("不存在的客户")
	// }
	// if customer.Status != 1 {
	// 	return nil, errors.New("客户已禁用")
	// }
	mapWhere := make(map[string]interface{})
	mapWhere["scan"] = req.Value
	//mapWhere["cu_id"] = customer.Id
	user, err := l.svcCtx.MemUserSev.GetByMap(l.ctx, mapWhere, "")
	if err == nil && user.Id > 0 {
		sTime := *user.ScanTime
		sTime = sTime.Add(time.Minute * 5)
		if time.Now().After(sTime) {
			return nil, errors.New("失败，请在5分钟内登录")
		}
		if user.Status != 1 {
			return nil, errors.New("用户状态异常")
		}
		myResp, err := GenerateToken(l.svcCtx.Config.JwtAuth.AccessExpire, l.svcCtx.Config.JwtAuth.AccessSecret, user)
		if err != nil {
			return nil, errors.New("生成token失败")
		}
		return myResp, nil
	} else {
		return nil, errors.New("扫码不存在")
	}
}
