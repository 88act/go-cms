package logic

import (
	"context"

	"go-cms/app/basic/cmd/rpc/internal/svc"
	"go-cms/app/basic/cmd/rpc/pb"
	"go-cms/common/mycache"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyCodeLogic {
	return &VerifyCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 图形码
func (l *VerifyCodeLogic) VerifyCode(in *pb.VerifyCodeReq) (*pb.VerifyCodeResp, error) {
	code, _ := mycache.GetCache().Get(in.Key)
	if code.(string) == in.Code {
		return &pb.VerifyCodeResp{Status: 1}, nil
	} else {
		return &pb.VerifyCodeResp{Status: 0}, nil
	}
}
