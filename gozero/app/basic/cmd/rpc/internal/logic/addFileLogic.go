package logic

import (
	"context"

	"go-cms/app/basic/cmd/rpc/internal/svc"
	"go-cms/app/basic/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFileLogic {
	return &AddFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// addFile 新增文件
func (l *AddFileLogic) AddFile(in *pb.AddFileReq) (*pb.AddFileResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddFileResp{}, nil
}
