package cms

import (
	"context"
	"fmt"

	"go-cms/app/cms/cmd/api/internal/svc"
	"go-cms/app/cms/cmd/api/internal/types"
	"go-cms/app/cms/model"
	"go-cms/common/mycache"

	"github.com/gogf/gf/util/gconv"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type ArtListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArtListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ArtListLogic {
	return ArtListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArtListLogic) ArtList(req *types.PageInfoReq) (resp *types.ArtListResp, err error) {
	cacheKey := fmt.Sprintf("%s_%s_%d_%d", "artList", req.Key, req.Page, req.PageSize)
	if cacheData, err := mycache.GetCache().GetObj(cacheKey); err == nil {
		respCa := cacheData.(types.ArtListResp)
		return &respCa, nil
	} else {
		seq := model.CmsArtSearch{}
		seq.Page = req.Page
		seq.PageSize = req.PageSize
		catId := gconv.Int64(req.Key)
		if catId <= 0 {
			return resp, errors.New("参数错误")
		}
		seq.CatId = &catId
		list, total, err := l.svcCtx.CmsArtSev.GetActList(l.ctx, seq)
		if err != nil {
			return nil, err
		}
		var list2 []types.ArtDetailResp
		for _, item := range list {
			var obj types.ArtDetailResp
			_ = copier.Copy(&obj, item)
			obj.CreatedAt = item.CreatedAt.Unix()
			list2 = append(list2, obj)
		}
		myResp := types.ArtListResp{
			Total: total,
			List:  list2,
		}
		mycache.GetCache().SetObj(cacheKey, myResp, 0)
		return &myResp, nil
	}

}
