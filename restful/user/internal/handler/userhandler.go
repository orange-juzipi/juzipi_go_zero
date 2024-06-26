package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"juzipi_go_zero/restful/user/internal/logic"
	"juzipi_go_zero/restful/user/internal/svc"
	"juzipi_go_zero/restful/user/internal/types"
)

func UserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserLogic(r.Context(), svcCtx)
		resp, err := l.User(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
