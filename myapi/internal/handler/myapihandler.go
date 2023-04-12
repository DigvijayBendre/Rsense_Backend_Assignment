package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"myapi/internal/logic"
	"myapi/internal/svc"
	"myapi/internal/types"
)

func MyapiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewMyapiLogic(r.Context(), svcCtx)
		resp, err := l.Myapi(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
