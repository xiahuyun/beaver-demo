// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"

	"beaver-server/app/usercenter/api/internal/logic"
	"beaver-server/app/usercenter/api/internal/svc"
	"beaver-server/app/usercenter/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 用户认证
func authenticationUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AuthenticationUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAuthenticationUserLogic(r.Context(), svcCtx)
		resp, err := l.AuthenticationUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
