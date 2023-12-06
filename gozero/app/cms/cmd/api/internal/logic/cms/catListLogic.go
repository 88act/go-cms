package cms

import (
	"context"

	"go-cms/app/cms/cmd/api/internal/svc"
	"go-cms/app/cms/cmd/api/internal/types"
	"go-cms/app/cms/model"
	"go-cms/common/mycache"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type CatListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCatListLogic(ctx context.Context, svcCtx *svc.ServiceContext) CatListLogic {
	return CatListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CatListLogic) CatList() (resp *types.CatListResp, err error) {
	cacheKey := "catList"
	if cacheData, err := mycache.GetCache().GetObj(cacheKey); err == nil {
		respCa := cacheData.(types.CatListResp)
		return &respCa, nil
	} else {
		seq := model.CmsCatSearch{}
		seq.Page = 1
		seq.PageSize = 1000
		list, total, err := l.svcCtx.CmsCatSev.GetList(l.ctx, seq, "")
		if err != nil {
			return nil, err
		}
		var list2 []types.Cat
		for _, item := range list {
			var obj types.Cat
			_ = copier.Copy(&obj, item)
			obj.CreatedAt = item.CreatedAt.Unix()
			list2 = append(list2, obj)
		}
		myResp := types.CatListResp{
			Total: total,
			List:  list2,
		}
		mycache.GetCache().SetObj(cacheKey, myResp, 0)
		return &myResp, nil
	}

}
