package act

import (
	"net/http"

	"looklook/common/result"

	"looklook/app/act/cmd/api/internal/logic/act"
	"looklook/app/act/cmd/api/internal/svc"
	"looklook/app/act/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ShopListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShopListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := act.NewShopListLogic(r.Context(), svcCtx)
		resp, err := l.ShopList(req)
		result.HttpResult(r, w, resp, err)
	}
}
