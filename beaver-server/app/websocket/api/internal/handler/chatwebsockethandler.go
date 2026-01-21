// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"

	"beaver-server/app/websocket/api/internal/logic"
	"beaver-server/app/websocket/api/internal/svc"
	"beaver-server/app/websocket/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// WebSocket连接
func chatWebsocketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 打印所有请求头信息用于调试
		logx.Info("WebSocket请求头信息:")
		for k, v := range r.Header {
			logx.Infof("  %s: %v", k, v)
		}

		var req types.ChatWebsocketRequest
		if err := httpx.Parse(r, &req); err != nil {
			logx.Errorf("parse request failed: %v", err)
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewChatWebsocketLogic(r.Context(), svcCtx)
		resp, err := l.ChatWebsocket(&req, w, r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
