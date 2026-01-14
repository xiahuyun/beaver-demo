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

// 用户登录
func loginUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginUserLogic(r.Context(), svcCtx)
		resp, err := l.LoginUser(&req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, &types.LoginUserResponse{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
