package basic

import (
	"context"
	"go-cms/app/basic/cmd/api/internal/svc"
	"go-cms/app/basic/cmd/api/internal/types"
	"go-cms/common/baseModel"
	"go-cms/common/ctxdata"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileByKeyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFileByKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileByKeyLogic {
	return &GetFileByKeyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFileByKeyLogic) GetFileByKey(req *types.ValReq) (resp *types.FileDetailResp, err error) {

	if req.Val == "" {
		return nil, errors.New("参数错误")
	}
	mapData := make(map[string]interface{})
	mapData["md5"] = req.Val
	userId := ctxdata.GetUidFromCtx(l.ctx)
	info := types.FileInfo{}
	//logx.Errorf("====1====userid === %v", userId)
	if basicFile, err := l.svcCtx.BasicFileSev.GetByMap(l.ctx, mapData, ""); err == nil {
		//logx.Errorf("====2====ubasicFileuserid === %v", basicFile.UserId)
		if userId > 0 && basicFile.UserId != userId {
			//不是自己的图片 ，创建一个自己的
			basicFileNew := baseModel.BasicFile{}
			_ = copier.Copy(&basicFileNew, basicFile)
			return &types.FileDetailResp{Info: info}, nil
		} else {
			_ = copier.Copy(&info, basicFile)
			return &types.FileDetailResp{Info: info}, nil
		}
	} else {
		return nil, errors.New("不存在")

	}
}
