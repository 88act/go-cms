package user

import (
	"net/http"

	"go-cms/app/usercenter/cmd/api/internal/logic/user"
	"go-cms/app/usercenter/cmd/api/internal/svc"
	"go-cms/common/result"
)

func TokenVerifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewTokenVerifyLogic(r.Context(), svcCtx)
		resp, err := l.TokenVerify()
		result.HttpResult(r, w, resp, err, "")
	}
}
