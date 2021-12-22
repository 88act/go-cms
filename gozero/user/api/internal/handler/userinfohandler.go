package handler

import (
	"fmt"
	"net/http"

	"datacenter/user/api/internal/logic"
	"datacenter/user/api/internal/svc"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func userInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("-----111-----UserCheckMiddleware userId = ")
		userId := r.Header.Get("x-user-id")
		l := logic.NewUserInfoLogic(r.Context(), ctx)
		resp, err := l.UserInfo(userId)

		if err != nil {
			fmt.Println(err)
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
