package logic

import (
	"context"
	"strconv"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"looklook/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

var ErrUserNoExistsError = xerr.NewErrMsg("用户不存在")

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {

	logx.Errorf("GetUserInfo 11111")
	logx.Errorf(strconv.FormatInt(in.Id, 10))
	logx.Errorf("GetUserInfo 11111222")
	//user, err := l.svcCtx.UserModel.FindOne(l.ctx,in.Id)
	user, err := l.svcCtx.MemUserSev.Get(l.ctx, in.Id, "")
	logx.Errorf("GetUserInfo 11111aaa")
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetUserInfo find user db err , id:%d , err:%v", in.Id, err)
	}
	if user == nil {
		return nil, errors.Wrapf(ErrUserNoExistsError, "id:%d", in.Id)
	}
	logx.Errorf("GetUserInfo 11111222")
	var respUser usercenter.MemUser
	_ = copier.Copy(&respUser, user)
	logx.Errorf("GetUserInfo 11111333--------")
	logx.Errorv(&respUser)
	logx.Errorf("GetUserInfo 11111333-----")
	return &usercenter.GetUserInfoResp{
		User: &respUser,
	}, nil

}
