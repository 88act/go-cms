package order

import (
	"net/http"

	"looklook/common/result"

	"looklook/app/order/cmd/api/internal/logic/order"
	"looklook/app/order/cmd/api/internal/svc"
	"looklook/app/order/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func OrderListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := order.NewOrderListLogic(r.Context(), svcCtx)
		resp, err := l.OrderList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
