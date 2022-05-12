package act

import (
	"context"

	"looklook/app/act/cmd/api/internal/svc"
	"looklook/app/act/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShopListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShopListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ShopListLogic {
	return ShopListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShopListLogic) ShopList(req types.ShopListReq) (resp *types.ShopListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
