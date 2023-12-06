package cms

import (
	"context"

	"go-cms/app/cms/cmd/api/internal/svc"
	"go-cms/app/cms/cmd/api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type ArtDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArtDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) ArtDetailLogic {
	return ArtDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArtDetailLogic) ArtDetail(req *types.IdReq) (resp *types.ArtDetailResp, err error) {
	art, err := l.svcCtx.CmsArtSev.Get(l.ctx, req.Id, "")
	if err != nil {
		return nil, err
	}
	var obj types.ArtDetailResp
	_ = copier.Copy(&obj, art)
	obj.CreatedAt = art.CreatedAt.Unix()
	return &obj, nil
}
