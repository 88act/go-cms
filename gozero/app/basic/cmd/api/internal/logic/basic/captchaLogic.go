package basic

import (
	"context"

	"go-cms/app/basic/cmd/api/internal/svc"
	"go-cms/app/basic/cmd/api/internal/types"
	"go-cms/common/mycache"

	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/logx"
)

type CaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) CaptchaLogic {
	return CaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
// var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

func (l *CaptchaLogic) Captcha(req *types.CaptchaReq) (resp *types.CaptchaResp, err error) {
	// -------rpc 模式----------
	// rpcResp, err := l.svcCtx.BasicRpc.Captcha(l.ctx, &basic.CaptchaReq{})
	// if err != nil {
	// 	return nil, err
	// }
	// myResp := types.CaptchaResp{}
	// _ = copier.Copy(&myResp, rpcResp)
	// return &myResp, nil

	//--------本地模式 ----------------------
	// todo: add your logic here and delete this line
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(40, 120, 4, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	//cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	id, content, answer := cp.Driver.GenerateIdQuestionAnswer()
	item, err := cp.Driver.DrawCaptcha(content)
	if err != nil {
		return nil, err
	}
	key := "cap_" + id
	// 放入缓存
	//logx.Errorv("放入缓存 = " + answer)
	mycache.GetCache().Set(key, answer, 300)
	b64s := item.EncodeB64string()
	return &types.CaptchaResp{Image: b64s, Key: key, ExpireTime: 300}, nil

}
