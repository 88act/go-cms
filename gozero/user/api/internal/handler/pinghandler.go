package handler

import (
	"fmt"
	"net/http"

	"datacenter/user/api/internal/logic"
	"datacenter/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func pingHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("-----111-----pingHandler userId = ")
		l := logic.NewPingLogic(r.Context(), ctx)
		err := l.Ping()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
