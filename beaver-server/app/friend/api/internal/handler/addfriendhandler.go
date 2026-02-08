// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"

	"beaver-server/app/friend/api/internal/logic"
	"beaver-server/app/friend/api/internal/svc"
	"beaver-server/app/friend/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 发送好友验证请求
func addFriendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddFriendReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAddFriendLogic(r.Context(), svcCtx)
		resp, err := l.AddFriend(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
