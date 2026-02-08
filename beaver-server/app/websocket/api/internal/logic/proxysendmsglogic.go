// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"beaver-server/app/pkg/websocket"
	wstypes "beaver-server/app/pkg/websocket"
	"beaver-server/app/websocket/api/internal/svc"
	"beaver-server/app/websocket/api/internal/types"

	wsutils "beaver-server/app/websocket/api/internal/logic/websocket/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProxySendMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// HTTP API发送消息
func NewProxySendMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProxySendMsgLogic {
	return &ProxySendMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProxySendMsgLogic) ProxySendMsg(req *types.ProxySendMsgReq) (resp *types.ProxySendMsgRes, err error) {
	fmt.Println("收到ws转发的消息")
	// 将map转换为json.RawMessage
	bodyBytes, err := json.Marshal(req.Body)
	if err != nil {
		return nil, err
	}

	content := wstypes.WebsocketContent{
		Timestamp: 0,
		Data: wstypes.WebsocketData{
			Type:           websocket.WebsocketType(req.Type),
			Body:           bodyBytes,
			ConversationID: req.ConversationId,
		},
	}

	fmt.Println("消息内容：", string(bodyBytes))
	fmt.Println("发送者ID: ", req.UserID, "目标ID: ", req.TargetID)

	fmt.Println("命令类型：", req.Command, "，消息类型：", req.Type, "会话ID: ", req.ConversationId)

	wsutils.SendMsgToUser(req.TargetID, wstypes.WebsocketCommand(req.Command), content)

	return &types.ProxySendMsgRes{}, nil
}
