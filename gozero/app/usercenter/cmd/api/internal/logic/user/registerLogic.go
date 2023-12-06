package user

import (
	"context"
	"time"

	"go-cms/app/usercenter/cmd/api/internal/svc"
	"go-cms/app/usercenter/cmd/api/internal/types"
	"go-cms/app/usercenter/model"
	"go-cms/common/ctxdata"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.LoginReq) (resp *types.LoginResp, err error) {
	mapData := map[string]interface{}{"username": req.Username}
	user, err := l.svcCtx.MemUserSev.GetByMap(l.ctx, mapData, "")
	if err == nil && user.Id > 0 {
		return resp, errors.New("账号已存在")
	} else {
		var data = model.MemUser{}
		//data.Username = in.Mobile
		data.Username = req.Username
		data.Password = req.Password
		data.PasswordSlat = "" // tool.Krand(8, tool.KC_RAND_KIND_ALL)
		//pwd := data.Password + data.PasswordSlat
		//data.Password = tool.Md5ByString(pwd)
		//if utils.IsEmpty(data.Nickname) {
		data.Nickname = req.Username
		//}
		data.Status = 1
		if userId, err := l.svcCtx.MemUserSev.Create(l.ctx, data); err == nil {
			data.Id = userId
			myResp, err := GenerateToken(l.svcCtx.Config.JwtAuth.AccessExpire, l.svcCtx.Config.JwtAuth.AccessSecret, data)
			if err != nil {
				return nil, errors.New("生成token失败")
			}
			return myResp, nil
		} else {
			return resp, errors.New("注册出错 ")
		}
	}
}

func GenerateToken(accessExpire int64, accessSecret string, user model.MemUser) (*types.LoginResp, error) {
	now := time.Now().Unix()

	accessToken, err := GetJwtToken(accessSecret, now, accessExpire, user.Id, user.UserType, user.CuId)
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func GetJwtToken(secretKey string, iat, seconds, userId int64, userType int, cuId int64) (string, error) {

	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[ctxdata.CtxKeyJwtUserId] = userId
	claims[ctxdata.CtxKeyJwtUserType] = userType
	claims[ctxdata.CtxKeyCuId] = cuId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
