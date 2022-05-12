package logic

import (
	"context"

	"go-cms/app/act/cmd/rpc/act"
	"go-cms/app/act/cmd/rpc/internal/svc"
	"go-cms/app/act/cmd/rpc/pb"
	"go-cms/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type ActDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActDetailLogic {
	return &ActDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// actDetail
func (l *ActDetailLogic) ActDetail(in *pb.ActDetailReq) (*pb.ActDetailResp, error) {

	actAct, err := l.svcCtx.ActActSev.Get(l.ctx, in.Id, "")
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_RecordNotFound), "活动不存在 :%s,err:%v", in.Id, err)
	}
	var respAct act.ActInfo
	_ = copier.Copy(&respAct, actAct)
	//logx.Errorv(respAct)
	return &pb.ActDetailResp{Act: &respAct}, nil
}
