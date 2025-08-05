package user

import (
	"net/http"

	"douyin-mall/atlas/internal/logic/user"
	"douyin-mall/atlas/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 用户登录
func UserLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUserLoginLogic(r.Context(), svcCtx)
		resp, err := l.UserLogin()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
