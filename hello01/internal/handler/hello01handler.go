package handler

import (
	"net/http"

	"hello01/internal/logic"
	"hello01/internal/svc"
	"hello01/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func Hello01Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewHello01Logic(r.Context(), svcCtx) //逻辑层，将服务的上下文传进去
		resp, err := l.Hello01(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
