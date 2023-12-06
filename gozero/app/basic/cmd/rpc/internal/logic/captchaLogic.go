package logic

import (
	"context"

	"go-cms/app/basic/cmd/rpc/internal/svc"
	"go-cms/app/basic/cmd/rpc/pb"
	"go-cms/common/mycache"

	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/logx"
)

type CaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CaptchaLogic {
	return &CaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
// var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

// 图形码
func (l *CaptchaLogic) Captcha(in *pb.CaptchaReq) (*pb.CaptchaResp, error) {
	// todo: add your logic here and delete this line
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(40, 120, 4, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	//cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	id, content, answer := cp.Driver.GenerateIdQuestionAnswer()
	item, err := cp.Driver.DrawCaptcha(content)
	if err != nil {
		return &pb.CaptchaResp{}, err
	}
	key := "cap_" + id
	//放入缓存
	logx.Errorv("放入缓存 = " + answer)
	mycache.GetCache().Set(key, answer, 0)
	b64s := item.EncodeB64string()
	return &pb.CaptchaResp{Image: b64s, Key: key, ExpireTime: 120}, nil
}
