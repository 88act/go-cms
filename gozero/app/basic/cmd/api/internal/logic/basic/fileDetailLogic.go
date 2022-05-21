package basic

import (
	"context"

	"go-cms/app/basic/cmd/api/internal/svc"
	"go-cms/app/basic/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) FileDetailLogic {
	return FileDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileDetailLogic) FileDetail(req *types.FileDetailReq) (resp *types.FileDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
