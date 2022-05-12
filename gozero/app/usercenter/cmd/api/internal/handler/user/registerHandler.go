package user

import (
	"net/http"

	"go-cms/app/usercenter/cmd/api/internal/logic/user"
	"go-cms/app/usercenter/cmd/api/internal/svc"
	"go-cms/app/usercenter/cmd/api/internal/types"
	"go-cms/common/result"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		logx.Errorf("Register查询sql=")
		l := user.NewRegisterLogic(r.Context(), ctx)
		resp, err := l.Register(req)
		result.HttpResult(r, w, resp, err)
	}
}