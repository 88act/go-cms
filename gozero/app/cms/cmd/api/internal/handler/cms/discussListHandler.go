package cms

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-cms/app/cms/cmd/api/internal/logic/cms"
	"go-cms/app/cms/cmd/api/internal/svc"
	"go-cms/app/cms/cmd/api/internal/types"
	"go-cms/common/result"
)

func DiscussListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := cms.NewDiscussListLogic(r.Context(), svcCtx)
		resp, err := l.DiscussList(&req)
		result.HttpResult(r, w, resp, err, req)
	}
}
