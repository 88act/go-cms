package basic

import (
	"context"
	"go-cms/app/basic/cmd/api/internal/svc"
	"go-cms/app/basic/cmd/api/internal/types"

	"go-cms/common/ctxdata"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type MyFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMyFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MyFileListLogic {
	return &MyFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MyFileListLogic) MyFileList(req *types.PageInfoReq) (resp *types.FileListResp, err error) {

	mapData := make(map[string]interface{})
	mapData["user_id"] = ctxdata.GetUidFromCtx(l.ctx)
	//	logx.Errorf("===== MyFileList userId ====%v ", ctxdata.GetUidFromCtx(l.ctx))
	if list, err := l.svcCtx.BasicFileSev.GetListByMap(l.ctx, mapData, "", "id desc"); err == nil {
		myResp := new(types.FileListResp)
		myResp.List = make([]types.FileInfo, 0)
		for _, v := range list {
			info := types.FileInfo{}
			_ = copier.Copy(&info, v)
			// info.Guid = v.Guid
			// info.Path = v.Path
			// info.Size = v.Size
			// info.MediaType = v.MediaType
			myResp.List = append(myResp.List, info)
			myResp.Total = 0
		}
		return myResp, nil
	} else {
		return nil, errors.New("不存在")
	}
}
