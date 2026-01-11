// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"beaver-server/app/usercenter/api/internal/logic"
	"beaver-server/app/usercenter/api/internal/svc"
	"beaver-server/app/usercenter/api/internal/types"
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取邮箱验证码
func getEmailVerificationCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetEmailVerificationCodeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// TODO: add middleware to set requestId
		requestId := r.Header.Get("requestId")
		r = r.WithContext(context.WithValue(r.Context(), "requestId", requestId))

		l := logic.NewGetEmailVerificationCodeLogic(r.Context(), svcCtx)
		resp, err := l.GetEmailVerificationCode(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
