package logic

import (
	"context"

	"go-cms/app/act/cmd/rpc/act"
	"go-cms/app/act/cmd/rpc/internal/svc"
	"go-cms/app/act/cmd/rpc/pb"
	"go-cms/app/act/model"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type ActListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActListLogic {
	return &ActListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// actList
func (l *ActListLogic) ActList(in *pb.ActListReq) (*pb.ActListResp, error) {
	var seq model.ActActSearch
	_ = copier.Copy(&seq, in)
	list, _, _ := l.svcCtx.ActActSev.GetList(l.ctx, seq, "")
	var resp []*act.ActInfo
	if len(list) > 0 {
		for _, actAct := range list {
			var obj act.ActInfo
			_ = copier.Copy(&obj, actAct)
			//act.CreateTime = homestayOrder.CreateTime.Unix()
			//act.LiveStartDate = homestayOrder.LiveStartDate.Unix()
			//act.LiveEndDate = homestayOrder.LiveEndDate.Unix()
			resp = append(resp, &obj)
		}
	}
	//var actListResp *act.ActListResp
	//_ = copier.Copy(actListResp, list)
	return &pb.ActListResp{List: resp}, nil
}
