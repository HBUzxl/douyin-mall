package user

import (
	"net/http"

	"github.com/HBUzxl/douyin-mall/atlas/internal/logic/user"
	"github.com/HBUzxl/douyin-mall/atlas/internal/svc"
	"github.com/HBUzxl/douyin-mall/atlas/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 添加用户黑名单
func UserAddUserBlacklistHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddUserBlacklistReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserAddUserBlacklistLogic(r.Context(), svcCtx)
		resp, err := l.UserAddUserBlacklist(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
