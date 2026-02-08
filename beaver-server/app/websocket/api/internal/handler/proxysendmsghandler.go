// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"

	"beaver-server/app/websocket/api/internal/logic"
	"beaver-server/app/websocket/api/internal/svc"
	"beaver-server/app/websocket/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// HTTP API发送消息
func proxySendMsgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProxySendMsgReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewProxySendMsgLogic(r.Context(), svcCtx)
		resp, err := l.ProxySendMsg(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
