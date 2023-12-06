package cms

import (
	"net/http"

	"go-cms/app/cms/cmd/api/internal/logic/cms"
	"go-cms/app/cms/cmd/api/internal/svc"
	"go-cms/common/result"
)

func CatListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := cms.NewCatListLogic(r.Context(), svcCtx)
		resp, err := l.CatList()
		result.HttpResult(r, w, resp, err, "")
	}
}
