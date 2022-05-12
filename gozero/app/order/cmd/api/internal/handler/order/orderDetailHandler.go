package order

import (
	"net/http"

	"go-cms/common/result"

	"go-cms/app/order/cmd/api/internal/logic/order"
	"go-cms/app/order/cmd/api/internal/svc"
	"go-cms/app/order/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func OrderDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := order.NewOrderDetailLogic(r.Context(), svcCtx)
		resp, err := l.OrderDetail(&req)
		result.HttpResult(r, w, resp, err)
	}
}
