package basic

import (
	"context"

	"go-cms/app/basic/cmd/api/internal/svc"
	"go-cms/app/basic/cmd/api/internal/types"
	"go-cms/common/ctxdata"
	"go-cms/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DelFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelFileListLogic {
	return &DelFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelFileListLogic) DelFileList(req *types.GuidListReq) (resp *types.OkResp, err error) {
	uid := ctxdata.GetUidFromCtx(l.ctx)
	if uid > 0 {
		mapWhere := make(map[string]interface{})
		mapWhere["user_id"] = uid
		//	logx.Errorf("===== MyFileList userId ====%v ", ctxdata.GetUidFromCtx(l.ctx))
		if err := l.svcCtx.BasicFileSev.DeleteByGuids(l.ctx, mapWhere, req.GuidList); err == nil {
			return &types.OkResp{Status: 1, Msg: "成功"}, nil
		} else {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "删除失败", err)
		}
	} else {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.Fail), "删除失败，认证失败", err)
	}

}
