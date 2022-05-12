package act

import (
	"context"

	"go-cms/app/act/cmd/api/internal/svc"
	"go-cms/app/act/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShopDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShopDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) ShopDetailLogic {
	return ShopDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShopDetailLogic) ShopDetail(req types.ShopDetailReq) (resp *types.ShopDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
