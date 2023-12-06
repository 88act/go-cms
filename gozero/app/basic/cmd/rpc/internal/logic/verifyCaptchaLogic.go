package logic

import (
	"context"

	"go-cms/app/basic/cmd/rpc/internal/svc"
	"go-cms/app/basic/cmd/rpc/pb"
	"go-cms/common/mycache"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyCaptchaLogic {
	return &VerifyCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 图形码
func (l *VerifyCaptchaLogic) VerifyCaptcha(in *pb.VerifyCodeReq) (*pb.VerifyCodeResp, error) {
	// todo: add your logic here and delete this line
	code, _ := mycache.GetCache().Get(in.Key)
	codeStr := code.(string)
	if codeStr == in.Code {
		return &pb.VerifyCodeResp{Status: 1}, nil
	} else {
		return &pb.VerifyCodeResp{Status: 0}, nil
	}
}
