package handler

import (
	"net/http"

	"atlas/internal/logic"
	"atlas/internal/svc"
	"atlas/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AtlasHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAtlasLogic(r.Context(), svcCtx)
		resp, err := l.Atlas(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
