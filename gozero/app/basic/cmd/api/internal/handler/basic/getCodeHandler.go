package basic

import (
	"net/http"

	"go-cms/common/result"

	"go-cms/app/basic/cmd/api/internal/logic/basic"
	"go-cms/app/basic/cmd/api/internal/svc"
	"go-cms/app/basic/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := basic.NewGetCodeLogic(r.Context(), svcCtx)
		resp, err := l.GetCode(&req)
		result.HttpResult(r, w, resp, err)
	}
}
