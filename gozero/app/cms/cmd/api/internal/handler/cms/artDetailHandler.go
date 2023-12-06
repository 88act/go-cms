package cms

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-cms/app/cms/cmd/api/internal/logic/cms"
	"go-cms/app/cms/cmd/api/internal/svc"
	"go-cms/app/cms/cmd/api/internal/types"
	"go-cms/common/result"
)

func ArtDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := cms.NewArtDetailLogic(r.Context(), svcCtx)
		resp, err := l.ArtDetail(&req)
		result.HttpResult(r, w, resp, err, req)
	}
}
