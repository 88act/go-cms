package basic

import (
	"net/http"

	"go-cms/app/basic/cmd/api/internal/logic/basic"
	"go-cms/app/basic/cmd/api/internal/svc"
	"go-cms/app/basic/cmd/api/internal/types"
	"go-cms/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req types.PageInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := basic.NewFileListLogic(r.Context(), svcCtx)
		resp, err := l.FileList(&req)
		result.HttpResult(r, w, resp, err, req)
	}
}
