package user

import (
	"net/http"

	"go-cms/app/usercenter/cmd/api/internal/logic/user"
	"go-cms/app/usercenter/cmd/api/internal/svc"
	"go-cms/common/ctxdata"
	"go-cms/common/result"
)

func DetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewDetailLogic(r.Context(), svcCtx)
		userId := ctxdata.GetUidFromCtx(r.Context())
		resp, err := l.Detail(userId)
		result.HttpResult(r, w, resp, err, userId)
	}
}
