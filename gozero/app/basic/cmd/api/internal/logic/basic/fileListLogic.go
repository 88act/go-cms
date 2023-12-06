package basic

import (
	"context"

	"go-cms/app/basic/cmd/api/internal/svc"
	"go-cms/app/basic/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) FileListLogic {
	return FileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileListLogic) FileList(req *types.PageInfoReq) (out *types.FileListResp, err error) {

	// fileListResp, err := l.svcCtx.BasicRpc.FileList(l.ctx, &basic.FileListReq{Ids: req.Ids})
	// if err != nil {
	// 	return nil, err
	// }

	// var resp types.FileListResp
	// _ = copier.Copy(&resp, fileListResp)

	return out, nil
}
