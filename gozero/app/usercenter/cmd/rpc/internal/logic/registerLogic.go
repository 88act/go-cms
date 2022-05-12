package logic

import (
	"context"
	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"looklook/app/usercenter/model"
	"looklook/common/tool"
	"looklook/common/xerr"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

var ErrUserAlreadyRegisterError = xerr.NewErrMsg("user has been 333332222")

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *usercenter.RegisterReq) (*usercenter.RegisterResp, error) {

	logx.Errorf("Register == 222222222 \r\n")
	mapData := map[string]interface{}{"mobile": in.Mobile}
	user, err := l.svcCtx.MemUserSev.GetByMap(l.ctx, mapData, "")
	//logx.Errorf("Register err==========" + err.Error())

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "mobile:%s,err:%v", in.Mobile, err)
	}
	if user != nil {
		return nil, errors.Wrapf(ErrUserAlreadyRegisterError, "Register user exists mobile:%s,err:%v", in.Mobile, err)
	}
	logx.Errorf("Register == 333333 \r\n")
	//var userId int64
	var data = model.MemUser{}
	_ = copier.Copy(&data, in)
	data.Username = in.Mobile
	//data.Mobile = in.Mobile
	data.Password = in.Password
	data.PasswordSlat = tool.Krand(8, tool.KC_RAND_KIND_ALL)
	logx.Errorf("Register == 444 " + data.Mobile)
	logx.Errorf("Register == 444 PasswordSlat " + data.PasswordSlat)

	pwd := data.Password + data.PasswordSlat
	logx.Errorf("Register == 4442  pwd =" + pwd)
	data.Password = tool.Md5ByString(pwd)
	logx.Errorf("Register == 4443 " + data.Password)
	// if utils.IsEmpty(data.Nickname) {
	// 	data.Nickname = tool.Krand(8, tool.KC_RAND_KIND_ALL)
	// }
	logx.Errorf("Register == 555 " + data.Nickname)
	//var data = model.MemUser{Username:in.username,Nickname:, Mobile: in.Mobile, Password: in.Password, PasswordSlat: in.PasswordSlat}
	userId, err := l.svcCtx.MemUserSev.Create(l.ctx, data)
	logx.Errorf("Register == 66666 " + strconv.FormatInt(userId, 10))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Register db user Insert err:%v,user:%+v", err, data)
	}
	//session := db.WithContext(ctx)

	// if err := l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
	// 	user := new(model.User)
	// 	user.Mobile = in.Mobile
	// 	if len(user.Nickname) == 0 {
	// 		user.Nickname = tool.Krand(8, tool.KC_RAND_KIND_ALL)
	// 	}
	// 	if len(in.Password) > 0 {
	// 		user.Password = tool.Md5ByString(in.Password)
	// 	}
	// 	logx.Errorf("Register == 444444 \r\n")
	// 	//insertResult, err := l.svcCtx.UserModel.Insert(ctx, session, user)
	// 	lastId, err := l.svcCtx.UserModel.AddOne(ctx, user)
	// 	if err != nil {
	// 		return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Register db user Insert err:%v,user:%+v", err, user)
	// 	}
	// 	//lastId, err := insertResult.LastInsertId()
	// 	logx.Errorf("Register == 555555  ")
	// 	logx.Errorf(string(lastId))
	// 	if err != nil {
	// 		return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Register db user insertResult.LastInsertId err:%v,user:%+v", err, user)
	// 	}

	// 	logx.Errorf("Register == 66666 \r\n")

	// 	userId = lastId

	// 	// userAuth := new(model.UserAuth)
	// 	// userAuth.UserId = lastId
	// 	// userAuth.AuthKey = in.AuthKey
	// 	// userAuth.AuthType = in.AuthType
	// 	// if _, err := l.svcCtx.UserAuthModel.Insert(ctx, session, userAuth); err != nil {
	// 	// 	return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Register db user_auth Insert err:%v,userAuth:%v", err, userAuth)
	// 	// }
	// 	return nil
	// }); err != nil {
	// 	return nil, err
	// }

	//2、Generate the token, so that the service doesn't call rpc internally
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&usercenter.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId : %d", userId)
	}

	return &usercenter.RegisterResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}
