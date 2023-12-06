package basic

import (
	"context"

	"go-cms/app/basic/cmd/api/internal/svc"
	"go-cms/app/basic/cmd/api/internal/types"

	"github.com/pkg/errors"
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

func (l *FileDetailLogic) FileDetail(req *types.ValReq) (resp *types.FileDetailResp, err error) {
	if req.Val == "" {
		return nil, errors.New("参数错误")
	}
	mapData := make(map[string]interface{})
	mapData["guid"] = req.Val
	//logx.Errorf("====1====userid === %v", userId)
	if basicFile, err := l.svcCtx.BasicFileSev.GetByMap(l.ctx, mapData, ""); err == nil {
		basicFileNew := types.FileDetailResp{}
		basicFileNew.Info.Guid = basicFile.Guid
		basicFileNew.Info.Name = basicFile.Name
		basicFileNew.Info.Path = basicFile.Path
		basicFileNew.Info.MediaType = basicFile.MediaType
		basicFileNew.Info.Size = basicFile.Size
		return &basicFileNew, nil
	} else {
		return resp, err
	}
}
