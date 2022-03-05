package handler

import (
	"fmt"
	"net/http"

	"datacenter/user/api/internal/logic"
	"datacenter/user/api/internal/svc"
	"datacenter/user/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func registerHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		fmt.Println("registerHandler----1---------")
		fmt.Println(req)
		l := logic.NewRegisterLogic(r.Context(), ctx)
		fmt.Println("registerHandler---2---------")
		err := l.Register(req)
		fmt.Println("registerHandler----3---------")
		fmt.Println(err)
		fmt.Println("registerHandler----4---------")
		fmt.Println(w)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
